package rpc

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
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
