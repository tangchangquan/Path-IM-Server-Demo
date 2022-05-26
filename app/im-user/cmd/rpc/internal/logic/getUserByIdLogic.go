package logic

import (
	"context"
	"fmt"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Path-IM-Server/app/im-user/model"
	"github.com/showurl/Path-IM-Server/common/xcache/global"
	xormerr "github.com/showurl/Path-IM-Server/common/xorm/err"

	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserResp, error) {
	if in.UserId == "" {
		return nil, fmt.Errorf("用户id不能为空")
	}
	user := &model.User{}
	err := l.rep.DetailCache.FirstById(user, in.UserId)
	if err != nil {
		if xormerr.RecordNotFound(err) {
			err = nil
		}
		if err == global.RedisErrorNotExists {
			err = nil
		}
		if err != nil {
			l.Errorf("get user by id error: %s", err.Error())
			return &pb.GetUserResp{
				User: nil,
				BaseResp: &pb.BaseResp{
					ErrCode: -1,
					ErrMsg:  err.Error(),
				},
			}, err
		}
	}
	var userInfo *pb.User
	if user != nil {
		userInfo = &pb.User{
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
		}
	}
	return &pb.GetUserResp{
		User: userInfo,
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}
