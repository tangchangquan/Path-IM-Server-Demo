package rpc

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"testing"
)

var ctx = context.Background()

func TestApplyFriendLogic(t *testing.T) {

	resp, err := relationService.ApplyFriend(ctx, &pb.ApplyFriendReq{
		ApplyUserId: "1",
		UserId:      "2",
		Message:     "1对2的好友申请",
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp:%+v", resp)
}
