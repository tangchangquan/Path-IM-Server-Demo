package types

import (
	"encoding/json"
	"github.com/showurl/Zero-IM-Server/common/types"
)

const (
	ApplyFriendContentType        int32 = types.NotificationUser2User
	ApplyFriendSuccessContentType int32 = types.NotificationUser2User
	DeleteFriendContentType       int32 = types.NotificationUser2User
)

func NewApplyFriendContent() []byte {
	buf, _ := json.Marshal(&struct {
		Text string
		Type int32
	}{"你有新的好友申请", 1})
	return buf
}
func NewApplyFriendSuccessContent() []byte {
	buf, _ := json.Marshal(&struct {
		Text string
		Type int32
	}{"已同意好友申请", 2})
	return buf
}
func NewDeleteFriendContent() []byte {
	buf, _ := json.Marshal(&struct {
		Text string
		Type int32
	}{"已删除好友", 3})
	return buf

}
