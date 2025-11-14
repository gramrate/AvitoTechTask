package user

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/types"
	"context"
	"fmt"
)

// GetUserPullRequests возвращает все PR, где пользователь назначен ревьювером.
func (s *Service) GetUserPullRequests(ctx context.Context, req *dto.GetUsersPRRequest) (*dto.GetUsersPRResponse, error) {

	var statusPtr *types.PullRequestStatus
	if req.Status != "" {
		status, err := types.FromString(req.Status)
		if err != nil {
			return nil, fmt.Errorf("invalid status '%s': %w", req.Status, err)
		}
		statusPtr = &status
	}

	pullRequestEntities, err := s.pullRequestsRepo.GetByReviewerIDInTransaction(ctx, req.UserID, statusPtr)
	if err != nil {
		return nil, fmt.Errorf("failed to get user's assigned pull requests: %w", err)
	}

	// Преобразуем ent.PullRequest в dto.PullRequest
	var pullRequests []*dto.PullRequestShort
	for _, prEntity := range pullRequestEntities {
		prDTO := &dto.PullRequestShort{
			PullRequestID:   prEntity.ID,
			PullRequestName: prEntity.PullRequestName,
			AuthorID:        prEntity.AuthorID,
			Status:          prEntity.Status,
		}

		pullRequests = append(pullRequests, prDTO)
	}

	// Создаем response
	response := &dto.GetUsersPRResponse{
		UserID:       req.UserID,
		PullRequests: pullRequests,
	}

	return response, nil
}
