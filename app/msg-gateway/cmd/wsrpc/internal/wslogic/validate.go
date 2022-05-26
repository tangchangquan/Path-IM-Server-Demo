package wslogic

import (
	"github.com/golang/protobuf/proto"
	msggatewaypb "github.com/showurl/Path-IM-Server/app/msg-gateway/cmd/wsrpc/pb"
	chatpb "github.com/showurl/Path-IM-Server/app/msg/cmd/rpc/pb"
	"github.com/showurl/Path-IM-Server/common/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type Req struct {
	ReqIdentifier int32  `json:"reqIdentifier" validate:"required"`
	Token         string `json:"token" `
	SendID        string `json:"sendID" validate:"required"`
	MsgIncr       string `json:"msgIncr" validate:"required"`
	Data          []byte `json:"data"`
}
type Resp struct {
	ReqIdentifier int32  `json:"reqIdentifier"`
	MsgIncr       string `json:"msgIncr"`
	ErrCode       int32  `json:"errCode"`
	ErrMsg        string `json:"errMsg"`
	Data          []byte `json:"data"`
}

type SeqData struct {
	SeqBegin int64 `mapstructure:"seqBegin" validate:"required"`
	SeqEnd   int64 `mapstructure:"seqEnd" validate:"required"`
}
type MsgData struct {
	PlatformID  int32                  `mapstructure:"platformID" validate:"required"`
	SessionType int32                  `mapstructure:"sessionType" validate:"required"`
	MsgFrom     int32                  `mapstructure:"msgFrom" validate:"required"`
	ContentType int32                  `mapstructure:"contentType" validate:"required"`
	RecvID      string                 `mapstructure:"recvID" validate:"required"`
	ForceList   []string               `mapstructure:"forceList"`
	Content     string                 `mapstructure:"content" validate:"required"`
	Options     map[string]interface{} `mapstructure:"options" validate:"required"`
	ClientMsgID string                 `mapstructure:"clientMsgID" validate:"required"`
	OfflineInfo map[string]interface{} `mapstructure:"offlineInfo" validate:"required"`
	Ext         map[string]interface{} `mapstructure:"ext"`
}
type MaxSeqResp struct {
	MaxSeq int64 `json:"maxSeq"`
}
type PullMessageResp struct {
}
type SeqListData struct {
	SeqList []int64 `mapstructure:"seqList" validate:"required"`
}

func (l *MsggatewayLogic) argsValidate(m *msggatewaypb.Req, r int32) (isPass bool, errCode int32, errMsg string, returnData interface{}) {
	switch r {
	case types.WSSendMsg:
		data := chatpb.MsgData{}
		if err := proto.Unmarshal(m.Data, &data); err != nil {
			logx.WithContext(l.ctx).Error("Decode Data struct  err", err.Error(), r)
			return false, 203, err.Error(), nil
		}
		if err := validate.Struct(data); err != nil {
			logx.WithContext(l.ctx).Error("data args validate  err", err.Error(), r)
			return false, 204, err.Error(), nil

		}
		return true, 0, "", data
	case types.WSPullMsgBySeqList:
		data := chatpb.PullMessageBySeqListReq{}
		if err := proto.Unmarshal(m.Data, &data); err != nil {
			logx.WithContext(l.ctx).Error("Decode Data struct  err", err.Error(), r)
			return false, 203, err.Error(), nil
		}
		if err := validate.Struct(data); err != nil {
			logx.WithContext(l.ctx).Error("data args validate  err", err.Error(), r)
			return false, 204, err.Error(), nil
		}
		return true, 0, "", data
	case types.WSPullMsgByGroupSeqList:
		data := chatpb.PullMessageBySuperGroupSeqListReq{}
		if err := proto.Unmarshal(m.Data, &data); err != nil {
			logx.WithContext(l.ctx).Error("Decode Data struct  err", err.Error(), r)
			return false, 203, err.Error(), nil
		}
		if err := validate.Struct(data); err != nil {
			logx.WithContext(l.ctx).Error("data args validate  err", err.Error(), r)
			return false, 204, err.Error(), nil
		}
		return true, 0, "", data

	default:
	}
	return false, 204, "args err", nil
}
