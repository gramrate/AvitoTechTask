package pull_request

import (
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/internal/domain/types"
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/pullrequest"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// UpdateStatusAndGetWithRelations обновляет статус и возвращает PR с отношениями
func (r *Repo) UpdateStatusAndGetWithRelations(ctx context.Context, prID uuid.UUID, newStatus types.PullRequestStatus) (*ent.PullRequest, error) {
	// Сначала обновляем статус
	_, err := r.client.PullRequest.UpdateOneID(prID).
		SetStatus(newStatus).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.ErrPRNotFound
		}
		return nil, fmt.Errorf("updating pull request status: %w", err)
	}

	// Затем возвращаем PR с отношениями
	return r.GetWithReviewers(ctx, prID)
}

// ReassignReviewer переназначает ревьювера в PR в рамках одной транзакции
func (r *Repo) ReassignReviewer(ctx context.Context, prID uuid.UUID, oldReviewerID uuid.UUID, newReviewerID uuid.UUID) (*ent.PullRequest, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}

	// Получаем PR с ревьюверами чтобы получить текущий список
	pr, err := tx.PullRequest.Query().
		Where(pullrequest.IDEQ(prID)).
		WithReviewers().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, rollback(tx, errorz.ErrPRNotFound)
		}
		return nil, rollback(tx, fmt.Errorf("checking PR existence: %w", err))
	}

	// Создаем новый список ревьюверов (заменяем старого на нового)
	var newReviewerIDs []uuid.UUID
	if pr.Edges.Reviewers != nil {
		for _, reviewer := range pr.Edges.Reviewers {
			if reviewer.ID != oldReviewerID {
				newReviewerIDs = append(newReviewerIDs, reviewer.ID)
			}
		}
	}

	// Добавляем нового ревьювера, если он указан
	if newReviewerID != uuid.Nil {
		newReviewerIDs = append(newReviewerIDs, newReviewerID)
	}

	// Обновляем PR с новым списком ревьюверов
	_, err = tx.PullRequest.UpdateOneID(prID).
		ClearReviewers().                  // Очищаем всех ревьюверов
		AddReviewerIDs(newReviewerIDs...). // Добавляем новый список
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("failed to update reviewers: %w", err))
	}

	// Получаем обновленный PR с ревьюверами
	updatedPR, err := tx.PullRequest.Query().
		Where(pullrequest.IDEQ(prID)).
		WithReviewers().
		Only(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("failed to get updated PR with reviewers: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return updatedPR, nil
}
