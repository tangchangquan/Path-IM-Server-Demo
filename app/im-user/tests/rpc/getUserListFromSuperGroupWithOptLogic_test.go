package rpc

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestGetUserListFromSuperGroupWithOpt(t *testing.T) {
	resp, err := imUserService.GetUserListFromSuperGroupWithOpt(
		context.Background(),
		&pb.GetUserListFromSuperGroupWithOptReq{
			SuperGroupID: "supergroup_0",
			Opts: []pb.RecvMsgOpt{
				pb.RecvMsgOpt_ReceiveMessage,
				pb.RecvMsgOpt_ReceiveNotNotifyMessage,
			},
			//UserIDList:   []string{"user_0", "user_1"},
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
