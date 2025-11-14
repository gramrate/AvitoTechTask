package pull_request

import (
	"AvitoTechTask/internal/domain/types"
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/pullrequest"
	"AvitoTechTask/pkg/ent/user"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *PullRequestRepository) Get(ctx context.Context, id uuid.UUID) (*ent.PullRequest, error) {
	return r.client.PullRequest.Get(ctx, id)
}

// GetByReviewerIDInTransaction - возвращает PR для ревьювера в одной транзакции
// status - опциональный фильтр по статусу (если nil, то не фильтруется)
func (r *PullRequestRepository) GetByReviewerIDInTransaction(ctx context.Context, reviewerID uuid.UUID, status *types.PullRequestStatus) ([]*ent.PullRequest, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}

	// Базовый запрос - PR, где пользователь ревьювер
	query := tx.PullRequest.Query().
		Where(pullrequest.HasReviewersWith(user.IDEQ(reviewerID))).
		Order(ent.Desc(pullrequest.FieldCreatedAt)).
		WithAuthor().
		WithReviewers()

	// Добавляем фильтр по статусу только если он передан
	if status != nil {
		query = query.Where(pullrequest.StatusEQ(*status))
	}

	pullRequests, err := query.All(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("getting pull requests for reviewer: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return pullRequests, nil
}

// Вспомогательная функция для rollback
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
	}
	return err
}
