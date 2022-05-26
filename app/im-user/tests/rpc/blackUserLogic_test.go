package rpc

import (
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestBlackUserLogic(t *testing.T) {
	resp, err := relationService.BlackUser(ctx, &pb.BlackUserReq{
		SelfId: "1",
		UserId: "2",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}

func TestDeleteBlackLogic(t *testing.T) {
	resp, err := relationService.DeleteBlack(ctx, &pb.DeleteBlackReq{
		SelfId: "1",
		UserId: "2",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}

func TestGetBlackListLogic(t *testing.T) {
	resp, err := relationService.GetBlackList(ctx, &pb.GetBlackListReq{
		SelfId: "1",
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
