package user_v1

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete implements UserServiceServer.Delete
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	reqBuf := strings.Builder{}
	idBuf := strings.Builder{}
	dlineBuf := strings.Builder{}

	fmt.Fprintf(&idBuf, "{Id: %d}", req.Id.GetValue())
	if dLine, ok := ctx.Deadline(); ok {
		fmt.Fprintf(&dlineBuf, "{%s}", dLine.String())
	}
	fmt.Fprintf(&reqBuf, "DeleteRequest: {User: %s, Deadline: %s}", idBuf.String(), dlineBuf.String())

	log.Println(color.MagentaString("[gRPC]"), color.BlueString(reqBuf.String()))

	if err := i.userService.Delete(ctx, req.Id.GetValue()); err != nil {
		log.Println(color.MagentaString("[gRPC]"), color.RedString(fmt.Sprintf("User: %s: %v", idBuf.String(), err)))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
