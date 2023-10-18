package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	defConfigPath = "./config/auth.yml"
)

// Config is a struct that holds all the configuration for the service
type Config struct {
	Server  `yaml:"server"`
	Storage `yaml:"storage"`
}

// Server is a struct that holds all the configuration for the server
type Server struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8081"`
}

type Storage struct {
	DBType   DBType `yaml:"db_type" env-default:"inram"`
	Postgres `yaml:"postgres"`
}

type DBType string

type Postgres struct {
	Host     string `yaml:"host" env-default:"PG_HOST"`
	Port     string `yaml:"port" env-default:"PG_PORT"`
	User     string `yaml:"user" env-default:"PG_USER"`
	Password string `yaml:"pass" env-default:"PG_PASSWORD"`
	Database string `yaml:"database" env-default:"PG_DATABASE"`
}

// MustConfig reads the config from the environment and panics if it fails
func MustConfig() *Config {
	configPath := os.Getenv("AUTH_CONFIG_PATH")
	if configPath == "" {
		configPath = defConfigPath
		log.Printf("AUTH_CONFIG_PATH is not set, using default config: %s", configPath)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	return &cfg
}
