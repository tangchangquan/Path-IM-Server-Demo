package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	WhiteRateLimit RateLimitConfig
	Redis          redis.RedisConf
	rest.RestConf
	ImUserRpc   zrpc.RpcClientConf
	UserRpc     zrpc.RpcClientConf
	RelationRpc zrpc.RpcClientConf
}

type RateLimitConfig struct {
	Seconds int
	Quota   int
}
