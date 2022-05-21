package rpc

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestGetUserByIdsLogic(t *testing.T) {
	resp, err := userService.GetUserByIds(
		context.Background(),
		&pb.GetUserByIdsReq{
			UserIds: []string{
				"2e5f401f9cfab6b1475c6ebcd901e852",
			},
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
