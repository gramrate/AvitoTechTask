package pull_request

import (
	"AvitoTechTask/pkg/ent"
	"context"
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

	return create.Save(ctx)
}
