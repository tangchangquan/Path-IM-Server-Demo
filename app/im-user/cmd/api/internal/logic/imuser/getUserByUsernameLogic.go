package imuser

import (
	"context"
	"fmt"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/types"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(req *types.GetUserByUsernameReq) (resp *types.GetUserByUsernameResp, err error) {
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
	resp = &types.GetUserByUsernameResp{UserModel: types.UserModel{
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
	}}
	return
}
