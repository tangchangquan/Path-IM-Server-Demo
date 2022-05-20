package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/model"
	"github.com/showurl/Zero-IM-Server/common/xcache"
	"github.com/showurl/Zero-IM-Server/common/xcache/dc"
	"github.com/showurl/Zero-IM-Server/common/xcache/global"
	"github.com/showurl/Zero-IM-Server/common/xcache/rc"
	"github.com/showurl/Zero-IM-Server/common/xorm"
	"gorm.io/gorm"
)

type Rep struct {
	svcCtx        *svc.ServiceContext
	Cache         redis.UniversalClient
	Mysql         *gorm.DB
	DetailCache   *dc.DbMapping
	RelationCache *rc.RelationMapping
}

func (r *Rep) CheckDefaultSuperGroup() {
	_ = xorm.Insert(r.Mysql, model.DefaultGroup)
}

var rep *Rep

func NewRep(svcCtx *svc.ServiceContext) *Rep {
	if rep != nil {
		return rep
	}
	rep = &Rep{
		svcCtx: svcCtx,
		Cache:  xcache.GetClient(svcCtx.Config.RedisConfig.RedisConf, global.DB(svcCtx.Config.RedisConfig.DB)),
		Mysql:  xorm.GetClient(svcCtx.Config.MysqlConfig),
	}
	rep.DetailCache = dc.GetDbMapping(rep.Cache, rep.Mysql)
	rep.RelationCache = rc.NewRelationMapping(rep.Mysql, rep.Cache)
	rep.CheckDefaultSuperGroup()
	return rep
}
