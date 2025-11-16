package pull_request

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/types"
	"AvitoTechTask/pkg/ent"
	"context"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, req *dto.CreatePRRequest) (*dto.CreatePRResponse, error) {
	availableReviewers, err := s.pullRequestsRepo.GetAvailableReviewers(ctx, req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get available reviewers: %w", err)
	}
	selectedReviewerIDs := s.selectRandomReviewerIDs(availableReviewers, 2)

	pr := &ent.PullRequest{
		ID:                req.PullRequestID,
		PullRequestName:   req.PullRequestName,
		AuthorID:          req.AuthorID,
		Status:            types.PullRequestStatusOpen,
		AssignedReviewers: selectedReviewerIDs,
		Edges: ent.PullRequestEdges{
			Author:    &ent.User{ID: req.AuthorID},
			Reviewers: nil,
		},
	}

	created, err := s.pullRequestsRepo.Create(ctx, pr)
	if err != nil {
		return nil, fmt.Errorf("failed to create pull request: %w", err)
	}
	return &dto.CreatePRResponse{
		PR: dto.PullRequest{
			PullRequestID:     created.ID,
			PullRequestName:   created.PullRequestName,
			AuthorID:          created.AuthorID,
			Status:            created.Status.String(),
			AssignedReviewers: created.AssignedReviewers,
		},
	}, nil
}

// selectRandomReviewerIDs выбирает случайных ревьюверов из списка
func (s *Service) selectRandomReviewerIDs(reviewers []*ent.User, maxCount int) []uuid.UUID {
	if len(reviewers) == 0 {
		return []uuid.UUID{}
	}

	if len(reviewers) <= maxCount {
		ids := make([]uuid.UUID, len(reviewers))
		for i, reviewer := range reviewers {
			ids[i] = reviewer.ID
		}
		return ids
	}

	selected := make([]uuid.UUID, maxCount)
	indices := rand.Perm(len(reviewers))
	for i := 0; i < maxCount; i++ {
		selected[i] = reviewers[indices[i]].ID
	}
	return selected
}
