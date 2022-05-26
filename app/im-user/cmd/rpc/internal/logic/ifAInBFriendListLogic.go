package logic

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Path-IM-Server/app/im-user/model"
	xormerr "github.com/showurl/Path-IM-Server/common/xorm/err"

	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfAInBFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewIfAInBFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfAInBFriendListLogic {
	return &IfAInBFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

//  判断用户A是否在B好友列表中
func (l *IfAInBFriendListLogic) IfAInBFriendList(in *pb.IfAInBFriendListReq) (*pb.IfAInBFriendListResp, error) {
	friendlist := &model.Friendlist{}
	friendlist.SelfId = in.BUserID
	exist, err := l.rep.RelationCache.Exist(
		in.AUserID,
		friendlist,
		"user_id",
		map[string]interface{}{})
	if err != nil {
		if xormerr.TableNotFound(err) {
			_ = l.rep.Mysql.Table(friendlist.TableName()).AutoMigrate(friendlist)
		}
		return nil, err
	}
	return &pb.IfAInBFriendListResp{
		CommonResp: &pb.CommonResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
		IsInFriendList: exist,
	}, nil
}
