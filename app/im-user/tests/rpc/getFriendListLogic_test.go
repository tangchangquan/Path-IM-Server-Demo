package rpc

import (
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
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
