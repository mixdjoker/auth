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
	reqBuf.WriteString("UpdateRequest {\n")
	fmt.Fprintf(&reqBuf, "\tId: %d,\n\tName: %s\n\tEmail: %s\n\tRole: %s\n",
		req.Id.GetValue(),
		req.User.Name.GetValue(),
		req.User.Email.GetValue(),
		req.User.Role.String(),
	)
	if dLine, ok := ctx.Deadline(); ok {
		fmt.Fprintf(&reqBuf, "\tDeadline: %s\n", dLine.String())
	}
	reqBuf.WriteString("\t}")
	log.Println(color.MagentaString("[gRPC]"), color.BlueString(reqBuf.String()))

	err := i.userService.Update(ctx, dtohelper.ToModelUserFromUpdateRequest(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
