package repository

import (
	"context"
	"github.com/showurl/Zero-IM-Server/common/types"
	timeUtils "github.com/showurl/Zero-IM-Server/common/utils/time"
	"strconv"
	"time"
)

func (r *Rep) GetTokenMap(ctx context.Context, uid string, platform string) (map[string]int64, error) {
	key := types.RedisKeyToken + uid + ":" + platform
	result, err := r.Cache.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]int64{}, err
	}
	tokenMap := make(map[string]int64)
	for k, v := range result {
		tokenMap[k], _ = strconv.ParseInt(v, 10, 64)
	}
	return tokenMap, nil
}
func (r *Rep) SetTokenMap(ctx context.Context, uid string, platform string, token string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	expiredAt := timeUtils.Now().Add(time.Hour * 24 * time.Duration(r.svcCtx.Config.TokenRenewalDay))
	return r.Cache.HSet(ctx, key, token, expiredAt.UnixMilli()).Err()
}
func (r *Rep) DelTokenMap(ctx context.Context, uid string, platform string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	return r.Cache.Del(ctx, key).Err()
}

func (r *Rep) RenewalToken(ctx context.Context, uid string, platform string, token string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	expiredAt := timeUtils.Now().Add(time.Hour * 24 * time.Duration(r.svcCtx.Config.TokenRenewalDay))
	return r.Cache.HSet(ctx, key, token, expiredAt.UnixMilli()).Err()
}

func (r *Rep) DeleteToken(ctx context.Context, uid string, platform string, token string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	return r.Cache.HDel(ctx, key, token).Err()
}
