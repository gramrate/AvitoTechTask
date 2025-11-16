package service_provider

import (
	"AvitoTechTask/internal/domain/service/team"
)

func (s *ServiceProvider) TeamService() *team.Service {
	if s.teamService == nil {
		s.teamService = team.NewService(s.TeamRepo(), s.UserService())
	}

	return s.teamService
}
