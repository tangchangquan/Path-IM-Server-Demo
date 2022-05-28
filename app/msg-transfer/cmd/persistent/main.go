package main

import (
	"flag"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-transfer/cmd/persistent/internal/config"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-transfer/cmd/persistent/internal/server"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-transfer/cmd/persistent/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xconf"
)

var configFile = flag.String("f", "etc/persistent.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	xconf.MustLoad(*configFile, &c)
	err := c.ServiceConf.SetUp()
	if err != nil {
		panic(err)
	}
	s := server.NewMsgTransferPersistentServer(svc.NewServiceContext(c))
	s.Start()
}
