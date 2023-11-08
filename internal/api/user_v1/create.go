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
	// --- Only for logging
	reqBuf := strings.Builder{}
	userBuf := strings.Builder{}
	dlineBuf := strings.Builder{}
	fmt.Fprintf(&userBuf, "{Name: %s, Email: %s, Password: %s, Password confirm: %s, Role: %s}",
		req.User.GetName().GetValue(),
		req.User.GetEmail().GetValue(),
		req.Password.GetValue(),
		req.PasswordConfirm.GetValue(),
		req.User.Role.String(),
	)
	if dLine, ok := ctx.Deadline(); ok {
		fmt.Fprintf(&dlineBuf, "{%s}", dLine.String())
	}
	fmt.Fprintf(&reqBuf, "CreateRequest: {User: %s, Deadline: %s}", userBuf.String(), dlineBuf.String())
	// ---

	log.Println(color.MagentaString("[gRPC]"), color.BlueString(reqBuf.String()))

	id, err := i.userService.Create(ctx, dtohelper.ToModelNewUserFromCreateRequest(req))
	if err != nil {
		if strings.Contains(err.Error(), "ValidationError") {
			log.Println(color.MagentaString("[gRPC]"), color.RedString(fmt.Sprintf("User: %s: %v", userBuf.String(), err)))
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if strings.Contains(err.Error(), "CreationError") {
			log.Println(color.MagentaString("[gRPC]"), color.RedString(fmt.Sprintf("User: %s: %v", userBuf.String(), err)))
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.CreateResponse{
		Id: &wrapperspb.Int64Value{Value: id},
	}, nil
}
