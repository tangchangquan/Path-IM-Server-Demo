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

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	oldUser, err := NewGetUserByIdLogic(l.ctx, l.svcCtx).GetUserById(&pb.GetUserByIdReq{UserId: in.User.Id})
	if err != nil {
		return nil, err
	}
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
	updateMap := make(map[string]interface{})
	for _, field := range in.Fields {
		updateMap[field] = user.Value(field)
	}
	err = xorm.Transaction(l.rep.Mysql,
		// 插入数据
		func(tx *gorm.DB) error {
			return tx.Model(&model.User{}).Table(user.TableName()).
				Where("id = ?", user.Id).
				Updates(updateMap).Error
		},
		// 清除缓存
		l.rep.FuncDelUserCache(l.ctx, &model.User{
			Id:       oldUser.User.Id,
			Username: oldUser.User.Username,
		}, user),
	)
	if err != nil {
		return &pb.UpdateUserResp{
			BaseResp: &pb.BaseResp{
				ErrCode: -1,
				ErrMsg:  err.Error(),
			},
		}, err
	}
	// 预热数据
	go l.rep.WarmUpUser(l.ctx, user)
	return &pb.UpdateUserResp{
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}
