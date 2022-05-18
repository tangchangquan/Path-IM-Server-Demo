package model

type User struct {
	Id string `gorm:"column:id;primary_key;comment:主键id;type:char(64);"`
	// 账户
	Username string `gorm:"column:username;comment:用户名;type:varchar(31);index,unique;not null;"`
	Password string `gorm:"column:password;comment:密码;type:char(64);not null;"`
	// 用户信息显示
	Nickname string `gorm:"column:nickname;comment:昵称;type:varchar(31);not null;"`
	Sign     string `gorm:"column:sign;comment:个性签名;type:varchar(63);default:'';not null;"`
	Avatar   string `gorm:"column:avatar;comment:头像;type:varchar(255);default:'';not null;"`
	Province string `gorm:"column:province;index;comment:省;type:varchar(63);default:'';not null;"`
	City     string `gorm:"column:city;index;comment:市;type:varchar(63);default:'';not null;"`
	District string `gorm:"column:district;index;comment:区;type:varchar(63);default:'';not null;"`
	Birthday string `gorm:"column:birthday;comment:生日;type:char(10);default:'';not null;"`
	IsMale   bool   `gorm:"column:is_male;comment:是否是男性;default:1;notnull;"`
}

func (u *User) TableName() string {
	return "users"
}
