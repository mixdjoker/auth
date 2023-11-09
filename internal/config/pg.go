package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// PgConfiger interface for config Postgres
type PgConfiger interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

// NewPGConfig is a function that returns a new instance of the PgConfiger interface
func NewPGConfig(env EnvConfiger) (PgConfiger, error) {
	args := make(map[string]int)
	args[env.PGHostEnvName()] = len(os.Getenv(env.PGHostEnvName()))
	args[env.PGPortEnvName()] = len(os.Getenv(env.PGPortEnvName()))
	args[env.PGUserEnvName()] = len(os.Getenv(env.PGUserEnvName()))
	args[env.PGPasswordEnvName()] = len(os.Getenv(env.PGPasswordEnvName()))
	args[env.PGDatabaseEnvName()] = len(os.Getenv(env.PGDatabaseEnvName()))
	args[env.PGSSLModeEnvName()] = len(os.Getenv(env.PGSSLModeEnvName()))

	for k, v := range args {
		if v == 0 {
			return nil, errors.Errorf("env variable %s not found", k)
		}
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv(env.PGHostEnvName()),
		os.Getenv(env.PGPortEnvName()),
		os.Getenv(env.PGUserEnvName()),
		os.Getenv(env.PGPasswordEnvName()),
		os.Getenv(env.PGDatabaseEnvName()),
		os.Getenv(env.PGSSLModeEnvName()),
	)

	return &pgConfig{dsn: dsn}, nil
}

// DSN is a method that returns a DSN string
func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
