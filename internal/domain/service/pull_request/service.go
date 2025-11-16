package pull_request

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/types"
	"AvitoTechTask/pkg/ent"
	"context"

	"github.com/google/uuid"
)

type userService interface {
	UpdateActivity(ctx context.Context, req *dto.SetUserActivityRequest) (*dto.SetUserActivityResponse, error)
}

type pullRequestsRepo interface {
	Create(ctx context.Context, prEntity *ent.PullRequest) (*ent.PullRequest, error)
	GetWithReviewers(ctx context.Context, id uuid.UUID) (*ent.PullRequest, error)
	UpdateStatusAndGetWithRelations(ctx context.Context, prID uuid.UUID, newStatus types.PullRequestStatus) (*ent.PullRequest, error)
	GetAvailableReviewers(ctx context.Context, authorID uuid.UUID) ([]*ent.User, error)
	GetAvailableReviewersFromReviewerTeam(ctx context.Context, oldReviewerID uuid.UUID, prID uuid.UUID) ([]*ent.User, error)
	ReassignReviewer(ctx context.Context, prID uuid.UUID, oldReviewerID uuid.UUID, newReviewerID uuid.UUID) (*ent.PullRequest, error)
}

type Service struct {
	pullRequestsRepo pullRequestsRepo
	userService      userService
}

func NewService(pullRequestsRepo pullRequestsRepo, userService userService) *Service {
	return &Service{
		pullRequestsRepo: pullRequestsRepo,
		userService:      userService,
	}
}
