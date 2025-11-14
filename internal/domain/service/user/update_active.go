package user

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/pkg/ent"
	"context"
	"fmt"
)

func (s *Service) UpdateActivity(ctx context.Context, req *dto.SetUserActivityRequest) (*dto.SetUserActivityResponse, error) {
	user := &ent.User{
		ID:       req.UserID,
		IsActive: req.IsActive,
	}

	updatedUser, err := s.userRepo.UpdateActivity(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	response := &dto.SetUserActivityResponse{
		User: dto.User{
			UserID:   updatedUser.ID,
			Username: updatedUser.Username,
			IsActive: updatedUser.IsActive,
		},
	}

	if updatedUser.Edges.Team != nil {
		response.User.TeamName = updatedUser.Edges.Team.TeamName
	}

	return response, nil
}
