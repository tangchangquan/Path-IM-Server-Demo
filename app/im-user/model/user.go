package model

import (
	"github.com/showurl/Zero-IM-Server/common/utils/encrypt"
	"github.com/showurl/Zero-IM-Server/common/xorm/global"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type User struct {
	Id string `gorm:"column:id;primary_key;comment:主键id;type:char(32);"`
	// 账户
	Username string `gorm:"column:username;comment:用户名;type:char(31);uniqueIndex;"`
	Password string `gorm:"column:password;comment:密码;type:char(64);not null;"`
	// 用户信息显示
	Nickname     string `gorm:"column:nickname;comment:昵称;type:varchar(31);not null;"`
	Sign         string `gorm:"column:sign;comment:个性签名;type:varchar(63);default:'';not null;"`
	Avatar       string `gorm:"column:avatar;comment:头像;type:varchar(255);default:'';not null;"`
	Province     string `gorm:"column:province;index;comment:省;type:varchar(63);default:'';not null;"`
	City         string `gorm:"column:city;index;comment:市;type:varchar(63);default:'';not null;"`
	District     string `gorm:"column:district;index;comment:区;type:varchar(63);default:'';not null;"`
	Birthday     string `gorm:"column:birthday;comment:生日;type:char(10);default:'';not null;"`
	RegisterTime string `gorm:"column:register_time;comment:注册时间;type:char(19);not null;"`
	IsMale       bool   `gorm:"column:is_male;comment:是否是男性;default:1;notnull;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Id == "" {
		u.Id = global.GetID()
	}
	if u.Password != "" {
		u.Password = encrypt.Md5(u.Password)
	}
	return nil
}

func (u *User) GetIdString() string {
	return u.Id
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Value(field string) interface{} {
	t := reflect.TypeOf(*u)
	v := reflect.ValueOf(u).Elem()
	for i := 0; i < t.NumField(); i++ {
		if strings.Contains(t.Field(i).Tag.Get("gorm"), "column:"+field) {
			return v.Field(i).Interface()
		}
	}
	return nil
}
