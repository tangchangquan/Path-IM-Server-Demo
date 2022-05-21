package imuser

import (
	"context"
	"fmt"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"github.com/showurl/Zero-IM-Server/common/ctxdata"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnBlackUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnBlackUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnBlackUserLogic {
	return &UnBlackUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnBlackUserLogic) UnBlackUser(req *types.UnBlackUserReq) (resp *types.UnBlackUserResp, err error) {
	rpcResp, err := l.svcCtx.RelationService().DeleteBlack(l.ctx, &pb.DeleteBlackReq{
		SelfId: ctxdata.GetUidFromCtx(l.ctx),
		UserId: req.UserId,
	})
	if err != nil {
		l.Errorf("DeleteBlack rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("DeleteBlack rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.UnBlackUserResp{}
	return
}
