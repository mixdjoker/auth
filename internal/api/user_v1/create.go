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

// Create implements UserServiceServer.Create
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

	id, err := i.userService.Create(ctx, dtohelper.ToModelUserFromCreateRequest(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.CreateResponse{
		Id: &wrapperspb.Int64Value{Value: id},
	}, nil
}
