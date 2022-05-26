package xenv

import (
	"github.com/showurl/Path-IM-Server/common/xkafka"
	"github.com/showurl/Path-IM-Server/common/xmgo/global"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"os"
	"testing"
)

func TestSelfEnv(t *testing.T) {
	type RedisConfig struct {
		Conf redis.RedisConf
		DB   int
	}
	type MongoConfig struct {
		global.MongoConfig
		DBDatabase                      string
		DBTimeout                       int
		SingleChatMsgCollectionName     string
		SuperGroupChatMsgCollectionName string
	}
	type CallbackConfigItem struct {
		Enable          bool
		ContinueOnError bool
	}

	type CallbackConfig struct {
		CallbackWordFilter              CallbackConfigItem
		CallbackAtAllInSuperGroup       CallbackConfigItem
		CallbackBeforeSendGroupMsg      CallbackConfigItem
		CallbackAfterSendGroupMsg       CallbackConfigItem
		CallbackBeforeSendSuperGroupMsg CallbackConfigItem
		CallbackAfterSendSuperGroupMsg  CallbackConfigItem
		CallbackBeforeSendSingleMsg     CallbackConfigItem
		CallbackAfterSendSingleMsg      CallbackConfigItem
	}
	type SinglePushConsumerConfig struct {
		xkafka.ProducerConfig
		SinglePushGroupID string
	}
	type KafkaConfigOnline struct {
		xkafka.ProducerConfig
		MsgToMongoGroupID string
	}
	type KafkaConfig struct {
		Online         KafkaConfigOnline
		Offline        xkafka.ProducerConfig
		SinglePush     xkafka.ProducerConfig
		SuperGroupPush xkafka.ProducerConfig
	}
	os.Setenv("ZEROIM_IMUSERRPC_ETCD_ETCDCONF_HOSTS", "192.168.2.77:2379,192.168.2.77:2379")
	os.Setenv("ZEROIM_RESTCONF_SERVICECONF__HOST", "0.0.0.0")
	type conf struct {
		//rest.RestConf
		//zrpc.RpcServerConf
		//MsgGatewayRpc discov.EtcdConf
		//ImUserRpc zrpc.RpcClientConf
		RedisConfig RedisConfig
		//Mongo       MongoConfig
		//Callback CallbackConfig
		//SinglePushConsumer SinglePushConsumerConfig
		//Kafka KafkaConfig
		//UserRpc     zrpc.RpcClientConf
		//RelationRpc zrpc.RpcClientConf
	}
	c := &conf{}
	err := Parse(c)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(c.ImUserRpc.Etcd.Hosts)
	//t.Log(c.Name)
}
