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

type IfAInBBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewIfAInBBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfAInBBlacklistLogic {
	return &IfAInBBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

//  判断用户A是否在B黑名单中
func (l *IfAInBBlacklistLogic) IfAInBBlacklist(in *pb.IfAInBBlacklistReq) (*pb.IfAInBBlacklistResp, error) {
	blacklist := &model.Blacklist{}
	blacklist.SelfId = in.BUserID
	exist, err := l.rep.RelationCache.Exist(in.AUserID, blacklist, "user_id", map[string]interface{}{})
	if err != nil {
		if xormerr.TableNotFound(err) {
			_ = l.rep.Mysql.Table(blacklist.TableName()).AutoMigrate(blacklist)
		}
		return nil, err
	}
	return &pb.IfAInBBlacklistResp{
		CommonResp: &pb.CommonResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
		IsInBlacklist: exist,
	}, nil
}
