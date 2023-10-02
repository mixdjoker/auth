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
	"github.com/mixdjoker/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.MustConfig()
	aLog := log.New(os.Stdout, color.CyanString("[INFO] "), log.LstdFlags)

	aLog.Println("Starting auth service...")

	url := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.GRPCPort)

	lis, err := net.Listen("tcp", url)
	if err != nil {
		errStr := fmt.Sprintf("failed to listen: %v", err)
		aLog.Fatalf(color.RedString(errStr))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	rpcSrvV1 := handler.NewUserRPCServerV1(aLog)
	user_v1.RegisterUser_V1Server(s, rpcSrvV1)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.Serve(lis); err != nil {
			errStr := fmt.Sprintf("failed to serve: %v", err)
			aLog.Fatalf(color.RedString(errStr))
		}
	}()

	aLog.Println(color.GreenString("Auth service started successfully "), color.BlueString(url))

	<-done
	s.GracefulStop()
	aLog.Println(color.YellowString("Auth service stopped"))
}
