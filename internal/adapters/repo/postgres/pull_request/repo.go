package pull_request

import (
	"AvitoTechTask/pkg/ent"
)

type PullRequestRepository struct {
	client *ent.Client
}

func NewPullRequestRepository(client *ent.Client) *PullRequestRepository {
	return &PullRequestRepository{client: client}
}
