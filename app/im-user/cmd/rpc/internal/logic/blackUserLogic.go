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

type BlackUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewBlackUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlackUserLogic {
	return &BlackUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *BlackUserLogic) BlackUser(in *pb.BlackUserReq) (*pb.BlackUserResp, error) {
	blacklist := &model.Blacklist{
		UserId: in.UserId,
		SelfId: in.SelfId,
	}
	err := xorm.Transaction(l.rep.Mysql,
		func(tx *gorm.DB) error {
			return xorm.Insert(tx, blacklist)
		},
		l.rep.FuncDelBlackCache(l.ctx, blacklist),
	)
	return &pb.BlackUserResp{BaseResp: &pb.RelationBaseResp{}}, err
}
