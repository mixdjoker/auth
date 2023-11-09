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
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update implements UserServiceServer.Update
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	reqBuf := strings.Builder{}
	usrBuf := strings.Builder{}
	dlineBuf := strings.Builder{}

	fmt.Fprint(&usrBuf, "{")
	if req.Id != nil {
		fmt.Fprintf(&usrBuf, "Id: %d, ", req.Id.GetValue())
	}
	if req.User.Name != nil {
		fmt.Fprintf(&usrBuf, "Name: %s, ", req.User.Name.GetValue())
	}
	if req.User.Email != nil {
		fmt.Fprintf(&usrBuf, "Email: %s, ", req.User.Email.GetValue())
	}
	if req.User.Role != desc.Role_UNKNOWN {
		fmt.Fprintf(&usrBuf, "Role: %s, ", req.User.Role.String())
	}
	fmt.Fprintf(&usrBuf, "}")

	if dLine, ok := ctx.Deadline(); ok {
		fmt.Fprintf(&dlineBuf, "{%s}", dLine.String())
	}

	fmt.Fprintf(&reqBuf, "UpdateRequest: {User: %s, Deadline: %s}", usrBuf.String(), dlineBuf.String())
	log.Println(color.MagentaString("[gRPC]"), color.BlueString(reqBuf.String()))

	err := i.userService.Update(ctx, dtohelper.ToModelUserFromUpdateRequest(req))
	if err != nil {
		log.Println(color.MagentaString("[gRPC]"), color.RedString(fmt.Sprintf("User: %s: %v", usrBuf.String(), err)))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
