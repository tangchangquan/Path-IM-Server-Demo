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

type GetBlackListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetBlackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlackListLogic {
	return &GetBlackListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GetBlackListLogic) GetBlackList(in *pb.GetBlackListReq) (*pb.GetBlackListResp, error) {
	var blacklists []*model.Blacklist
	var resplist []*pb.GetBlackListResp_Black
	var count int64
	blacklist := &model.Blacklist{}
	blacklist.SelfId = in.SelfId
	tx := l.rep.Mysql.Model(blacklist).Table(blacklist.TableName())
	tx.Count(&count)
	err := xorm.Paging(tx, int(in.PageReq.Page), int(in.PageReq.PageSize)).Find(&blacklists).Error
	if err != nil {
		l.Errorf("get black list error: %s", err.Error())
		return nil, err
	}
	for _, f := range blacklists {
		resplist = append(resplist, &pb.GetBlackListResp_Black{
			UserId:     f.UserId,
			SelfId:     f.SelfId,
			CreateTime: timeUtils.Unix2Time(f.CreatedAt).Format("2006-01-02 15:04:05"),
		})
	}
	return &pb.GetBlackListResp{BlackList: resplist, BaseResp: &pb.RelationBaseResp{}, PageResp: &pb.PageResp{
		Page:     in.PageReq.Page,
		PageSize: in.PageReq.PageSize,
		Total:    count,
	}}, nil
}
