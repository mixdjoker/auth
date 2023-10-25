package config

import (
	"fmt"
	"net"
	"os"
)

// GrpcConfiger interface for grpc config
type GrpcConfiger interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

// NewGrpcConfig returns a new instance of GrpcConfiger
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

// Address returns grpc address
func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
