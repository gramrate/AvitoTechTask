package service_provider

import (
	"AvitoTechTask/pkg/logger"
)

type ServiceProvider struct {
	serverConfig serverConfig
	loggerConfig loggerConfig

	logger *logger.Logger
}

func New() *ServiceProvider {
	return &ServiceProvider{}
}
