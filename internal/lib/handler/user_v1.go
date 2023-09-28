package handler

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

// Get is a method that implements the Get method of the UserHandlerV1 interface
func (h *UserHandlerV1) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	id := req.GetId()

	if dline, ok := ctx.Deadline(); ok {
		h.log.Println(color.BlueString("Deadline: %v", dline))
	}

	h.log.Println(color.BlueString("Received:\n\tId: %v", id))

	role := gofakeit.RandString([]string{"ADMIN", "USER"})

	resp := desc.GetResponse{
		Id:        id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role(desc.Role_value[role]),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}

	rStr := fmt.Sprintf("Response:\n\tId: %v,\n\tName: %v,\n\temail: %v,\n\tRole: %v,\n\tCreatedAt: %v,\n\tUpdatedAt: %v\n",
		resp.Id,
		resp.Name,
		resp.Email,
		resp.Role,
		resp.CreatedAt,
		resp.UpdatedAt)

	h.log.Println(color.GreenString(rStr))

	return &resp, nil
}

// Update is a method that implements the Update method of the UserHandlerV1 interface
func Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// Delete is a method that implements the Delete method of the UserHandlerV1 interface
func Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
