package repository

import (
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-transfer/cmd/persistent/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm"
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
