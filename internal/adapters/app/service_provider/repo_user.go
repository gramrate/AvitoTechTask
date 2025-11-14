package service_provider

import (
	"AvitoTechTask/internal/adapters/repo/postgres/user"
)

func (s *ServiceProvider) UserRepo() *user.Repo {
	if s.userRepo == nil {
		s.userRepo = user.NewRepo(s.DB())
	}

	return s.userRepo
}
