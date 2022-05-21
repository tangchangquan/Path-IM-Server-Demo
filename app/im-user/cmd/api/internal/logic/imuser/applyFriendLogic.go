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

type ApplyFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyFriendLogic {
	return &ApplyFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyFriendLogic) ApplyFriend(req *types.ApplyFriendReq) (resp *types.ApplyFriendResp, err error) {
	rpcResp, err := l.svcCtx.RelationService().ApplyFriend(l.ctx, &pb.ApplyFriendReq{
		ApplyUserId: ctxdata.GetUidFromCtx(l.ctx),
		UserId:      req.UserId,
		Message:     req.Message,
	})
	if err != nil {
		l.Errorf("ApplyFriend rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("ApplyFriend rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.ApplyFriendResp{}
	return
}
