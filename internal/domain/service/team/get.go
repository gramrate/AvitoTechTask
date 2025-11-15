package team

import (
	"AvitoTechTask/internal/domain/dto"
	"context"
	"fmt"
)

// GetByNameWithMembers возвращает команду с участниками по имени
func (s *Service) GetByNameWithMembers(ctx context.Context, req dto.GetTeamRequest) (*dto.GetTeamResponse, error) {
	team, err := s.teamRepo.GetByNameWithMembers(ctx, req.TeamName)
	if err != nil {
		return nil, fmt.Errorf("failed to get team: %w", err)
	}

	// Конвертируем ent пользователей в TeamMember DTO
	var members []dto.TeamMember
	if team.Edges.Members != nil {
		members = make([]dto.TeamMember, len(team.Edges.Members))
		for i, user := range team.Edges.Members {
			members[i] = dto.TeamMember{
				UserID:   user.ID,
				Username: user.Username,
				IsActive: user.IsActive,
			}
		}
	}

	resp := &dto.GetTeamResponse{
		TeamName: team.TeamName,
		Members:  members,
	}

	return resp, nil
}
