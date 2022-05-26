package main

import (
	"flag"
	"github.com/showurl/Path-IM-Server/app/msg-transfer/cmd/persistent/internal/config"
	"github.com/showurl/Path-IM-Server/app/msg-transfer/cmd/persistent/internal/server"
	"github.com/showurl/Path-IM-Server/app/msg-transfer/cmd/persistent/internal/svc"
	"github.com/showurl/Path-IM-Server/common/xconf"
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
