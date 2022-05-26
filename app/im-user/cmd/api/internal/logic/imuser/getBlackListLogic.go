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

type GetBlackListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBlackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlackListLogic {
	return &GetBlackListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBlackListLogic) GetBlackList(req *types.GetBlackListReq) (resp *types.GetBlackListResp, err error) {
	rpcResp, err := l.svcCtx.RelationService().GetBlackList(l.ctx, &pb.GetBlackListReq{
		SelfId: ctxdata.GetUidFromCtx(l.ctx),
		PageReq: &pb.PageReq{
			Page:     int32(req.PageNo),
			PageSize: int32(req.PageSize),
		},
	})
	if err != nil {
		l.Errorf("GetBlackList rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("GetBlackList rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.GetBlackListResp{}
	var userIds []string
	for _, black := range rpcResp.BlackList {
		userIds = append(userIds, black.UserId)
	}
	// 获取用户信息
	userListResp, err := l.svcCtx.UserService().GetUserByIds(l.ctx, &pb.GetUserByIdsReq{
		UserIds: userIds,
	})
	if err != nil {
		l.Errorf("GetUserByIds rpc error: %v", err)
		return
	}
	if userListResp.BaseResp.ErrCode != 0 {
		l.Errorf("GetUserByIds rpc error: %v", userListResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", userListResp.BaseResp.ErrMsg)
		return
	}
	userMap := make(map[string]*pb.User)
	for _, user := range userListResp.Users {
		userMap[user.Id] = user
	}
	for _, black := range rpcResp.BlackList {
		user, ok := userMap[black.UserId]
		if !ok {
			user = &pb.User{
				Id:       black.UserId,
				Nickname: "用户已注销",
				Sign:     "",
				Avatar:   "",
			}
		}
		resp.BlackList = append(resp.BlackList, types.RealtionUser{
			Id:         black.UserId,
			Nickname:   user.Nickname,
			Sign:       user.Sign,
			Avatar:     user.Avatar,
			CreateTime: black.CreateTime,
		})
	}
	return
}
