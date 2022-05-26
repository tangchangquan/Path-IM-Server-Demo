package rpc

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
	"time"
)

func TestInsertUserLogic(t *testing.T) {
	resp, err := userService.InsertUser(
		context.Background(),
		&pb.InsertUserReq{
			User: &pb.User{
				Id:           "1",
				Username:     "user01",
				Password:     "123456",
				Nickname:     "第user01个注册的人",
				Sign:         "第user01个注册的人，吃瓜",
				Avatar:       "https://go.dev/images/gophers/pilot-bust.svg",
				Province:     "北京市",
				City:         "北京市",
				District:     "通州区",
				Birthday:     "1999-03-06",
				RegisterTime: time.Now().Format("2006-01-02 15:04:05"),
				IsMale:       true,
			},
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
