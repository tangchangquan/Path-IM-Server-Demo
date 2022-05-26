package rpclogic

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"
	"github.com/showurl/Path-IM-Server/app/msg-gateway/cmd/wsrpc/internal/wslogic"
	"github.com/showurl/Path-IM-Server/common/types"

	"github.com/showurl/Path-IM-Server/app/msg-gateway/cmd/wsrpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersOnlineStatusLogic struct {
	ctx    context.Context
	svcCtx *rpcsvc.ServiceContext
	logx.Logger
}

func NewGetUsersOnlineStatusLogic(ctx context.Context, svcCtx *rpcsvc.ServiceContext) *GetUsersOnlineStatusLogic {
	return &GetUsersOnlineStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersOnlineStatusLogic) GetUsersOnlineStatus(req *pb.GetUsersOnlineStatusReq) (*pb.GetUsersOnlineStatusResp, error) {
	var resp pb.GetUsersOnlineStatusResp
	logic := wslogic.NewMsggatewayLogic(nil, nil)
	for _, userID := range req.UserIDList {
		platformList := []string{
			types.IOSPlatformStr,
			types.AndroidPlatformStr,
			types.WindowsPlatformStr,
			types.OSXPlatformStr,
			types.WebPlatformStr,
			types.MiniWebPlatformStr,
			types.LinuxPlatformStr,
		}
		temp := new(pb.GetUsersOnlineStatusResp_SuccessResult)
		temp.UserID = userID
		for _, platform := range platformList {
			if conn := logic.GetUserConn(userID, platform); conn != nil {
				ps := new(pb.GetUsersOnlineStatusResp_SuccessDetail)
				ps.Platform = platform
				ps.Status = types.OnlineStatus
				temp.Status = types.OnlineStatus
				temp.DetailPlatformStatus = append(temp.DetailPlatformStatus, ps)

			}
		}
		if temp.Status == types.OnlineStatus {
			resp.SuccessResult = append(resp.SuccessResult, temp)
		}
	}
	return &resp, nil
}
