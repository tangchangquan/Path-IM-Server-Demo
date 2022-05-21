package imuser

import (
	"context"
	"fmt"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"github.com/showurl/Zero-IM-Server/common/ctxdata"
	"github.com/showurl/Zero-IM-Server/common/utils/encrypt"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 获取用户信息
	rpcResp, err := l.svcCtx.UserService().GetUserByUsername(l.ctx, &pb.GetUserByUsernameReq{
		Username: req.Username,
	})
	if err != nil {
		l.Errorf("GetSelfInfo rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("GetSelfInfo rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	// 判断密码是否一致
	l.Infof("rpcResp user: %+v", rpcResp.User.String())
	if encrypt.Md5(req.Password) != rpcResp.User.Password {
		// 密码错误
		err = fmt.Errorf("密码错误")
		return
	}
	loginResp, err := l.svcCtx.UserService().LoginById(l.ctx, &pb.LoginByIdReq{
		UserId:   rpcResp.User.Id,
		Platform: ctxdata.GetPlatformFromCtx(l.ctx),
	})
	if err != nil {
		l.Errorf("LoginById rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("LoginById rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.LoginResp{
		UserModel: types.UserModel{
			ID:           rpcResp.User.Id,
			Username:     rpcResp.User.Username,
			Nickname:     rpcResp.User.Nickname,
			Sign:         rpcResp.User.Sign,
			Avatar:       rpcResp.User.Avatar,
			Province:     rpcResp.User.Province,
			City:         rpcResp.User.City,
			District:     rpcResp.User.District,
			Birthday:     rpcResp.User.Birthday,
			RegisterTime: rpcResp.User.RegisterTime,
			IsMale:       rpcResp.User.IsMale,
		},
		Token: loginResp.Token,
	}
	return
}
