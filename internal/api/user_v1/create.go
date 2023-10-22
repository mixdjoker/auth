package user_v1

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	reqBuf := strings.Builder{}
	reqBuf.WriteString("CreateRequest {\n")
	fmt.Fprintf(&reqBuf, "\tName: %s,\n\tEmail: %s,\n\tPassword: %s,\n\tPassword confirm: %s,\n\tRole: %s\n",
		req.GetName(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetPasswordConfirm(),
		req.GetRole().String(),
	)
	if dLine, ok := ctx.Deadline(); ok {
		fmt.Fprintf(&reqBuf, "\tDeadline: %s\n", dLine.String())
	}
	reqBuf.WriteString("\t}")
	log.Println(color.MagentaString("[gRPC]"),  color.BlueString(reqBuf.String()))

	
	return &desc.CreateResponse{}, nil
}
