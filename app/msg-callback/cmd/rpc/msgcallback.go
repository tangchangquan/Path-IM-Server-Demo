package main

import (
	"flag"
	"fmt"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xconf"

	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-callback/cmd/rpc/internal/config"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-callback/cmd/rpc/internal/server"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-callback/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-callback/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/msgcallback.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	xconf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewMsgcallbackServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMsgcallbackServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
