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
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	reqBuf := strings.Builder{}
	fmt.Fprintf(&reqBuf, "GetRequest {\n\tId: %d,\n", req.Id.Value)
	if dLine, ok := ctx.Deadline(); ok {
		fmt.Fprintf(&reqBuf, "\tDeadline: %s\n", dLine.String())
	}
	reqBuf.WriteString("\t}")
	log.Println(color.MagentaString("[gRPC]"), color.BlueString(reqBuf.String()))

	user, err := i.userService.Get(ctx, req.Id.GetValue())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	descUserInfo := &desc.UserInfo{
		Id:        &wrapperspb.Int64Value{Value: user.ID},
		CreatedAt: &timestamppb.Timestamp{Seconds: user.CreatedAt.Unix(), Nanos: int32(user.CreatedAt.Nanosecond())},
		User: &desc.User{
			Name:  &wrapperspb.StringValue{Value: user.Name},
			Email: &wrapperspb.StringValue{Value: user.Email},
			Role:  desc.Role(user.Role),
		},
	}

	if user.UpdatedAt != nil {
		descUserInfo.UpdatedAt = &timestamppb.Timestamp{Seconds: user.UpdatedAt.Unix(), Nanos: int32(user.UpdatedAt.Nanosecond())}
	}

	return &desc.GetResponse{
		UserInfo: descUserInfo,
	}, nil
}
