package config

import (
	"github.com/showurl/Path-IM-Server/common/xkafka"
	"github.com/showurl/Path-IM-Server/common/xorm/global"
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
