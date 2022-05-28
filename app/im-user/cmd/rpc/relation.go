package main

import (
	"flag"
	"fmt"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xconf"

	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/config"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/server"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var RelationConfigFile = flag.String("f", "etc/relation.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	xconf.MustLoad(*RelationConfigFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewRelationServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterRelationServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
