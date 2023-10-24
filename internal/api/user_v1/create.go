package user_v1

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/mixdjoker/auth/internal/dtohelper"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	reqBuf := strings.Builder{}
	reqBuf.WriteString("CreateRequest {\n")
	fmt.Fprintf(&reqBuf, "\tName: %s,\n\tEmail: %s,\n\tPassword: %s,\n\tPassword confirm: %s,\n\tRole: %s\n",
		req.User.GetName().GetValue(),
		req.User.GetEmail().GetValue(),
		req.Password.GetValue(),
		req.PasswordConfirm.GetValue(),
		req.User.Role.String(),
	)
	if dLine, ok := ctx.Deadline(); ok {
		fmt.Fprintf(&reqBuf, "\tDeadline: %s\n", dLine.String())
	}
	reqBuf.WriteString("\t}")
	log.Println(color.MagentaString("[gRPC]"), color.BlueString(reqBuf.String()))

	if err := validateUserCreateRequest(req); err != nil {
		return nil, err
	}

	id, err := i.userService.Create(ctx, dtohelper.ToModelUserFromCreateRequest(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.CreateResponse{
		Id: &wrapperspb.Int64Value{Value: id},
	}, nil
}

func validateUserCreateRequest(req *desc.CreateRequest) error {
	if req.User == nil {
		return status.Errorf(codes.InvalidArgument, "User is required")
	}

	if req.User.Name == nil {
		return status.Errorf(codes.InvalidArgument, "Name is required")
	}

	if req.User.Email == nil {
		return status.Errorf(codes.InvalidArgument, "Email is required")
	}

	if req.Password == nil {
		return status.Errorf(codes.InvalidArgument, "Password is required")
	}

	if req.PasswordConfirm == nil {
		return status.Errorf(codes.InvalidArgument, "Password confirm is required")
	}

	if req.Password.GetValue() != req.PasswordConfirm.GetValue() {
		return status.Errorf(codes.InvalidArgument, "Password and password confirm must be equal")
	}

	if req.User.Role == desc.Role_UNKNOWN {
		return status.Errorf(codes.InvalidArgument, "Role is required")
	}

	return nil
}
