package repository

import (
	"context"
	"github.com/showurl/Path-IM-Server/common/types"
)

func (r *Rep) GetUserMaxSeq(uid string) (uint64, error) {
	key := types.RedisKeyUserIncrSeq + uid
	count, err := r.Cache.Get(context.Background(), key).Int64()
	return uint64(count), err
}
func (r *Rep) GetUserMinSeq(uid string) (uint64, error) {
	key := types.RedisKeyUserMinSeq + uid
	count, err := r.Cache.Get(context.Background(), key).Int64()
	return uint64(count), err
}

func (r *Rep) GetSuperGroupMaxSeq(groupId string) (uint64, error) {
	key := types.RedisKeyGroupIncrSeq + groupId
	count, err := r.Cache.Get(context.Background(), key).Int64()
	return uint64(count), err
}

func (r *Rep) GetSuperGroupMinSeq(groupId string) (uint64, error) {
	key := types.RedisKeyGroupMinSeq + groupId
	count, err := r.Cache.Get(context.Background(), key).Int64()
	return uint64(count), err
}
