package rpc

import (
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestDeleteFriendLogic(t *testing.T) {
	resp, err := relationService.DeleteFriend(ctx, &pb.DeleteFriendReq{
		SelfId: "1",
		UserId: "2",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
