package service_provider

import (
	prRepo "AvitoTechTask/internal/adapters/repo/postgres/pull_request"
	teamRepo "AvitoTechTask/internal/adapters/repo/postgres/team"
	userRepo "AvitoTechTask/internal/adapters/repo/postgres/user"
	"AvitoTechTask/internal/domain/service/user"
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/logger"
)

type ServiceProvider struct {
	serverConfig   serverConfig
	loggerConfig   loggerConfig
	postgresConfig postgresConfig

	db *ent.Client

	userRepo *userRepo.Repo
	teamRepo *teamRepo.Repo
	prRepo   *prRepo.Repo

	userService *user.Service

	logger *logger.Logger
}

func New() *ServiceProvider {
	return &ServiceProvider{}
}
