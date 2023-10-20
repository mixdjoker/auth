package app

import "google.golang.org/grpc"


type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

