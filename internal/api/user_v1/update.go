package user_v1

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	reqBuf := strings.Builder{}
	reqBuf.WriteString("UpdateRequest {\n")
	fmt.Fprintf(&reqBuf, "\tId: %d,\n\tName: %s\n\tEmail: %s\n\tRole: %s\n",
		req.Id.Value,
		req.User.Name.Value,
		req.User.Email.Value,
		req.User.Role.String(),
	)
	if dLine, ok := ctx.Deadline(); ok {
		fmt.Fprintf(&reqBuf, "\tDeadline: %s\n", dLine.String())
	}
	reqBuf.WriteString("\t}")
	log.Println(color.MagentaString("[gRPC]"), color.BlueString(reqBuf.String()))

	return &emptypb.Empty{}, nil
}
