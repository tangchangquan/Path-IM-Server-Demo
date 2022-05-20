package repository

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/model"
	"github.com/showurl/Zero-IM-Server/common/xcache/rc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

func (r *Rep) WarmUpUser(ctx context.Context, user *model.User) {
	var err error
	err = r.DetailCache.FirstById(user, user.Id)
	if err != nil {
		logx.WithContext(ctx).Errorf("WarmUpUser: %s", err)
	}
	err = r.DetailCache.FirstByWhere(user, map[string][]interface{}{
		"username = ?": {user.Username},
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("WarmUpUser: %s", err)
	}
}
func (r *Rep) FuncDelUserCache(ctx context.Context, users ...*model.User) func(tx *gorm.DB) error {
	if len(users) == 0 {
		return func(tx *gorm.DB) error {
			return nil
		}
	}
	return func(_ *gorm.DB) error {
		var keys []string
		for _, user := range users {
			idK := r.DetailCache.KeyById(user, user.Id)
			unameK := r.DetailCache.Key(user,
				map[string][]interface{}{
					"username=?": {user.Username},
				})
			logx.WithContext(ctx).Infof("del user cache keys: %+v", []string{idK, unameK})
			keys = append(keys, idK, unameK)
		}
		err := r.Cache.Del(ctx, keys...).Err()
		if err != nil {
			return err
		}
		return nil
	}
}
func (r *Rep) FuncDelInsertUserCache(ctx context.Context, user *model.User, group *model.Group) func(tx *gorm.DB) error {
	var keys []string
	idK := r.DetailCache.KeyById(user, user.Id)
	unameK := r.DetailCache.Key(user,
		map[string][]interface{}{
			"username=?": {user.Username},
		})
	keys = append(keys, idK, unameK)
	record := &model.SuperGroupConversationRecord{
		UserId:     "",
		GroupId:    group.Id,
		RecvMsgOpt: 0,
		Remark:     "",
	}
	groupK, _ := r.RelationCache.Key(
		record,
		"user_id",
		map[string]interface{}{
			"recv_msg_opt": 0,
		},
		rc.Order("user_id"),
	)
	keys = append(keys, groupK)
	logx.WithContext(ctx).Infof("del user cache keys: %+v", keys)
	return func(tx *gorm.DB) error {
		err := r.Cache.Del(ctx, keys...).Err()
		if err != nil {
			return err
		}
		return nil
	}
}
