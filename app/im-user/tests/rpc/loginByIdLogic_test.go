package rpc

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestLoginByIdLogic(t *testing.T) {
	resp, err := userService.LoginById(context.Background(), &pb.LoginByIdReq{
		UserId:   "2e5f401f9cfab6b1475c6ebcd901e852",
		Platform: "IOS",
	})

	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}
