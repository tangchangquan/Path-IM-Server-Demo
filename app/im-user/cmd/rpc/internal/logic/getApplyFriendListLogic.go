package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/repository"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/model"
	timeUtils "github.com/Path-IM/Path-IM-Server-Demo/common/utils/time"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm"

	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplyFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetApplyFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplyFriendListLogic {
	return &GetApplyFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GetApplyFriendListLogic) GetApplyFriendList(in *pb.GetApplyFriendListReq) (*pb.GetApplyFriendListResp, error) {
	var (
		count   int64
		records []*model.FriendApplyRecord
		record  = &model.FriendApplyRecord{}
		list    []*pb.GetApplyFriendListResp_ApplyFriend
		err     error
	)
	record.SelfId = in.SelfId
	tx := l.rep.Mysql.Model(record).
		Table(record.TableName())
	tx.Count(&count)
	tx = xorm.Paging(tx, int(in.PageReq.Page), int(in.PageReq.PageSize))
	err = tx.Order("created_at DESC").Find(&records).Error
	for _, applyRecord := range records {
		list = append(list, &pb.GetApplyFriendListResp_ApplyFriend{
			Id:          applyRecord.Id,
			ApplyUserId: applyRecord.UserId,
			UserId:      applyRecord.SelfId,
			Message:     applyRecord.Message,
			CreateTime:  timeUtils.Unix2Time(applyRecord.CreatedAt).Format("2006-01-02 15:04:05"),
			Status:      pb.GetApplyFriendListResp_ApplyFriend_Status(applyRecord.Status),
		})
	}
	return &pb.GetApplyFriendListResp{
		BaseResp: &pb.RelationBaseResp{},
		PageResp: &pb.PageResp{
			Page:     in.PageReq.Page,
			PageSize: in.PageReq.PageSize,
			Total:    count,
		},
		ApplyFriendList: list,
	}, err
}
