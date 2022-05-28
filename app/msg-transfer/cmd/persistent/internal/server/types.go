package server

import (
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-transfer/cmd/persistent/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xkafka"
	"sync"
)

const OnlineTopicBusy = 1
const OnlineTopicVacancy = 0

type fcb func(msg []byte, msgKey string) error
type Cmd2Value struct {
	Cmd   int
	Value interface{}
}

type MsgTransferPersistentServer struct {
	svcCtx                  *svc.ServiceContext
	msgHandle               map[string]fcb
	persistentConsumerGroup *xkafka.MConsumerGroup
	cmdCh                   chan Cmd2Value
	w                       *sync.Mutex
	OnlineTopicStatus       int
}
