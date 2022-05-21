package model

type Blacklist struct {
	UserId string `gorm:"column:user_id;primary_key;type:char(32);comment:主键 用户id;"`
	SelfId string `gorm:"column:self_id;type:char(32);comment:自己id,分表键;"`
	// 拉黑时间
	CreatedAt int64 `gorm:"column:created_at;type:int(10);comment:拉黑时间 秒级时间戳;index;"`
}

func (b *Blacklist) TableName() string {
	return "blacklist_" + b.SelfId
}
