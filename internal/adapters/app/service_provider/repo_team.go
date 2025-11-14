package service_provider

import (
	"AvitoTechTask/internal/adapters/repo/postgres/team"
)

func (s *ServiceProvider) TeamRepo() *team.Repo {
	if s.teamRepo == nil {
		s.teamRepo = team.NewRepo(s.DB())
	}

	return s.teamRepo
}
