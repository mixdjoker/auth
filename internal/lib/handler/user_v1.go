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

// UserRPCServerV1 is a struct that implements the User_V1Server interface
type UserRPCServerV1 struct {
	desc.UnimplementedUser_V1Server
	log *log.Logger
}

// NewUserRPCServerV1 returns a new UserRPCServerV1
func NewUserRPCServerV1(log *log.Logger) *UserRPCServerV1 {
	return &UserRPCServerV1{
		log: log,
	}
}

// Create is a method that implements the Create method of the User_V1Server interface
func (s *UserRPCServerV1) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	resStr := fmt.Sprintf("Received Create:\n\tName: %v,\n\tEmail: %v,\n\tPassword: %v,\n\tPassword confirm: %v,\n\tRole: %v\n",
		req.GetName(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetPasswordConfirm(),
		req.GetRole(),
	)
	s.log.Println(color.BlueString(resStr))

	if dline, ok := ctx.Deadline(); ok {
		s.log.Println(color.BlueString("Deadline: %v", dline))
	}

	randInt64, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<62))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	id := randInt64.Int64()

	respStr := fmt.Sprintf("Response Create:\n\tId: %v\n", id)
	s.log.Println(color.GreenString(respStr))

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// Get is a method that implements the Get method of the User_V1Server interface
func (s *UserRPCServerV1) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	s.log.Println(color.BlueString("Received Get:\n\tId: %v", req.GetId()))

	if dline, ok := ctx.Deadline(); ok {
		s.log.Println(color.BlueString("Deadline: %v", dline))
	}

	role := gofakeit.RandString([]string{"ADMIN", "USER"})
	resp := desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.BeerName(),
		Email:     gofakeit.Email(),
		Role:      desc.Role(desc.Role_value[role]),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}

	respStr := fmt.Sprintf("Response Get:\n\tId: %v,\n\tName: %v,\n\tEmail: %v,\n\tRole: %v,\n\tCreatedAt: %v,\n\tUpdatedAt: %v\n",
		resp.Id,
		resp.Name,
		resp.Email,
		resp.Role,
		resp.CreatedAt,
		resp.UpdatedAt,
	)

	s.log.Println(color.GreenString(respStr))

	return &resp, nil
}

// Update is a method that implements the Update method of the User_V1Server interface
func (s *UserRPCServerV1) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	buf := strings.Builder{}
	buf.WriteString("Received Update:\n")
	idStr := fmt.Sprintf("\tId: %v\n", req.GetId())
	buf.WriteString(idStr)

	if req.Name != nil {
		buf.WriteString(fmt.Sprintf("\tName: %v\n", req.GetName().GetValue()))
	}

	if req.Email != nil {
		buf.WriteString(fmt.Sprintf("\tEmail: %v\n", req.GetEmail().GetValue()))
	}

	if req.Role != desc.Role_UNKNOWN {
		buf.WriteString(fmt.Sprintf("\tRole: %v\n", req.GetRole()))
	}

	if dline, ok := ctx.Deadline(); ok {
		log.Println(color.BlueString("Deadline: %v", dline))
	}

	s.log.Println(color.BlueString(buf.String()))

	return &emptypb.Empty{}, nil
}

// Delete is a method that implements the Delete method of the User_V1Server interface
func (s *UserRPCServerV1) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if dline, ok := ctx.Deadline(); ok {
		log.Println(color.BlueString("Deadline: %v", dline))
	}

	log.Println(color.BlueString("Received Delete:\n\tId: %v", req.GetId()))

	return &emptypb.Empty{}, nil
}
