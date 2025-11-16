package pull_request

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/types"
	"context"
)

func (s *Service) Merge(ctx context.Context, req *dto.MergePRRequest) (*dto.MergePRResponse, error) {
	pr, err := s.pullRequestsRepo.UpdateStatusAndGetWithRelations(ctx, req.PullRequestID, types.PullRequestStatusMerged)
	if err != nil {
		return nil, err
	}
	return &dto.MergePRResponse{PR: dto.PullRequest{
		PullRequestID:     pr.ID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		Status:            pr.Status.String(),
		AssignedReviewers: pr.AssignedReviewers,
	}}, nil
}
