package imuser

import (
	"context"
	"fmt"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
	"github.com/showurl/Path-IM-Server/common/ctxdata"
	"github.com/showurl/Path-IM-Server/common/xorm/global"

	"github.com/showurl/Path-IM-Server/app/im-user/cmd/api/internal/svc"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	rpcResp, err := l.svcCtx.UserService().GetUserByUsername(l.ctx, &pb.GetUserByUsernameReq{
		Username: req.Username,
	})
	if err != nil {
		l.Errorf("GetUserByUsername rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("GetUserByUsername rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	if rpcResp.User != nil && rpcResp.User.Id != "" {
		l.Errorf("GetUserByUsername rpc error: %v", "用户名已存在")
		err = fmt.Errorf("%v", "用户名已存在")
		return
	}
	uid := global.GetID()
	insertUserResp, err := l.svcCtx.UserService().InsertUser(l.ctx, &pb.InsertUserReq{
		User: &pb.User{
			Id:           uid,
			Username:     req.Username,
			Password:     req.Password,
			Nickname:     "Zero用户",
			Avatar:       "",
			Sign:         "",
			Province:     "",
			City:         "",
			District:     "",
			Birthday:     "",
			RegisterTime: "",
			IsMale:       true,
		},
	})
	if err != nil {
		l.Errorf("InsertUser rpc error: %v", err)
		return
	}
	if insertUserResp.BaseResp.ErrCode != 0 {
		l.Errorf("InsertUser rpc error: %v", insertUserResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", insertUserResp.BaseResp.ErrMsg)
		return
	}
	// loginbyid
	loginResp, err := l.svcCtx.UserService().LoginById(l.ctx, &pb.LoginByIdReq{
		UserId:   uid,
		Platform: ctxdata.GetPlatformFromCtx(l.ctx),
	})
	if err != nil {
		l.Errorf("LoginById rpc error: %v", err)
		return
	}
	if loginResp.BaseResp.ErrCode != 0 {
		l.Errorf("LoginById rpc error: %v", loginResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", "注册成功 请登录")
		return
	}
	// userinfo
	userInfoResp, err := l.svcCtx.UserService().GetUserById(l.ctx, &pb.GetUserByIdReq{
		UserId: uid,
	})
	if err != nil {
		l.Errorf("GetUserById rpc error: %v", err)
		return
	}
	if userInfoResp.BaseResp.ErrCode != 0 {
		l.Errorf("GetUserById rpc error: %v", userInfoResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", "注册成功 请登录")
		return
	}
	resp = &types.RegisterResp{
		UserModel: types.UserModel{
			ID:           userInfoResp.User.Id,
			Username:     userInfoResp.User.Username,
			Nickname:     userInfoResp.User.Nickname,
			Sign:         userInfoResp.User.Sign,
			Avatar:       userInfoResp.User.Avatar,
			Province:     userInfoResp.User.Province,
			City:         userInfoResp.User.City,
			District:     userInfoResp.User.District,
			Birthday:     userInfoResp.User.Birthday,
			RegisterTime: userInfoResp.User.RegisterTime,
			IsMale:       userInfoResp.User.IsMale,
		},
		Token: loginResp.Token,
	}
	return
}
