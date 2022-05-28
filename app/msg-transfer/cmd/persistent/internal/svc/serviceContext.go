package svc

import (
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-transfer/cmd/persistent/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
