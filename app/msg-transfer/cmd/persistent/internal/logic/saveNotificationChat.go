package logic

import (
	"context"
	chatpb "github.com/Path-IM/Path-IM-Server-Demo/app/msg/cmd/rpc/pb"
)

func (l *MsgTransferPersistentOnlineLogic) saveNotificationChat(ctx context.Context, key string, c *chatpb.MsgDataToMQ) error {
	// 不存储通知消息
	return nil
}
