package rpc

import (
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestGetFriendListLogic(t *testing.T) {
	resp, err := relationService.GetFriendList(ctx, &pb.GetFriendListReq{
		SelfId: "1",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
