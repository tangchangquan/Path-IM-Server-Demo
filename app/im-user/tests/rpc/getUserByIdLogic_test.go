package rpc

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestGetUserByIdLogic(t *testing.T) {
	resp, err := userService.GetUserById(
		context.Background(),
		&pb.GetUserByIdReq{
			UserId: "2e5f401f9cfab6b1475c6ebcd901e852",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
