package repository

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/model"
	"gorm.io/gorm"
)

func (r *Rep) FuncDelFriendCache(ctx context.Context, friends ...*model.Friendlist) func(tx *gorm.DB) error {
	var keys []string
	for _, friend := range friends {
		k, _ := r.RelationCache.Key(friend, "user_id", map[string]interface{}{})
		keys = append(keys, k)
	}
	return func(_ *gorm.DB) error {
		return r.Cache.Del(ctx, keys...).Err()
	}
}

func (r *Rep) FuncDelBlackCache(ctx context.Context, blacklists ...*model.Blacklist) func(tx *gorm.DB) error {
	var keys []string
	for _, blacklist := range blacklists {
		k, _ := r.RelationCache.Key(blacklist, "user_id", map[string]interface{}{})
		keys = append(keys, k)
	}
	return func(_ *gorm.DB) error {
		return r.Cache.Del(ctx, keys...).Err()
	}
}
