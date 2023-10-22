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

type EnvConfiger interface {
	GrpcEnvConfiger
	PgEnvConfiger
}

type GrpcEnvConfiger interface {
	GrpcHostEnvName() string
	GrpcPortEnvName() string
}

type PgEnvConfiger interface {
	PGHostEnvName() string
	PGPortEnvName() string
	PGUserEnvName() string
	PGPasswordEnvName() string
	PGDatabaseEnvName() string
	PGSSLModeEnvName() string
}

// EnvConfig ...
type EnvConfig struct {
	Grpc    `yaml:"grpc"`
	Storage `yaml:"storage"`
	EnvFile `yaml:"env_file"`
}

// Server is a struct that holds all the configuration for the server
type Grpc struct {
	Host string `yaml:"host" env-default:"GRPC_HOST"`
	Port string `yaml:"port" env-default:"GRPC_PORT"`
}

type Storage struct {
	Postgres `yaml:"postgres"`
}

type Postgres struct {
	Host     string `yaml:"host" env-default:"PG_HOST"`
	Port     string `yaml:"port" env-default:"PG_PORT"`
	User     string `yaml:"user" env-default:"PG_USER"`
	Password string `yaml:"pass" env-default:"PG_PASSWORD"`
	Database string `yaml:"database" env-default:"PG_DATABASE"`
	SslMode  string `yaml:"sslmode" env-default:"PG_SSLMODE"`
}

type EnvFile struct {
	Path string `yaml:"path" env-default:"./.env"`
}

// MustConfig reads the config from the environment and panics if it fails
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

func (cfg *EnvConfig) GrpcHostEnvName() string {
	return cfg.Grpc.Host
}

func (cfg *EnvConfig) GrpcPortEnvName() string {
	return cfg.Grpc.Port
}

func (cfg *EnvConfig) PGHostEnvName() string {
	return cfg.Storage.Postgres.Host
}

func (cfg *EnvConfig) PGPortEnvName() string {
	return cfg.Storage.Postgres.Port
}

func (cfg *EnvConfig) PGUserEnvName() string {
	return cfg.Storage.Postgres.User
}

func (cfg *EnvConfig) PGPasswordEnvName() string {
	return cfg.Storage.Postgres.Password
}

func (cfg *EnvConfig) PGDatabaseEnvName() string {
	return cfg.Storage.Postgres.Database
}

func (cfg *EnvConfig) PGSSLModeEnvName() string {
	return cfg.Storage.Postgres.SslMode
}
