package rpc

import (
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/imuserservice"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/relationservice"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	imuserConf = zrpc.RpcClientConf{
		Endpoints: []string{
			"192.168.2.77:10240",
		},
		Target:   "",
		App:      "",
		Token:    "",
		NonBlock: true,
		Timeout:  0,
	}
	imUserService = imuserservice.NewImUserService(zrpc.MustNewClient(imuserConf))
	userConf      = zrpc.RpcClientConf{
		Endpoints: []string{
			"192.168.2.77:10260",
		},
		Target:   "",
		App:      "",
		Token:    "",
		NonBlock: true,
		Timeout:  0,
	}
	userService  = userservice.NewUserService(zrpc.MustNewClient(userConf))
	relationConf = zrpc.RpcClientConf{
		Endpoints: []string{
			"192.168.2.77:10270",
		},
		Target:   "",
		App:      "",
		Token:    "",
		NonBlock: true,
		Timeout:  0,
	}
	relationService = relationservice.NewRelationService(zrpc.MustNewClient(relationConf))
)
