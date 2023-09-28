package handler

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/fatih/color"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserHandlerV1 is a struct that implements the UserHandlerV1 interface
type UserHandlerV1 struct {
	desc.UnimplementedUser_V1Server
	log *log.Logger
}

// NewUserHandlerV1 returns a new UserHandlerV1
func NewUserHandlerV1(log *log.Logger) *UserHandlerV1 {
	return &UserHandlerV1{
		log: log,
	}
}

// Create is a method that implements the Create method of the UserHandlerV1 interface
func (h *UserHandlerV1) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	name := req.GetName()
	email := req.GetEmail()
	pass := req.GetPassword()
	passConfirm := req.GetPasswordConfirm()
	role := req.GetRole()

	if dline, ok := ctx.Deadline(); ok {
		h.log.Println(color.BlueString("Deadline: %v", dline))
	}

	resStr := fmt.Sprintf("Received:\n\tName: %v,\n\temail: %v,\n\tPassword: %v,\n\tPassword confirm: %v,\n\tRole: %v\n", name, email, pass, passConfirm, role)
	h.log.Println(color.BlueString(resStr))

	randInt64, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<63-1))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.CreateResponse{
		Id: randInt64.Int64(),
	}, nil
}
