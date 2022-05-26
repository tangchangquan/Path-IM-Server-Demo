package rpc

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestGetUserByUsernameLogic(t *testing.T) {
	resp, err := userService.GetUserByUsername(
		context.Background(),
		&pb.GetUserByUsernameReq{
			Username: "showurl01",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
