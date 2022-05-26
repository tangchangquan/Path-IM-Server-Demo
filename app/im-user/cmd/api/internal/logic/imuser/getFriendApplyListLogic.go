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

type GetFriendApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendApplyListLogic {
	return &GetFriendApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendApplyListLogic) GetFriendApplyList(req *types.GetFriendApplyListReq) (resp *types.GetFriendApplyListResp, err error) {
	rpcResp, err := l.svcCtx.RelationService().GetApplyFriendList(l.ctx, &pb.GetApplyFriendListReq{
		SelfId: ctxdata.GetUidFromCtx(l.ctx),
		PageReq: &pb.PageReq{
			Page:     int32(req.PageNo),
			PageSize: int32(req.PageSize),
		},
	})
	if err != nil {
		l.Errorf("GetApplyFriendList rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("GetApplyFriendList rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.GetFriendApplyListResp{}
	var userIds []string
	for _, black := range rpcResp.ApplyFriendList {
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
	for _, apply := range rpcResp.ApplyFriendList {
		user, ok := userMap[apply.UserId]
		if !ok {
			user = &pb.User{
				Id:       apply.UserId,
				Nickname: "用户已注销",
				Sign:     "",
				Avatar:   "",
			}
		}
		resp.ApplyList = append(resp.ApplyList, types.GetFriendApplyItem{
			Id:         apply.UserId,
			Nickname:   user.Nickname,
			Sign:       user.Sign,
			Avatar:     user.Avatar,
			CreateTime: apply.CreateTime,
			Status:     int(apply.Status),
			StatusStr:  apply.Status.String(),
		})
	}
	return
}
