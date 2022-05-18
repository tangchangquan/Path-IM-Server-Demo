package rpc

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
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
