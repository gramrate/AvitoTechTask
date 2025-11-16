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

	// Проверяем существование PR
	_, err = tx.PullRequest.Get(ctx, prID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, rollback(tx, errorz.ErrPRNotFound)
		}
		return nil, rollback(tx, fmt.Errorf("checking PR existence: %w", err))
	}

	// Проверяем существование старого ревьювера
	_, err = tx.User.Get(ctx, oldReviewerID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, rollback(tx, errorz.ErrUserNotFound)
		}
		return nil, rollback(tx, fmt.Errorf("checking old reviewer existence: %w", err))
	}

	// Если указан новый ревьювер, проверяем его существование
	if newReviewerID != uuid.Nil {
		_, err = tx.User.Get(ctx, newReviewerID)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, rollback(tx, errorz.ErrUserNotFound)
			}
			return nil, rollback(tx, fmt.Errorf("checking new reviewer existence: %w", err))
		}
	}

	// Удаляем старого ревьювера и добавляем нового (если он указан)
	update := tx.PullRequest.UpdateOneID(prID).
		RemoveReviewerIDs(oldReviewerID)

	if newReviewerID != uuid.Nil {
		update = update.AddReviewerIDs(newReviewerID)
	}

	_, err = update.Save(ctx)
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
