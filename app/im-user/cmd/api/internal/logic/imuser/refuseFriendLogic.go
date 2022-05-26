package imuser

import (
	"context"
	"fmt"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
	"github.com/showurl/Path-IM-Server/common/ctxdata"

	"github.com/showurl/Path-IM-Server/app/im-user/cmd/api/internal/svc"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefuseFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefuseFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefuseFriendLogic {
	return &RefuseFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefuseFriendLogic) RefuseFriend(req *types.RefuseFriendReq) (resp *types.RefuseFriendResp, err error) {
	rpcResp, err := l.svcCtx.RelationService().UpdateApplyFriendStatus(l.ctx, &pb.UpdateApplyFriendStatusReq{
		Status: pb.UpdateApplyFriendStatusReq_REFUSE,
		Id:     req.ApplyId,
		SelfId: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		l.Errorf("RefuseFriend rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("RefuseFriend rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.RefuseFriendResp{}
	return
}
