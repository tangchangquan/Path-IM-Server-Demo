package rpc

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
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
