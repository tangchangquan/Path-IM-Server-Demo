package rpc

import (
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/imuserservice"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	conf = zrpc.RpcClientConf{
		Endpoints: []string{
			"192.168.2.77:10240",
		},
		Target:   "",
		App:      "",
		Token:    "",
		NonBlock: true,
		Timeout:  0,
	}
	imUserService = imuserservice.NewImUserService(zrpc.MustNewClient(conf))
)
