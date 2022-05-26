package rpc

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
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
