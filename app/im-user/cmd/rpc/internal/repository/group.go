package repository

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/types"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/model"
	chatpb "github.com/Path-IM/Path-IM-Server-Demo/app/msg/cmd/rpc/pb"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm"
	xormerr "github.com/Path-IM/Path-IM-Server-Demo/common/xorm/err"
	"gorm.io/gorm"
)

func (r *Rep) FuncJoinGroup(user *model.User, group *model.Group) func(tx *gorm.DB) error {
	record := &model.SuperGroupConversationRecord{
		UserId:     user.Id,
		GroupId:    group.Id,
		RecvMsgOpt: 0,
		Remark:     "",
	}
	return func(tx *gorm.DB) error {
		err := xorm.Insert(tx, record)
		if err != nil {
			if xormerr.DuplicateError(err) {
				err = nil
			} else {
				return err
			}
		}
		return nil
	}
}

func (r *Rep) SendGroupTextMsg(
	ctx context.Context, user *model.User, group *model.Group, text string) error {
	_, err := r.svcCtx.MsgRpc().SendMsg(
		ctx,
		&chatpb.SendMsgReq{
			MsgData: types.NewGroupTextMsg(user.Id, group.Id, text),
		},
	)
	return err
}
