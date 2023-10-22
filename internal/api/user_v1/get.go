package user_v1

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	reqBuf := strings.Builder{}
	fmt.Fprintf(&reqBuf, "GetRequest {\n\tId: %d,\n", req.GetId())
	if dLine, ok := ctx.Deadline(); ok {
		fmt.Fprintf(&reqBuf, "\tDeadline: %s\n", dLine.String())
	}
	reqBuf.WriteString("\t}")
	log.Println(color.MagentaString("[gRPC]"), color.BlueString(reqBuf.String()))

	return &desc.GetResponse{}, nil
}
