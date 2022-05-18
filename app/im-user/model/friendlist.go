package model

type Friendlist struct {
	UserId string `gorm:"column:user_id;primary_key;type:char(64);comment:主键 用户id;"`
	SelfId string `gorm:"column:self_id;type:char(64);comment:自己id,分表键;"`
	// 加好友时间
	CreatedAt int64 `gorm:"column:created_at;type:bigint(10);comment:加好友时间 秒级时间戳;"`
}

func (b *Friendlist) TableName() string {
	return "friendlist_" + b.SelfId
}
