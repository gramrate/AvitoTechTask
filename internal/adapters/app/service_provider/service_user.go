package service_provider

import "AvitoTechTask/internal/domain/service/user"

func (s *ServiceProvider) UserService() *user.Service {
	if s.userService == nil {
		s.userService = user.NewService(s.UserRepo(), s.PRRepo())
	}

	return s.userService
}
