package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/repository"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/model"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm"
	"gorm.io/gorm"

	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBlackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewDeleteBlackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBlackLogic {
	return &DeleteBlackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *DeleteBlackLogic) DeleteBlack(in *pb.DeleteBlackReq) (*pb.DeleteBlackResp, error) {
	blacklist := &model.Blacklist{
		UserId: in.UserId,
		SelfId: in.SelfId,
	}
	err := xorm.Transaction(l.rep.Mysql,
		func(tx *gorm.DB) error {
			err := tx.Model(blacklist).Table(blacklist.TableName()).Where("user_id = ?", blacklist.UserId).
				Unscoped().Delete(&model.Blacklist{}).Error
			if err != nil {
				return err
			}
			return nil
		},
		l.rep.FuncDelBlackCache(l.ctx, blacklist),
	)
	return &pb.DeleteBlackResp{BaseResp: &pb.RelationBaseResp{}}, err
}
