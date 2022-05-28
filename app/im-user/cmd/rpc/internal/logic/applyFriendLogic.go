package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/repository"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/types"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/model"
	chatpb "github.com/Path-IM/Path-IM-Server-Demo/app/msg/cmd/rpc/pb"
	timeUtils "github.com/Path-IM/Path-IM-Server-Demo/common/utils/time"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm/global"
	"gorm.io/gorm"

	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewApplyFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyFriendLogic {
	return &ApplyFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *ApplyFriendLogic) ApplyFriend(in *pb.ApplyFriendReq) (*pb.ApplyFriendResp, error) {
	m := &model.FriendApplyRecord{}
	m.SelfId = in.UserId
	count := int64(0)
	l.rep.Mysql.Model(m).Table(m.TableName()).
		Where("user_id = ?", in.ApplyUserId).
		Where("status = ?", 0).Count(&count)
	if count > 0 {
		return &pb.ApplyFriendResp{
			BaseResp: &pb.RelationBaseResp{},
		}, nil
	}
	record := &model.FriendApplyRecord{
		Id:        global.GetID(),
		UserId:    in.ApplyUserId,
		SelfId:    in.UserId,
		CreatedAt: timeUtils.Now().Unix(),
		Status:    0,
		Message:   in.Message,
	}
	err := xorm.Transaction(l.rep.Mysql,
		// 插入记录
		func(tx *gorm.DB) error {
			return xorm.Insert(tx, record)
		},
		// 发送通知
		func(_ *gorm.DB) error {
			_, err := l.svcCtx.MsgRpc().SendMsg(
				l.ctx,
				&chatpb.SendMsgReq{
					MsgData: types.NewSingleChatMsgNotification(
						in.ApplyUserId,
						in.UserId,
						types.NewApplyFriendContent(),
						types.ApplyFriendContentType,
						&chatpb.OfflinePushInfo{
							Title: "你有新的好友申请",
						},
					),
				},
			)
			if err != nil {
				return err
			}
			return nil
		},
	)
	return &pb.ApplyFriendResp{BaseResp: &pb.RelationBaseResp{
		ErrCode: 0,
		ErrMsg:  "",
	}}, err
}
