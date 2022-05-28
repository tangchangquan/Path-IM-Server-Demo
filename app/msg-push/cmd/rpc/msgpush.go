package main

import (
	"flag"
	"fmt"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xconf"

	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-push/cmd/rpc/internal/config"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-push/cmd/rpc/internal/server"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-push/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-push/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/msgpush.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	xconf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewMsgPushServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMsgPushServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
