package user

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/pkg/ent"
	"context"
	"fmt"
)

func (s *Service) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	user := &ent.User{
		Username: req.Username,
		IsActive: req.IsActive,
		Edges: ent.UserEdges{
			Team: &ent.Team{
				ID: req.TeamID,
			},
		},
	}

	createdUser, err := s.userRepo.CreateWithTeam(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	response := &dto.CreateUserResponse{
		UserID:   createdUser.ID,
		Username: createdUser.Username,
		IsActive: createdUser.IsActive,
	}

	if createdUser.Edges.Team != nil {
		response.TeamName = createdUser.Edges.Team.TeamName
	}

	return response, nil
}
