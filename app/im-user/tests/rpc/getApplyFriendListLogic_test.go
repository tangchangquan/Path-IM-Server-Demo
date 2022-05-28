package rpc

import (
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestGetApplyFriendListLogic(t *testing.T) {
	resp, err := relationService.GetApplyFriendList(ctx, &pb.GetApplyFriendListReq{
		SelfId: "2",
		PageReq: &pb.PageReq{
			Page:     1,
			PageSize: 20,
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
