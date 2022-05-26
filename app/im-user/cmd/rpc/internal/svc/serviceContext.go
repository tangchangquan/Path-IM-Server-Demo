package svc

import (
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/config"
	"github.com/showurl/Path-IM-Server/app/msg/cmd/rpc/chat"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	msgRpc chat.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
func (s *ServiceContext) MsgRpc() chat.Chat {
	if s.msgRpc == nil {
		s.msgRpc = chat.NewChat(zrpc.MustNewClient(s.Config.MsgRpc))
	}
	return s.msgRpc
}
