package rpc

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
	"time"
)

func TestUpdateUserLogic(t *testing.T) {
	resp, err := userService.UpdateUser(
		context.Background(),
		&pb.UpdateUserReq{
			User: &pb.User{
				Id:           "2e5f401f9cfab6b1475c6ebcd901e852",
				Username:     "showurl-01",
				Password:     "123456",
				Nickname:     "第1个注册的人",
				Sign:         "showurl-01",
				Avatar:       "https://go.dev/images/gophers/pilot-bust.svg",
				Province:     "北京市",
				City:         "北京市",
				District:     "通州区",
				Birthday:     "1999-03-06",
				RegisterTime: time.Now().Format("2006-01-02 15:04:05"),
				IsMale:       true,
			},
			Fields: []string{"username", "sign", "is_male"},
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
