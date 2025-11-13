package service_provider

import (
	"AvitoTechTask/internal/adapters/config"

	"fmt"
)

type serverConfig interface {
	Address() string
	Port() int
	Host() string
	EnabledTLS() bool
	DevMode() bool
}

func (s *ServiceProvider) ServerConfig() serverConfig {
	if s.serverConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get http config: %w", err))
		}

		s.serverConfig = cfg
	}

	return s.serverConfig

}
