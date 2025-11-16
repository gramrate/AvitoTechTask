package service_provider

import (
	"AvitoTechTask/internal/domain/service/pull_request"
)

func (s *ServiceProvider) PRService() *pull_request.Service {
	if s.prService == nil {
		s.prService = pull_request.NewService(s.PRRepo(), s.UserService())
	}

	return s.prService
}
