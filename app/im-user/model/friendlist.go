package model

type Friendlist struct {
	UserId string `gorm:"column:user_id;primary_key;type:char(32);comment:主键 用户id;"`
	SelfId string `gorm:"column:self_id;type:char(32);comment:自己id,分表键;"`
	// 加好友时间
	CreatedAt int64  `gorm:"column:created_at;type:int(10);comment:加好友时间 秒级时间戳;"`
	Remark    string `gorm:"column:remark;type:varchar(255);comment:备注;"`
}

func (b *Friendlist) TableName() string {
	return "friendlist_" + b.SelfId
}

type FriendApplyRecord struct {
	Id     string `gorm:"column:id;primary_key;type:char(32);comment:主键"`
	UserId string `gorm:"column:user_id;index;type:char(32);comment:主键 用户id 申请人的id;"`
	SelfId string `gorm:"column:self_id;type:char(32);comment:自己id,分表键;"`
	// 申请时间
	CreatedAt int64 `gorm:"column:created_at;type:int(10);comment:加好友时间 秒级时间戳;index;"`
	// 申请状态
	Status int `gorm:"column:status;type:tinyint(1);comment:申请状态 0:申请中 1:同意 2:拒绝;default:0;"`
	// 申请消息
	Message string `gorm:"column:message;type:varchar(255);comment:申请消息;"`
}

func (f *FriendApplyRecord) TableName() string {
	return "friend_apply_record_" + f.SelfId
}
