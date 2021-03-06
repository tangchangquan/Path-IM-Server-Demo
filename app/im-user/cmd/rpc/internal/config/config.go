package config

import (
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm/global"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	TokenRateLimiter RateLimitConfig
	zrpc.RpcServerConf
	RedisConfig            RedisConfig
	MysqlConfig            global.MysqlConfig
	TokenSecret            string          `json:",default=zeroimserver"`
	TokenRenewalDay        int64           `json:",default=30"` // 用户每次连接websocket 自动续签的的天数 默认30天
	MsgGatewayRpc          discov.EtcdConf `json:",optional"`
	MsgRpc                 zrpc.RpcClientConf
	MsgGatewayRpcK8sTarget string `json:",optional"`
}

type RedisConfig struct {
	redis.RedisConf
	DB int
}

type RateLimitConfig struct {
	Seconds int
	Quota   int
}
