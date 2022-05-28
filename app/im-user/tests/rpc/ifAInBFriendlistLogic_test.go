package rpc

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestIfAInBFriendlist(t *testing.T) {
	resp, err := imUserService.IfAInBFriendList(
		context.Background(),
		&pb.IfAInBFriendListReq{
			AUserID: "a",
			BUserID: "b",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
