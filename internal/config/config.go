package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	_CONFIG_PATH = "./config/auth.yml"
)

// Config is a struct that holds all the configuration for the service
type Config struct {
	Server `yaml:"server"`
}

// Server is a struct that holds all the configuration for the server
type Server struct {
	Host     string `yaml:"host" env-default:"localhost"`
	GRPCPort string `yaml:"grpc_port" env-default:"8081"`
}

// MustConfig reads the config from the environment and panics if it fails
func MustConfig() *Config {
	configPath := os.Getenv("AUTH_CONFIG_PATH")
	if configPath == "" {
		configPath = _CONFIG_PATH
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
