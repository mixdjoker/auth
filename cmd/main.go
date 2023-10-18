package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/mixdjoker/auth/internal/config"
	"github.com/mixdjoker/auth/internal/lib/handler"
	"github.com/mixdjoker/auth/internal/storage/psql"
	"github.com/mixdjoker/auth/internal/storage/ram"
	"github.com/mixdjoker/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.MustConfig()
	aLog := log.New(os.Stdout, color.CyanString("[AUTH] "), log.LstdFlags)

	aLog.Println("Starting auth service...")

	url := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	lis, err := net.Listen("tcp", url)
	if err != nil {
		errStr := fmt.Sprintf("failed to listen: %v", err)
		aLog.Fatalf(color.RedString(errStr))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	var repo handler.UserRepository

	switch cfg.Storage.DBType {
	case "inram":
		repo = ram.NewUserStore()
	case "postgres":
		repo = psql.NewUserStore(cfg)
	default:
		errStr := fmt.Sprintf("unknown db type: %s", cfg.Storage.DBType)
		aLog.Fatalf(color.RedString(errStr))
	}

	rpcSrvV1 := handler.NewUserRPCServerV1(aLog, repo)
	user_v1.RegisterUser_V1Server(s, rpcSrvV1)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.Serve(lis); err != nil {
			errStr := fmt.Sprintf("failed to serve: %v", err)
			aLog.Fatalf(color.RedString(errStr))
		}
	}()

	aLog.Println(color.GreenString("Auth server started successfully "), color.BlueString(url))

	<-done
	s.GracefulStop()
	aLog.Println(color.YellowString("Auth service stopped"))

	repo.Close()
	aLog.Println(color.YellowString("Auth server repo closed"))
}
