package config

import (
	"github.com/showurl/Zero-IM-Server/common/xorm/global"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConfig RedisConfig
	MysqlConfig global.MysqlConfig
}

type RedisConfig struct {
	redis.RedisConf
	DB int
}
