package team

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/pkg/ent"
	"context"
	"fmt"
)

// Create создает команду с указанными пользователями
func (s *Service) Create(ctx context.Context, req *dto.CreateTeamRequest) (*dto.CreateTeamResponse, error) {
	team, err := s.teamRepo.Create(ctx, &ent.Team{
		TeamName: req.TeamName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create team: %w", err)
	}

	// Создаем срез заранее с нужным размером
	members := make([]dto.TeamMember, len(req.Members))

	// Создаем пользователей и привязываем к команде
	for i, member := range req.Members {
		// Создаем пользователя через сервис пользователей
		userResp, err := s.userService.Create(ctx, &dto.CreateUserRequest{
			Username: member.Username,
			TeamID:   team.ID,
			IsActive: member.IsActive,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create user %s: %w", member.Username, err)
		}

		// Записываем по индексу
		members[i] = dto.TeamMember{
			UserID:   userResp.User.UserID,
			Username: userResp.User.Username,
			IsActive: userResp.User.IsActive,
		}
	}

	// Формируем ответ без дополнительного получения команды
	response := &dto.CreateTeamResponse{
		Team: dto.Team{
			TeamName: team.TeamName,
			Members:  members,
		},
	}

	return response, nil
}
