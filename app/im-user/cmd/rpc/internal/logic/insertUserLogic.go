package logic

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Zero-IM-Server/app/im-user/model"
	"github.com/showurl/Zero-IM-Server/common/xorm"
	"gorm.io/gorm"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewInsertUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUserLogic {
	return &InsertUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *InsertUserLogic) InsertUser(in *pb.InsertUserReq) (*pb.InsertUserResp, error) {
	user := &model.User{
		Id:           in.User.Id,
		Username:     in.User.Username,
		Password:     in.User.Password,
		Nickname:     in.User.Nickname,
		Sign:         in.User.Sign,
		Avatar:       in.User.Avatar,
		Province:     in.User.Province,
		City:         in.User.City,
		District:     in.User.District,
		Birthday:     in.User.Birthday,
		RegisterTime: in.User.RegisterTime,
		IsMale:       in.User.IsMale,
	}
	err := xorm.Transaction(l.rep.Mysql,
		// 插入数据
		func(tx *gorm.DB) error {
			return xorm.Insert(tx, user)
		},
		// 清除缓存
		l.rep.FuncDelUserCache(l.ctx, user),
	)
	if err != nil {
		return &pb.InsertUserResp{
			BaseResp: &pb.BaseResp{
				ErrCode: -1,
				ErrMsg:  err.Error(),
			},
		}, err
	}
	// 预热数据
	go l.rep.WarmUpUser(l.ctx, user)
	return &pb.InsertUserResp{
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}
