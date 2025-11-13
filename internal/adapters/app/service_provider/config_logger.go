package service_provider

import (
	"AvitoTechTask/internal/adapters/config"
	"fmt"
	"time"
)

type loggerConfig interface {
	Debug() bool
	LogToFile() bool
	LogsDir() string
	TimeLocation() *time.Location
}

func (s *ServiceProvider) LoggerConfig() loggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get server config: %w", err))
		}
		s.loggerConfig = cfg
	}

	return s.loggerConfig
}
