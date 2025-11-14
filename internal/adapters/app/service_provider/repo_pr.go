package service_provider

import (
	"AvitoTechTask/internal/adapters/repo/postgres/pull_request"
)

func (s *ServiceProvider) PRRepo() *pull_request.Repo {
	if s.prRepo == nil {
		s.prRepo = pull_request.NewRepo(s.DB())
	}

	return s.prRepo
}
