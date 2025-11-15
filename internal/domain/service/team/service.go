package team

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/pkg/ent"
	"context"
	"fmt"
)

//// Service интерфейс для работы с командами
//type Service interface {
//	// Create создает команду с указанными пользователями
//	Create(ctx context.Context, req *dto.CreateTeamRequest) (*dto.CreateTeamResponse, error)
//
//	// GetByNameWithMembers возвращает команду с участниками по имени
//	GetByNameWithMembers(ctx context.Context, req dto.GetTeamRequest) (*dto.GetTeamResponse, error)
//}

type teamRepo interface {
	Create(ctx context.Context, teamEntity *ent.Team) (*ent.Team, error)
	GetByNameWithMembers(ctx context.Context, name string) (*ent.Team, error)
}

type userService interface {
	Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
}

type Service struct {
	teamRepo    teamRepo
	userService userService
}

func NewService(teamRepo teamRepo, userService userService) *Service {
	return &Service{teamRepo: teamRepo, userService: userService}
}
