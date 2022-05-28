package main

import (
	"flag"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-transfer/cmd/history-cassandra/internal/config"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-transfer/cmd/history-cassandra/internal/server"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-transfer/cmd/history-cassandra/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xconf"
)

var configFile = flag.String("f", "etc/history.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	xconf.MustLoad(*configFile, &c)
	err := c.ServiceConf.SetUp()
	if err != nil {
		panic(err)
	}
	s := server.NewMsgTransferHistoryServer(svc.NewServiceContext(c))
	s.Start()
}
