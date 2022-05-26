package logic

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/types"
	"github.com/showurl/Path-IM-Server/app/im-user/model"
	chatpb "github.com/showurl/Path-IM-Server/app/msg/cmd/rpc/pb"
	"github.com/showurl/Path-IM-Server/common/xorm"
	"gorm.io/gorm"

	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *DeleteFriendLogic) DeleteFriend(in *pb.DeleteFriendReq) (*pb.DeleteFriendResp, error) {
	friend1 := &model.Friendlist{
		UserId: in.UserId,
		SelfId: in.SelfId,
	}
	friend2 := &model.Friendlist{
		UserId: in.SelfId,
		SelfId: in.UserId,
	}
	err := xorm.Transaction(l.rep.Mysql,
		func(tx *gorm.DB) error {
			err := tx.Model(friend1).Table(friend1.TableName()).Where("user_id = ?", friend1.UserId).
				Unscoped().Delete(&model.Friendlist{}).Error
			if err != nil {
				return err
			}
			err = tx.Model(friend2).Table(friend2.TableName()).Where("user_id = ?", friend2.UserId).
				Unscoped().Delete(&model.Friendlist{}).Error
			if err != nil {
				return err
			}
			return nil
		},
		l.rep.FuncDelFriendCache(l.ctx, friend1, friend2),
		// 发通知
		func(_ *gorm.DB) error {
			_, err := l.svcCtx.MsgRpc().SendMsg(
				l.ctx,
				&chatpb.SendMsgReq{
					MsgData: types.NewSingleChatMsgNotification(
						in.SelfId,
						in.UserId,
						types.NewDeleteFriendContent(),
						types.DeleteFriendContentType,
						nil,
					),
				},
			)
			if err != nil {
				return err
			}
			return nil
		},
	)
	return &pb.DeleteFriendResp{BaseResp: &pb.RelationBaseResp{}}, err
}
