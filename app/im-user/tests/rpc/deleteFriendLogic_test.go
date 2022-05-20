package rpc

import (
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
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
