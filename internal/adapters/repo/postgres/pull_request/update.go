package pull_request

import (
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/ent/pullrequest"
	"context"

	"github.com/google/uuid"
)

func (r *Repo) GetWithReviewers(ctx context.Context, id uuid.UUID) (*ent.PullRequest, error) {
	return r.client.PullRequest.Query().
		Where(pullrequest.IDEQ(id)).
		WithReviewers().
		Only(ctx)
}
