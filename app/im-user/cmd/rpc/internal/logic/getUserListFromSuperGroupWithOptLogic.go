package logic

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Zero-IM-Server/app/im-user/model"
	"github.com/showurl/Zero-IM-Server/common/xcache/global"
	"github.com/showurl/Zero-IM-Server/common/xcache/rc"
	xormerr "github.com/showurl/Zero-IM-Server/common/xorm/err"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListFromSuperGroupWithOptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetUserListFromSuperGroupWithOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListFromSuperGroupWithOptLogic {
	return &GetUserListFromSuperGroupWithOptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

//  获取超级群成员列表 通过消息接收选项
func (l *GetUserListFromSuperGroupWithOptLogic) GetUserListFromSuperGroupWithOpt(in *pb.GetUserListFromSuperGroupWithOptReq) (*pb.GetUserListFromSuperGroupWithOptResp, error) {
	resp := &pb.GetUserListFromSuperGroupWithOptResp{
		CommonResp:    &pb.CommonResp{},
		UserIDOptList: nil,
	}
	record := &model.SuperGroupConversationRecord{}
	record.GroupId = in.SuperGroupID
	for _, opt := range in.Opts {
		var userIds []string
		err := l.rep.RelationCache.List(
			&userIds,
			0,
			-1,
			record,
			"user_id",
			map[string]interface{}{
				"recv_msg_opt": opt,
			},
			rc.Order("user_id"),
		)
		if err != nil {
			if xormerr.TableNotFound(err) {
				l.rep.Mysql.Table(record.TableName()).AutoMigrate(record)
			}
			if global.RedisErrorNotExists == err {
				continue
			}
			return nil, err
		}
		for _, id := range userIds {
			resp.UserIDOptList = append(resp.UserIDOptList, &pb.UserIDOpt{
				UserID: id,
				Opts:   opt,
			})
		}
	}
	return resp, nil
}
