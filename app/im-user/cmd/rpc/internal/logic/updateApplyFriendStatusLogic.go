package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/repository"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/types"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/model"
	chatpb "github.com/Path-IM/Path-IM-Server-Demo/app/msg/cmd/rpc/pb"
	timeUtils "github.com/Path-IM/Path-IM-Server-Demo/common/utils/time"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm"
	xormerr "github.com/Path-IM/Path-IM-Server-Demo/common/xorm/err"
	"gorm.io/gorm"

	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApplyFriendStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdateApplyFriendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApplyFriendStatusLogic {
	return &UpdateApplyFriendStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdateApplyFriendStatusLogic) UpdateApplyFriendStatus(in *pb.UpdateApplyFriendStatusReq) (*pb.UpdateApplyFriendStatusResp, error) {
	record := &model.FriendApplyRecord{}
	record.Id = in.Id
	record.SelfId = in.SelfId
	err := xorm.DetailByWhere(l.rep.Mysql, record, xorm.Where("id = ?", in.Id))
	if err != nil {
		l.Errorf("update apply friend status error: %s", err.Error())
		return nil, err
	}
	if record.Status != 0 {
		return &pb.UpdateApplyFriendStatusResp{BaseResp: &pb.RelationBaseResp{}}, nil
	}
	err = xorm.Transaction(l.rep.Mysql,
		func(tx *gorm.DB) error {
			if in.Status == pb.UpdateApplyFriendStatusReq_REFUSE {
				err := tx.Model(&model.FriendApplyRecord{}).
					Table(record.TableName()).Where("user_id = ?", record.UserId).
					Where("status = ?", 0).
					Updates(map[string]interface{}{"status": 2}).Error
				if err != nil {
					l.Errorf("update apply friend status error: %s", err.Error())
					return err
				}
				return nil
			} else {
				err := tx.Model(&model.FriendApplyRecord{}).
					Table(record.TableName()).Where("user_id = ?", record.UserId).
					Where("status = ?", 0).
					Updates(map[string]interface{}{"status": 1}).Error
				if err != nil {
					l.Errorf("update apply friend status error: %s", err.Error())
					return err
				}
				// 添加好友关系
				unix := timeUtils.Now().Unix()
				friend1 := &model.Friendlist{
					UserId:    record.UserId,
					SelfId:    record.SelfId,
					CreatedAt: unix,
					Remark:    "",
				}
				friend2 := &model.Friendlist{
					UserId:    record.SelfId,
					SelfId:    record.UserId,
					CreatedAt: unix,
					Remark:    "",
				}
				err = xorm.Insert(tx, friend1)
				if err != nil {
					if xormerr.DuplicateError(err) {
						err = nil
					} else {
						l.Errorf("insert friend error: %s", err.Error())
						return err
					}
				}
				err = xorm.Insert(tx, friend2)
				if err != nil {
					if xormerr.DuplicateError(err) {
						err = nil
					} else {
						l.Errorf("insert friend error: %s", err.Error())
						return err
					}
				}
				// 清除缓存
				err = l.rep.FuncDelFriendCache(l.ctx, friend1, friend2)(tx)
				if err != nil {
					return err
				}
				// 系统单聊消息
				_, err = l.svcCtx.MsgRpc().SendMsg(
					l.ctx,
					&chatpb.SendMsgReq{
						MsgData: types.NewSingleChatMsgNotification(
							record.SelfId,
							record.UserId,
							types.NewApplyFriendSuccessContent(),
							types.ApplyFriendSuccessContentType,
							nil,
						),
					},
				)
				if err != nil {
					return err
				}
				return nil
			}
		},
	)
	return &pb.UpdateApplyFriendStatusResp{BaseResp: &pb.RelationBaseResp{
		ErrCode: 0,
		ErrMsg:  "",
	}}, err
}
