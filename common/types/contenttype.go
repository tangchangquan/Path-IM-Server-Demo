package types

const ( ///消息类型
	Typing = 11

	Text    = 101
	Picture = 102
	Voice   = 103
	Video   = 104
	File    = 105
	Custom  = 106

	Location       = 109
	Merger         = 110
	Revoke         = 111
	HasReadReceipt = 112
	Card           = 113
	Quote          = 114
	Common         = 200
	GroupMsg       = 201

	NotificationUser2User = 301 // 用户对用户的通知 比如添加好友请求 打招呼... 不是好友也可以发送 但是拉黑就不能发送了
)

var ContentType2PushContent = map[int64]string{
	Picture:  "[图片]",
	Voice:    "[语音]",
	Video:    "[视频]",
	File:     "[文件]",
	Text:     "你收到了一条文本消息",
	GroupMsg: "你收到一条群聊消息",
	Common:   "你收到一条新消息",
}
