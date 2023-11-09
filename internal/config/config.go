package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	appEnv        = "AUTH_CONFIG_PATH"
	defConfigPath = "./config/auth.yml"
	defEnvPath    = "./.env"
)

// EnvConfiger interface for config environment
type EnvConfiger interface {
	GrpcEnvConfiger
	PgEnvConfiger
}

// GrpcEnvConfiger interface for config GRPC environment
type GrpcEnvConfiger interface {
	GrpcHostEnvName() string
	GrpcPortEnvName() string
}

// PgEnvConfiger interface for config Postgres environment
type PgEnvConfiger interface {
	PGHostEnvName() string
	PGPortEnvName() string
	PGUserEnvName() string
	PGPasswordEnvName() string
	PGDatabaseEnvName() string
	PGSSLModeEnvName() string
}

// EnvConfig is a struct that holds all the configuration for the server
type EnvConfig struct {
	Grpc    `yaml:"grpc"`
	Storage `yaml:"storage"`
	EnvFile `yaml:"env_file"`
}

// Grpc is a struct that holds all the configuration for the GRPC server
type Grpc struct {
	Host string `yaml:"host" env-default:"GRPC_HOST"`
	Port string `yaml:"port" env-default:"GRPC_PORT"`
}

// Storage is a struct that holds all the configuration for the storage
type Storage struct {
	Postgres `yaml:"postgres"`
}

// Postgres is a struct that holds all the configuration for the postgres
type Postgres struct {
	Host     string `yaml:"host" env-default:"PG_HOST"`
	Port     string `yaml:"port" env-default:"PG_PORT"`
	User     string `yaml:"user" env-default:"PG_USER"`
	Password string `yaml:"pass" env-default:"PG_PASSWORD"`
	Database string `yaml:"database" env-default:"PG_DATABASE"`
	SslMode  string `yaml:"sslmode" env-default:"PG_SSLMODE"`
}

// EnvFile is a struct that holds all the configuration for the env file
type EnvFile struct {
	Path string `yaml:"path" env-default:"./.env"`
}

// NewEnvConfig returns a new EnvConfig struct
func NewEnvConfig() (EnvConfiger, error) {
	configPath := os.Getenv(appEnv)
	if configPath == "" {
		configPath = defConfigPath
		log.Printf("%s is not set, using defaults: %s", appEnv, configPath)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg EnvConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	envPath := cfg.EnvFile.Path
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Printf("env file %s does not exist", envPath)
		log.Printf("trying to use default env file: %s", defEnvPath)
		envPath = defEnvPath
	}

	if err := godotenv.Overload(envPath); err != nil {
		log.Printf("failed to load env file: path: %s: %v", envPath, err)
	}

	return &cfg, nil
}

// GrpcHostEnvName returns the GRPC host environment name
func (cfg *EnvConfig) GrpcHostEnvName() string {
	return cfg.Grpc.Host
}

// GrpcPortEnvName returns the GRPC port environment name
func (cfg *EnvConfig) GrpcPortEnvName() string {
	return cfg.Grpc.Port
}

// PGHostEnvName returns the Postgres host environment name
func (cfg *EnvConfig) PGHostEnvName() string {
	return cfg.Storage.Postgres.Host
}

// PGPortEnvName returns the Postgres port environment name
func (cfg *EnvConfig) PGPortEnvName() string {
	return cfg.Storage.Postgres.Port
}

// PGUserEnvName returns the Postgres user environment name
func (cfg *EnvConfig) PGUserEnvName() string {
	return cfg.Storage.Postgres.User
}

// PGPasswordEnvName returns the Postgres password environment name
func (cfg *EnvConfig) PGPasswordEnvName() string {
	return cfg.Storage.Postgres.Password
}

// PGDatabaseEnvName returns the Postgres database environment name
func (cfg *EnvConfig) PGDatabaseEnvName() string {
	return cfg.Storage.Postgres.Database
}

// PGSSLModeEnvName returns the Postgres sslmode environment name
func (cfg *EnvConfig) PGSSLModeEnvName() string {
	return cfg.Storage.Postgres.SslMode
}
