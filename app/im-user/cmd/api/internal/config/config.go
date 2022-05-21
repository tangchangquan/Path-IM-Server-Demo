package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ImUserRpc   zrpc.RpcClientConf
	UserRpc     zrpc.RpcClientConf
	RelationRpc zrpc.RpcClientConf
}
