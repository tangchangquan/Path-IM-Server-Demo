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

type UpdateSelfInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSelfInfoLogic {
	return &UpdateSelfInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSelfInfoLogic) UpdateSelfInfo(req *types.ModifySelfInfoReq) (resp *types.ModifySelfInfoResp, err error) {
	rpcResp, err := l.svcCtx.UserService().UpdateUser(l.ctx, &pb.UpdateUserReq{
		User: &pb.User{
			Id:       ctxdata.GetUidFromCtx(l.ctx),
			Nickname: req.Nickname,
			Sign:     req.Sign,
			Avatar:   req.Avatar,
			Province: req.Province,
			City:     req.City,
			District: req.District,
			Birthday: req.Birthday,
			IsMale:   req.IsMale,
		},
		Fields: []string{
			"nickname",
			"sign",
			"avatar",
			"province",
			"city",
			"district",
			"birthday",
			"is_male",
		},
	})
	if err != nil {
		l.Errorf("UpdateUser rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("UpdateUser rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.ModifySelfInfoResp{}
	return
}
