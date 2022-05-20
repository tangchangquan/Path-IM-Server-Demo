package types

import (
	chatpb "github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/pb"
	"github.com/showurl/Zero-IM-Server/common/types"
	"github.com/showurl/Zero-IM-Server/common/utils"
	timeUtils "github.com/showurl/Zero-IM-Server/common/utils/time"
	"github.com/showurl/Zero-IM-Server/common/xorm/global"
)

func NewSingleChatMsgNotification(
	sendId string,
	receiveId string,
	content []byte,
	contentType int32,
	offlinePushInfo *chatpb.OfflinePushInfo,
) *chatpb.MsgData {
	options := map[string]bool{}
	utils.SetSwitchFromOptions(options, types.IsHistory, false)
	utils.SetSwitchFromOptions(options, types.IsPersistent, false)
	utils.SetSwitchFromOptions(options, types.IsUnreadCount, false)
	utils.SetSwitchFromOptions(options, types.IsConversationUpdate, false)
	utils.SetSwitchFromOptions(options, types.IsSenderConversationUpdate, false)
	utils.SetSwitchFromOptions(options, types.IsSenderSync, false)
	utils.SetSwitchFromOptions(options, types.IsOfflinePush, true)
	return &chatpb.MsgData{
		SendID: sendId,
		RecvID: receiveId,
		//GroupID:          "",
		ClientMsgID: global.GetID(),
		//ServerMsgID:      "",
		//SenderPlatformID: 0,
		//SenderNickname:   "",
		//SenderFaceURL:    "",
		SessionType: types.NotificationChatType,
		MsgFrom:     types.SysMsgType,
		ContentType: contentType,
		Content:     content,
		//Seq:              0,
		//SendTime:         0,
		CreateTime: timeUtils.Now().UnixMilli(),
		//Status:           0,
		Options:         options,
		OfflinePushInfo: offlinePushInfo,
		//AtUserIDList:     nil,
	}
}
