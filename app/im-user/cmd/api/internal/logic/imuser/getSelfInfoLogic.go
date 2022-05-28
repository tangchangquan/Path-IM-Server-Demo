package imuser

import (
	"context"
	"fmt"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"github.com/Path-IM/Path-IM-Server-Demo/common/ctxdata"

	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/api/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSelfInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSelfInfoLogic {
	return &GetSelfInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSelfInfoLogic) GetSelfInfo(req *types.GetSelfInfoReq) (resp *types.GetSelfInfoResp, err error) {
	rpcResp, err := l.svcCtx.UserService().GetUserById(l.ctx, &pb.GetUserByIdReq{
		UserId: ctxdata.GetUidFromCtx(l.ctx),
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
	resp = &types.GetSelfInfoResp{UserModel: types.UserModel{
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
