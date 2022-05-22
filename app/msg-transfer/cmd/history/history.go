package main

import (
	"flag"
	"github.com/showurl/Zero-IM-Server/app/msg-transfer/cmd/history/internal/config"
	"github.com/showurl/Zero-IM-Server/app/msg-transfer/cmd/history/internal/server"
	"github.com/showurl/Zero-IM-Server/app/msg-transfer/cmd/history/internal/svc"
	"github.com/showurl/Zero-IM-Server/common/xconf"
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
