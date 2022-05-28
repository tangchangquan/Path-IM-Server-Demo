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

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(in *pb.GetUserByUsernameReq) (*pb.GetUserResp, error) {
	user := &model.User{}
	user.Username = in.Username
	err := l.rep.DetailCache.FirstByWhere(
		user,
		map[string][]interface{}{
			"username=?": {in.Username},
		},
	)
	if err != nil {
		if xormerr.RecordNotFound(err) {
			err = nil
		}
		if err == global.RedisErrorNotExists {
			err = nil
		}
		if err != nil {
			l.Errorf("get user by username error: %s", err.Error())
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
