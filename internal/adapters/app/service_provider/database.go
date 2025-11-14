package service_provider

import (
	"AvitoTechTask/pkg/closer"
	"AvitoTechTask/pkg/ent"
	"context"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

func (s *ServiceProvider) DB() *ent.Client {
	if s.db == nil {
		s.Logger().Debugf("Connecting to database (dsn=%s)", s.PostgresConfig().DSN())
		db, err := ent.Open(dialect.Postgres, s.PostgresConfig().DSN())
		if err != nil {
			s.Logger().Panicf("failed to open database: %v", err)
		}
		client := db

		loggerCfg := s.LoggerConfig()
		if loggerCfg.Debug() {
			client = client.Debug()
		}

		if errMigrate := client.Schema.Create(
			context.Background(),
		); errMigrate != nil {
			s.Logger().Panicf("Failed to run migrations: %v", errMigrate)
		}

		closer.Add(client.Close)
		s.db = client
	}

	return s.db
}
