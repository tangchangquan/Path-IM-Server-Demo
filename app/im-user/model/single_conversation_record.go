package model

type SingleConversationRecord struct {
	ConversationId string `gorm:"column:conversation_id;primary_key;type:varchar(127);comment:主键 会话id;"`
	UserId         string `gorm:"column:user_id;type:char(64);comment:用户id;"`
	RecvMsgOpt     int8   `gorm:"column:recv_msg_opt;index;type:tinyint(1);comment:接收消息选项，0接收并提醒 1屏蔽消息 2接收但不提醒 默认0;default:0;"`
	Remark         string `gorm:"column:remark;type:varchar(255);comment:备注;default:'';"`
}

func (s *SingleConversationRecord) GetIdString() string {
	return s.ConversationId
}

func (s *SingleConversationRecord) TableName() string {
	return "single_conversation_record_" + s.UserId
}
