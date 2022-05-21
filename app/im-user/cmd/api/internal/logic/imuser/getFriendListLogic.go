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

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList(req *types.GetFriendListReq) (resp *types.GetFriendListResp, err error) {
	rpcResp, err := l.svcCtx.RelationService().GetFriendList(l.ctx, &pb.GetFriendListReq{
		SelfId: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		l.Errorf("GetFriendList rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("GetFriendList rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.GetFriendListResp{}
	var userIds []string
	for _, black := range rpcResp.FriendList {
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
	for _, friend := range rpcResp.FriendList {
		user, ok := userMap[friend.UserId]
		if !ok {
			user = &pb.User{
				Id:       friend.UserId,
				Nickname: "用户已注销",
				Sign:     "",
				Avatar:   "",
			}
		}
		resp.FriendList = append(resp.FriendList, types.RealtionUser{
			Id:         friend.UserId,
			Nickname:   user.Nickname,
			Sign:       user.Sign,
			Avatar:     user.Avatar,
			CreateTime: friend.CreateTime,
		})
	}
	return
}
