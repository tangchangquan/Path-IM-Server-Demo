package svc

import "github.com/Path-IM/Path-IM-Server-Demo/app/msg-callback/cmd/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
