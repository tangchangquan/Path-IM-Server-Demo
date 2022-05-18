package model

type SuperGroupConversationRecord struct {
	UserId     string `gorm:"column:user_id;primary_key;type:char(64);comment:主键 用户id;"`
	GroupId    string `gorm:"column:group_id;type:char(64);comment:群组id;"`
	RecvMsgOpt int8   `gorm:"column:recv_msg_opt;index;type:tinyint(1);comment:接收消息选项，0接收并提醒 1屏蔽消息 2接收但不提醒 默认0;default:0;"`
	Remark     string `gorm:"column:remark;type:varchar(255);comment:备注;default:'';"`
}

func (s *SuperGroupConversationRecord) GetIdString() string {
	return s.UserId
}

func (s *SuperGroupConversationRecord) TableName() string {
	return "supergroup_conversation_record_" + s.GroupId
}
