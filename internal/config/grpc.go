package config

import (
	"fmt"
	"net"
	"os"
)

type GrpcConfiger interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGrpcConfig(env EnvConfiger) (GrpcConfiger, error) {
	host := os.Getenv(env.GrpcHostEnvName())
	if len(host) == 0 {
		return nil, fmt.Errorf("grpc host not found")
	}

	port := os.Getenv(env.GrpcPortEnvName())
	if len(port) == 0 {
		return nil, fmt.Errorf("grpc port not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
