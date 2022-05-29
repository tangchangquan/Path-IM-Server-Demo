package svc

import (
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/api/internal/config"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/api/internal/middleware"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/imuserservice"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/relationservice"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/userservice"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	imUserService   imuserservice.ImUserService
	relationService relationservice.RelationService
	userService     userservice.UserService
	JwtAuth         rest.Middleware
	periodLimit     *limit.PeriodLimit
}

func (s *ServiceContext) PeriodLimit() *limit.PeriodLimit {
	if s.periodLimit == nil {
		s.periodLimit = limit.NewPeriodLimit(
			s.Config.WhiteRateLimit.Seconds,
			s.Config.WhiteRateLimit.Quota,
			newRedis(s.Config.Redis.Host, s.Config.Redis.Pass, s.Config.Redis.Type, s.Config.Redis.Tls),
			"periodlimit:whiteapi:",
			limit.Align(),
		)
	}
	return s.periodLimit
}

func (s *ServiceContext) ImUserService() imuserservice.ImUserService {
	if s.imUserService == nil {
		s.imUserService = imuserservice.NewImUserService(zrpc.MustNewClient(s.Config.ImUserRpc))
	}
	return s.imUserService
}
func (s *ServiceContext) UserService() userservice.UserService {
	if s.userService == nil {
		s.userService = userservice.NewUserService(zrpc.MustNewClient(s.Config.UserRpc))
	}
	return s.userService
}
func (s *ServiceContext) RelationService() relationservice.RelationService {
	if s.relationService == nil {
		s.relationService = relationservice.NewRelationService(zrpc.MustNewClient(s.Config.RelationRpc))
	}
	return s.relationService
}
func NewServiceContext(c config.Config) *ServiceContext {
	sctx := &ServiceContext{
		Config: c,
	}
	sctx.JwtAuth = middleware.NewJwtAuthMiddleware(
		func() imuserservice.ImUserService {
			return sctx.ImUserService()
		}, func() *limit.PeriodLimit {
			return sctx.PeriodLimit()
		}).Handle
	return sctx
}

func newRedis(addr string, password string, typ string, tls bool) *redis.Redis {
	ops := make([]redis.Option, 0)
	if password != "" {
		ops = append(ops, redis.WithPass(password))
	}
	if typ == "cluster" {
		ops = append(ops, redis.Cluster())
	}
	if tls {
		ops = append(ops, redis.WithTLS())
	}
	return redis.New(addr, ops...)
}
