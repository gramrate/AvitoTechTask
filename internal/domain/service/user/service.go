package user

import (
	"AvitoTechTask/internal/domain/types"
	"AvitoTechTask/pkg/ent"
	"context"

	"github.com/google/uuid"
)

type userRepo interface {
	CreateWithTeam(ctx context.Context, userEntity *ent.User) (*ent.User, error)
	UpdateActivity(ctx context.Context, userEntity *ent.User) (*ent.User, error)
}

type pullRequestsRepo interface {
	GetByReviewerIDInTransaction(ctx context.Context, reviewerID uuid.UUID, status *types.PullRequestStatus) ([]*ent.PullRequest, error)
}

type Service struct {
	userRepo         userRepo
	pullRequestsRepo pullRequestsRepo
}

func NewService(userRepo userRepo, pullRequestsRepo pullRequestsRepo) *Service {
	return &Service{
		userRepo:         userRepo,
		pullRequestsRepo: pullRequestsRepo,
	}
}
