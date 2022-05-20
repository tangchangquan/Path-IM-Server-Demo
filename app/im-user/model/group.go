package model

type Group struct {
	Id     string `gorm:"column:id;primary_key;comment:主键id"`
	Name   string `gorm:"column:name;comment:群组名称"`
	Avatar string `gorm:"column:avatar;comment:群组头像"`
}

func (g *Group) TableName() string {
	return "groups"
}

var (
	DefaultGroup = &Group{
		Id:     "supergroup_0",
		Name:   "炒鸡大群",
		Avatar: "https://go.dev/images/gophers/ladder.svg",
	}
)
