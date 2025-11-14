package service_provider

import (
	"AvitoTechTask/internal/adapters/config"
	"fmt"
)

type postgresConfig interface {
	DSN() string
}

func (s *ServiceProvider) PostgresConfig() postgresConfig {
	if s.postgresConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get http config: %w", err))
		}

		s.postgresConfig = cfg
	}

	return s.postgresConfig

}
