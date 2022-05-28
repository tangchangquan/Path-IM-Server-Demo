package main

import (
	"flag"
	"fmt"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xconf"

	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/api/internal/config"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/api/internal/handler"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/imuser.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	xconf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
