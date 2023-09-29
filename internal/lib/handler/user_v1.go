package handler

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"

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

// Create is a method that implements the Create method of the User_V1Server interface
func (h *UserHandlerV1) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	name := req.GetName()
	email := req.GetEmail()
	pass := req.GetPassword()
	passConfirm := req.GetPasswordConfirm()
	role := req.GetRole()

	if dline, ok := ctx.Deadline(); ok {
		h.log.Println(color.BlueString("Deadline: %v", dline))
	}

	resStr := fmt.Sprintf("Received Create:\n\tName: %v,\n\tEmail: %v,\n\tPassword: %v,\n\tPassword confirm: %v,\n\tRole: %v\n", name, email, pass, passConfirm, role)
	h.log.Println(color.BlueString(resStr))

	randInt64, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<63-1))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	id := randInt64.Int64()

	respStr := fmt.Sprintf("Response Create:\n\tId: %v\n", id)
	h.log.Println(color.GreenString(respStr))

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// Get is a method that implements the Get method of the User_V1Server interface
func (h *UserHandlerV1) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	id := req.GetId()

	if dline, ok := ctx.Deadline(); ok {
		h.log.Println(color.BlueString("Deadline: %v", dline))
	}

	h.log.Println(color.BlueString("Received Get:\n\tId: %v", id))

	role := gofakeit.RandString([]string{"ADMIN", "USER"})

	resp := desc.GetResponse{
		Id:        id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role(desc.Role_value[role]),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}

	rStr := fmt.Sprintf("Response Get:\n\tId: %v,\n\tName: %v,\n\tEmail: %v,\n\tRole: %v,\n\tCreatedAt: %v,\n\tUpdatedAt: %v\n",
		resp.Id,
		resp.Name,
		resp.Email,
		resp.Role,
		resp.CreatedAt,
		resp.UpdatedAt)

	h.log.Println(color.GreenString(rStr))

	return &resp, nil
}

// Update is a method that implements the Update method of the User_V1Server interface
func (h *UserHandlerV1) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	buf := strings.Builder{}

	buf.WriteString("Received Update:\n")

	idStr := fmt.Sprintf("\tId: %v\n", req.GetId())
	buf.WriteString(idStr)

	if req.Name != nil {
		name := req.GetName().GetValue()
		nameStr := fmt.Sprintf("\tName: %v\n", name)
		buf.WriteString(nameStr)
	}

	if req.Email != nil {
		email := req.GetEmail().GetValue()
		emailStr := fmt.Sprintf("\tEmail: %v\n", email)
		buf.WriteString(emailStr)
	}
	if dline, ok := ctx.Deadline(); ok {
		log.Println(color.BlueString("Deadline: %v", dline))
	}

	h.log.Println(color.BlueString(buf.String()))

	return &emptypb.Empty{}, nil
}

// Delete is a method that implements the Delete method of the User_V1Server interface
func (h *UserHandlerV1) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()

	if dline, ok := ctx.Deadline(); ok {
		log.Println(color.BlueString("Deadline: %v", dline))
	}

	log.Println(color.BlueString("Received Delete:\n\tId: %v", id))

	return &emptypb.Empty{}, nil
}
