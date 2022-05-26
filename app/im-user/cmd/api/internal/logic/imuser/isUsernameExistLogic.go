package imuser

import (
	"context"
	"fmt"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/showurl/Path-IM-Server/app/im-user/cmd/api/internal/svc"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUsernameExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsUsernameExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUsernameExistLogic {
	return &IsUsernameExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsUsernameExistLogic) IsUsernameExist(req *types.IsUsernameExistReq) (resp *types.IsUsernameExistResp, err error) {
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
	resp = &types.IsUsernameExistResp{
		Exist: rpcResp.User != nil && rpcResp.User.Id != "",
	}
	return
}
