package app

import (
	"context"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/mixdjoker/auth/internal/closer"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	appName = "auth"
)

// App is a struct that holds all the application dependencies
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp is a function that returns a new instance of the App struct
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run is a method that runs the application
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGrpcServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterUser_V1Server(a.grpcServer, a.serviceProvider.UserImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("Trying to start [%s] server on %s", color.BlueString(appName), a.serviceProvider.GrpcConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GrpcConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
