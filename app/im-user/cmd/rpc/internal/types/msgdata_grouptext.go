package types

import (
	"encoding/json"
	chatpb "github.com/Path-IM/Path-IM-Server-Demo/app/msg/cmd/rpc/pb"
	"github.com/Path-IM/Path-IM-Server-Demo/common/types"
	timeUtils "github.com/Path-IM/Path-IM-Server-Demo/common/utils/time"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm/global"
)

func NewGroupTextMsg(
	selfId string,
	groupId string,
	text string,
) *chatpb.MsgData {
	buf, _ := json.Marshal(map[string]string{
		"Text": text,
	})
	return &chatpb.MsgData{
		SendID: selfId,
		//RecvID:           "",
		GroupID:     groupId,
		ClientMsgID: global.GetID(),
		//ServerMsgID:      "",
		//SenderPlatformID: 0,
		//SenderNickname:   "",
		//SenderFaceURL:    "",
		SessionType: types.GroupChatType,
		MsgFrom:     types.SysMsgType,
		ContentType: types.Text,
		Content:     buf,
		//Seq:             0,
		//SendTime:        0,
		CreateTime: timeUtils.Now().UnixMilli(),
		//Status:          0,
		//Options:         nil,
		//OfflinePushInfo: nil,
		//AtUserIDList:    nil,
	}
}
