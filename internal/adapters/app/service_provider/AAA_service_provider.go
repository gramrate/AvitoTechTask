package service_provider

import (
	"AvitoTechTask/internal/adapters/controller/api/validator"
	prRepo "AvitoTechTask/internal/adapters/repo/postgres/pull_request"
	teamRepo "AvitoTechTask/internal/adapters/repo/postgres/team"
	userRepo "AvitoTechTask/internal/adapters/repo/postgres/user"
	"AvitoTechTask/internal/domain/service/pull_request"
	"AvitoTechTask/internal/domain/service/team"
	"AvitoTechTask/internal/domain/service/user"
	"AvitoTechTask/pkg/ent"
	"AvitoTechTask/pkg/logger"

	"github.com/go-playground/form"
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
	teamService *team.Service
	prService   *pull_request.Service

	logger      *logger.Logger
	validator   *validator.Validator
	formDecoder *form.Decoder
}

func New() *ServiceProvider {
	return &ServiceProvider{}
}
