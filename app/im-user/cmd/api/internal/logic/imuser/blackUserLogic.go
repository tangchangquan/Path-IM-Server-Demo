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

type BlackUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlackUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlackUserLogic {
	return &BlackUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlackUserLogic) BlackUser(req *types.BlackUserReq) (resp *types.BlackUserResp, err error) {
	rpcResp, err := l.svcCtx.RelationService().BlackUser(l.ctx, &pb.BlackUserReq{
		SelfId: ctxdata.GetUidFromCtx(l.ctx),
		UserId: req.UserId,
	})
	if err != nil {
		l.Errorf("AgreeFriend rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("AgreeFriend rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.BlackUserResp{}
	return
}
