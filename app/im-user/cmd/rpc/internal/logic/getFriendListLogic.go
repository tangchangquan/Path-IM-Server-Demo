package logic

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
	"github.com/showurl/Path-IM-Server/app/im-user/model"
	timeUtils "github.com/showurl/Path-IM-Server/common/utils/time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *pb.GetFriendListReq) (*pb.GetFriendListResp, error) {
	var friendlists []*model.Friendlist
	var resplist []*pb.GetFriendListResp_Friend
	friendlist := &model.Friendlist{}
	friendlist.SelfId = in.SelfId
	err := l.rep.Mysql.Model(friendlist).Table(friendlist.TableName()).Find(&friendlists).Error
	if err != nil {
		l.Errorf("get friend list error: %s", err.Error())
		return nil, err
	}
	for _, f := range friendlists {
		resplist = append(resplist, &pb.GetFriendListResp_Friend{
			UserId:     f.UserId,
			SelfId:     f.SelfId,
			Remark:     f.Remark,
			CreateTime: timeUtils.Unix2Time(f.CreatedAt).Format("2006-01-02 15:04:05"),
		})
	}
	return &pb.GetFriendListResp{
		BaseResp:   &pb.RelationBaseResp{},
		FriendList: resplist,
	}, err
}
