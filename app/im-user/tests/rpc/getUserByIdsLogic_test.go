package rpc

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
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
