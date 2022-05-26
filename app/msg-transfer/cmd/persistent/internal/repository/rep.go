package repository

import (
	"github.com/showurl/Path-IM-Server/app/msg-transfer/cmd/persistent/internal/svc"
	"github.com/showurl/Path-IM-Server/common/xorm"
	"gorm.io/gorm"
)

type Rep struct {
	svcCtx *svc.ServiceContext
	Mysql  *gorm.DB
}

var rep *Rep

func NewRep(svcCtx *svc.ServiceContext) *Rep {
	if rep != nil {
		return rep
	}
	rep = &Rep{
		svcCtx: svcCtx,
	}
	rep.Mysql = xorm.GetClient(svcCtx.Config.MysqlConfig)
	err := rep.Mysql.AutoMigrate()
	if err != nil {
		panic(err)
	}
	return rep
}
