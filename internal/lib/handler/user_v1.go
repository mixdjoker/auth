package handler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mixdjoker/auth/internal/model"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserRepository interface {
	Create(context.Context, model.User) (int64, error)
	Get(context.Context, int64) (model.User, error)
	Update(context.Context, model.User) error
	Delete(context.Context, int64) error
	Close()
}

// UserRPCServerV1 is a struct that implements the User_V1Server interface
type UserRPCServerV1 struct {
	desc.UnimplementedUser_V1Server
	log *log.Logger
	repo UserRepository
}

// NewUserRPCServerV1 returns a new UserRPCServerV1
func NewUserRPCServerV1(log *log.Logger, r UserRepository) *UserRPCServerV1 {
	return &UserRPCServerV1{
		log: log,
		repo: r,
	}
}

// Create is a method that implements the Create method of the User_V1Server interface
func (s *UserRPCServerV1) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	u := model.User{}

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

	err := validateUserCreateRequest(req)
	if err != nil {
		return nil, err
	}

	u.Name = req.GetName()
	u.Email = req.GetEmail()
	u.Password = req.GetPassword()
	u.Role = int(req.GetRole())

	id, err := s.repo.Create(ctx, u)
	if err != nil {
		s.log.Println(color.RedString("Error Create: %v", err))

		if outErr, ok := errUserEmailExists(err); ok {
			return nil, outErr
		}
		if outErr, ok := errDBConnectCheck(err); ok {
			return nil, outErr
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

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

	u, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		s.log.Println(color.RedString("Error Get: %v", err))

		if outErr, ok := errUserNotFound(err); ok {
			return nil, outErr
		}
		if outErr, ok := errDBConnectCheck(err); ok {
			return nil, outErr
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := desc.GetResponse{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      desc.Role(u.Role),
		CreatedAt: timestamppb.New(time.Unix(u.CreatedAt, 0)),
		UpdatedAt: timestamppb.New(time.Unix(u.UpdatedAt, 0)),
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
	u := model.User{}
	buf.WriteString("Received Update:\n")

	idStr := fmt.Sprintf("\tId: %v\n", req.GetId())
	u.ID = req.GetId()
	buf.WriteString(idStr)

	if req.Name != nil {
		buf.WriteString(fmt.Sprintf("\tName: %v\n", req.GetName().GetValue()))
		u.Name = req.GetName().GetValue()
	}

	if req.Email != nil {
		buf.WriteString(fmt.Sprintf("\tEmail: %v\n", req.GetEmail().GetValue()))
		u.Email = req.GetEmail().GetValue()
	}

	if req.Role != desc.Role_UNKNOWN {
		buf.WriteString(fmt.Sprintf("\tRole: %v\n", req.GetRole()))
		u.Role = int(req.GetRole())
	}

	if dline, ok := ctx.Deadline(); ok {
		log.Println(color.BlueString("Deadline: %v", dline))
	}

	s.log.Println(color.BlueString(buf.String()))

	err := s.repo.Update(ctx, u)
	if err != nil {
		s.log.Println(color.RedString("Error Update: %v", err))

		if outErr, ok := errUserNotFound(err); ok {
			return nil, outErr
		}
		if outErr, ok := errUserEmailExists(err); ok {
			return nil, outErr
		}
		if outErr, ok := errDBConnectCheck(err); ok {
			return nil, outErr
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// Delete is a method that implements the Delete method of the User_V1Server interface
func (s *UserRPCServerV1) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if dline, ok := ctx.Deadline(); ok {
		log.Println(color.BlueString("Deadline: %v", dline))
	}

	log.Println(color.BlueString("Received Delete:\n\tId: %v", req.GetId()))

	err := s.repo.Delete(ctx, req.GetId())
	if err != nil {
		s.log.Println(color.RedString("Error Delete: %v", err))

		if outErr, ok := errUserNotFound(err); ok {
			return nil, outErr
		}
		if outErr, ok := errDBConnectCheck(err); ok {
			return nil, outErr
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func validateUserCreateRequest(req *desc.CreateRequest) error {
	if req.GetPassword() != req.GetPasswordConfirm() {
		return status.Error(codes.InvalidArgument, "passwords do not match")
	}

	if req.GetRole() == desc.Role_UNKNOWN {
		return status.Error(codes.InvalidArgument, "role is unknown")
	}

	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is empty")
	}

	return nil
}

func errDBConnectCheck(err error) (error, bool) {
	if strings.HasPrefix(err.Error(), "failed to connect to") {
		return status.Error(codes.Internal, "failed to connect to database"), true
	}

	return err, false
}

func errUserNotFound(err error) (error, bool) {
	if strings.HasPrefix(err.Error(), "user with id") && strings.HasSuffix(err.Error(), "not found") {
		return status.Error(codes.NotFound, err.Error()), true
	}

	return err, false
}

func errUserEmailExists(err error) (error, bool) {
	if strings.HasPrefix(err.Error(), "user with email") && strings.HasSuffix(err.Error(), "already exists") {
		return status.Error(codes.InvalidArgument, err.Error()), true
	}

	return err, false
}
