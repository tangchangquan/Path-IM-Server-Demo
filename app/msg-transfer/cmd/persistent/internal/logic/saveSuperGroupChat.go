package logic

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/msg-transfer/model"
	chatpb "github.com/showurl/Path-IM-Server/app/msg/cmd/rpc/pb"
	xormerr "github.com/showurl/Path-IM-Server/common/xorm/err"
	"github.com/showurl/Path-IM-Server/common/xtrace"
)

func (l *MsgTransferPersistentOnlineLogic) saveSuperGroupChat(ctx context.Context, key string, c *chatpb.MsgDataToMQ) error {
	if c == nil || c.MsgData == nil || c.MsgData.ServerMsgID == "" || c.MsgData.GroupID == "" {
		l.Error("saveGroupChat error, c is nil or c.MsgData is nil")
		return nil
	}
	chat := &model.GroupChat{
		ServerMsgID:      c.MsgData.ServerMsgID,
		SendID:           c.MsgData.SendID,
		RecvID:           c.MsgData.RecvID,
		GroupID:          c.MsgData.GroupID,
		ClientMsgID:      c.MsgData.ClientMsgID,
		SenderPlatformID: c.MsgData.SenderPlatformID,
		SenderNickname:   c.MsgData.SenderNickname,
		SenderFaceURL:    c.MsgData.SenderFaceURL,
		SessionType:      c.MsgData.SessionType,
		MsgFrom:          c.MsgData.MsgFrom,
		ContentType:      c.MsgData.ContentType,
		Content:          string(c.MsgData.Content),
		Seq:              c.MsgData.Seq,
		SendTime:         c.MsgData.SendTime,
		CreateTime:       c.MsgData.CreateTime,
		OfflinePushInfo:  model.NewOfflinePushInfo(c.MsgData.OfflinePushInfo),
		AtUserIDList:     c.MsgData.AtUserIDList,
		Options:          c.MsgData.Options,
	}
	var err error
	xtrace.StartFuncSpan(ctx, "saveGroupChat", func(ctx context.Context) {
		err = chat.Insert(l.rep.Mysql)
		if err != nil {
			if xormerr.DuplicateError(err) {
				err = nil
			}
		}
	})
	if err != nil {
		l.Error("saveGroupChat error", err)
		return err
	}
	return nil
}
