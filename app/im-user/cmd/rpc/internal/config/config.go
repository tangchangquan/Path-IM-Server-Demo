package config

import (
	"github.com/showurl/Zero-IM-Server/common/xorm/global"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConfig     RedisConfig
	MysqlConfig     global.MysqlConfig
	TokenSecret     string          `json:",default=zeroimserver"`
	TokenRenewalDay int64           `json:",default=30"` // 用户每次连接websocket 自动续签的的天数 默认30天
	MsgGatewayRpc   discov.EtcdConf `json:",optional"`
	MsgRpc          zrpc.RpcClientConf
}

type RedisConfig struct {
	redis.RedisConf
	DB int
}
