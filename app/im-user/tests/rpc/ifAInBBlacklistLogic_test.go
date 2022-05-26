package rpc

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestIfAInBBlacklist(t *testing.T) {
	resp, err := imUserService.IfAInBBlacklist(
		context.Background(),
		&pb.IfAInBBlacklistReq{
			AUserID: "a",
			BUserID: "b",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
