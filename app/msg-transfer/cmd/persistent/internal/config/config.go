package config

import (
	"github.com/Path-IM/Path-IM-Server-Demo/common/xkafka"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm/global"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf
	Kafka       KafkaConfig
	MysqlConfig global.MysqlConfig
}
type KafkaConfigOnline struct {
	xkafka.ProducerConfig
	MsgToMysqlGroupID string
}
type KafkaConfig struct {
	Online KafkaConfigOnline
}
