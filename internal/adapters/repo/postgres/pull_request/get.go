package pull_request

import (
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/internal/domain/types"
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/pullrequest"
	"AvitoTechTask/pkg/ent/team"
	"AvitoTechTask/pkg/ent/user"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *Repo) Get(ctx context.Context, id uuid.UUID) (*ent.PullRequest, error) {
	pr, err := r.client.PullRequest.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.ErrPRNotFound
		}
		return nil, err
	}
	return pr, nil
}

func (r *Repo) GetWithReviewers(ctx context.Context, id uuid.UUID) (*ent.PullRequest, error) {
	pr, err := r.client.PullRequest.Query().
		Where(pullrequest.IDEQ(id)).
		WithReviewers().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.ErrPRNotFound
		}
		return nil, err
	}
	return pr, nil
}

// GetByReviewerIDInTransaction - возвращает PR для ревьювера в одной транзакции
// status - опциональный фильтр по статусу (если nil, то не фильтруется)
func (r *Repo) GetByReviewerIDInTransaction(ctx context.Context, reviewerID uuid.UUID, status *types.PullRequestStatus) ([]*ent.PullRequest, error) {
	// Проверяем существование пользователя
	_, err := r.client.User.Get(ctx, reviewerID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.ErrUserNotFound
		}
		return nil, fmt.Errorf("checking reviewer existence: %w", err)
	}

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

// GetAvailableReviewers возвращает доступных активных ревьюверов из команды автора (исключая самого автора)
func (r *Repo) GetAvailableReviewers(ctx context.Context, authorID uuid.UUID) ([]*ent.User, error) {
	// Получаем автора с его командой
	author, err := r.client.User.Query().
		Where(user.IDEQ(authorID)).
		WithTeam().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get author: %w", err)
	}

	// Если у автора нет команды - возвращаем пустой список
	if author.Edges.Team == nil {
		return []*ent.User{}, nil
	}

	// Получаем всех активных пользователей из команды автора, исключая самого автора
	availableReviewers, err := r.client.User.Query().
		Where(
			user.HasTeamWith(team.IDEQ(author.Edges.Team.ID)), // из команды автора
			user.IsActive(true),  // только активные
			user.IDNEQ(authorID), // исключая автора
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get available reviewers: %w", err)
	}

	return availableReviewers, nil
}

// GetAvailableReviewersFromReviewerTeam возвращает доступных ревьюверов из команды заменяемого ревьювера
// (исключая уже назначенных ревьюверов и автора PR)
func (r *Repo) GetAvailableReviewersFromReviewerTeam(ctx context.Context, oldReviewerID uuid.UUID, prID uuid.UUID) ([]*ent.User, error) {
	// Получаем PR с автором и ревьюверами
	pr, err := r.client.PullRequest.Query().
		Where(pullrequest.IDEQ(prID)).
		WithAuthor().
		WithReviewers().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.ErrPRNotFound
		}
		return nil, fmt.Errorf("failed to get pull request: %w", err)
	}

	// Получаем заменяемого ревьювера с его командой
	oldReviewer, err := r.client.User.Query().
		Where(user.IDEQ(oldReviewerID)).
		WithTeam().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get old reviewer: %w", err)
	}

	// Если у заменяемого ревьювера нет команды - возвращаем пустой список
	if oldReviewer.Edges.Team == nil {
		return []*ent.User{}, nil
	}

	// Собираем ID уже назначенных ревьюверов (исключая заменяемого)
	excludedUserIDs := []uuid.UUID{pr.AuthorID} // исключаем автора
	for _, reviewer := range pr.Edges.Reviewers {
		if reviewer.ID != oldReviewerID { // исключаем заменяемого ревьювера
			excludedUserIDs = append(excludedUserIDs, reviewer.ID)
		}
	}

	// Получаем активных пользователей из команды заменяемого ревьювера
	return r.client.User.Query().
		Where(
			user.HasTeamWith(team.IDEQ(oldReviewer.Edges.Team.ID)), // из команды заменяемого ревьювера
			user.IsActive(true),              // только активные
			user.IDNotIn(excludedUserIDs...), // исключаем автора и уже назначенных ревьюверов
		).
		All(ctx)
}
