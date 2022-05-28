package rpc

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestGetSingleConversationRecvMsgOptsLogic(t *testing.T) {
	resp, err := imUserService.GetSingleConversationRecvMsgOpts(
		context.Background(),
		&pb.GetSingleConversationRecvMsgOptsReq{
			UserID:         "1",
			ConversationID: "supergroup_0",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
