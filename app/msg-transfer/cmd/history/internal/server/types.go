package server

import (
	"github.com/showurl/Path-IM-Server/app/msg-transfer/cmd/history/internal/svc"
	"github.com/showurl/Path-IM-Server/common/xkafka"
	"sync"
)

const OnlineTopicBusy = 1
const OnlineTopicVacancy = 0

type fcb func(msg []byte, msgKey string) error
type Cmd2Value struct {
	Cmd   int
	Value interface{}
}

type MsgTransferHistoryServer struct {
	svcCtx               *svc.ServiceContext
	msgHandle            map[string]fcb
	historyConsumerGroup *xkafka.MConsumerGroup
	cmdCh                chan Cmd2Value
	w                    *sync.Mutex
	OnlineTopicStatus    int
}
