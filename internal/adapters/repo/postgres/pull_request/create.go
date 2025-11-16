package pull_request

import (
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/pkg/ent"
	"context"
	"fmt"
	"strings"
)

func (r *Repo) Create(ctx context.Context, prEntity *ent.PullRequest) (*ent.PullRequest, error) {
	create := r.client.PullRequest.Create().
		SetID(prEntity.ID).
		SetPullRequestName(prEntity.PullRequestName).
		SetAuthorID(prEntity.AuthorID).
		SetStatus(prEntity.Status).
		SetAssignedReviewers(prEntity.AssignedReviewers)

	if prEntity.MergedAt != nil {
		create.SetMergedAt(*prEntity.MergedAt)
	}

	pr, err := create.Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			// Проверяем, что ошибка связана с уникальностью ID PR
			if strings.Contains(err.Error(), "id") || strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "primary") {
				return nil, fmt.Errorf("create pull request: %w", errorz.ErrPRNameAlreadyUsed)
			}
		}
		return nil, fmt.Errorf("create pull request: %w", err)
	}
	return pr, nil
}
