package pull_request

import (
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/pullrequest"
	"context"

	"github.com/google/uuid"
)

func (r *PullRequestRepository) Get(ctx context.Context, id uuid.UUID) (*ent.PullRequest, error) {
	return r.client.PullRequest.Get(ctx, id)
}

func (r *PullRequestRepository) GetWithAuthor(ctx context.Context, id uuid.UUID) (*ent.PullRequest, error) {
	return r.client.PullRequest.Query().
		Where(pullrequest.IDEQ(id)).
		WithAuthor().
		Only(ctx)
}
