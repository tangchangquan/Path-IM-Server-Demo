package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/repository"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/model"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xcache/global"
	xormerr "github.com/Path-IM/Path-IM-Server-Demo/common/xorm/err"

	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSingleConversationRecvMsgOptsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetSingleConversationRecvMsgOptsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSingleConversationRecvMsgOptsLogic {
	return &GetSingleConversationRecvMsgOptsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

//  获取单聊会话的消息接收选项
func (l *GetSingleConversationRecvMsgOptsLogic) GetSingleConversationRecvMsgOpts(in *pb.GetSingleConversationRecvMsgOptsReq) (*pb.GetSingleConversationRecvMsgOptsResp, error) {
	record := &model.SingleConversationRecord{}
	record.UserId = in.UserID
	record.ConversationId = in.ConversationID
	err := l.rep.DetailCache.FirstByWhere(
		record,
		map[string][]interface{}{
			"conversation_id": {in.ConversationID},
		},
	)
	if err != nil {
		if xormerr.RecordNotFound(err) || err == global.RedisErrorNotExists {
			err = nil
		} else if xormerr.TableNotFound(err) {
			l.rep.Mysql.Table(record.TableName()).AutoMigrate(record)
			return nil, err
		} else {
			return nil, err
		}
	}
	return &pb.GetSingleConversationRecvMsgOptsResp{
		CommonResp: &pb.CommonResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
		Opts: pb.RecvMsgOpt(record.RecvMsgOpt),
	}, nil
}
