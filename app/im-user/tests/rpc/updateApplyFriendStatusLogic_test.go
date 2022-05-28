package rpc

import (
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestUpdateApplyFriendStatusLogic(t *testing.T) {
	resp, err := relationService.UpdateApplyFriendStatus(ctx, &pb.UpdateApplyFriendStatusReq{
		Status: pb.UpdateApplyFriendStatusReq_AGREE,
		Id:     "ea9ba42393b0c1e2a0260aaf4dead996",
		SelfId: "2",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
