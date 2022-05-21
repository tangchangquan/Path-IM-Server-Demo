package logic

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"github.com/showurl/Zero-IM-Server/app/im-user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetUserByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdsLogic {
	return &GetUserByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GetUserByIdsLogic) GetUserByIds(in *pb.GetUserByIdsReq) (*pb.GetUserListResp, error) {
	users := make([]*model.User, 0)
	err := l.rep.DetailCache.ListByIds(&model.User{}, &users, in.UserIds)
	if err != nil {
		l.Errorf("get user by ids error: %s", err.Error())
		return &pb.GetUserListResp{
			Users: nil,
			BaseResp: &pb.BaseResp{
				ErrCode: -1,
				ErrMsg:  err.Error(),
			},
		}, err
	}
	var userInfos []*pb.User
	for _, user := range users {
		if user != nil {
			userInfos = append(userInfos, &pb.User{
				Id:           user.Id,
				Username:     user.Username,
				Password:     user.Password,
				Nickname:     user.Nickname,
				Sign:         user.Sign,
				Avatar:       user.Avatar,
				Province:     user.Province,
				City:         user.City,
				District:     user.District,
				Birthday:     user.Birthday,
				RegisterTime: user.RegisterTime,
				IsMale:       user.IsMale,
			})
		}
	}

	return &pb.GetUserListResp{
		Users: userInfos,
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}
