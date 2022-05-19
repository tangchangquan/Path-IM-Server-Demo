package rpc

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
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
