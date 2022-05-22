package onlinemessagerelayservice

import (
	"context"
	"github.com/showurl/Zero-IM-Server/common/xetcd"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

// GetAll 获取所有 service
// @param msggatewayRpcEtcdKey msggateway-rpc 的 etcd key
func GetAll(
	ctx context.Context,
	etcdConf discov.EtcdConf,
	msggatewayRpcEtcdKey string,
) (services []OnlineMessageRelayService, err error) {
	var zrpcConns []zrpc.Client
	zrpcConns, err = xetcd.GetGoZeroZrpcConns(ctx, xetcd.NewClient(etcdConf), msggatewayRpcEtcdKey)
	if err != nil {
		return
	}
	for _, conn := range zrpcConns {
		service := NewOnlineMessageRelayService(conn)
		services = append(services, service)
	}
	return
}
