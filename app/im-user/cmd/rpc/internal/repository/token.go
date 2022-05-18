package repository

import (
	"context"
	"github.com/showurl/Zero-IM-Server/common/types"
	timeUtils "github.com/showurl/Zero-IM-Server/common/utils/time"
	"strconv"
	"time"
)

func (r *Rep) GetTokenMap(ctx context.Context, uid string, platform string) (map[string]int, error) {
	key := types.RedisKeyToken + uid + ":" + platform
	result, err := r.Cache.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]int{}, err
	}
	tokenMap := make(map[string]int)
	for k, v := range result {
		tokenMap[k], _ = strconv.Atoi(v)
	}
	return tokenMap, nil
}

func (r *Rep) RenewalToken(ctx context.Context, uid string, platform string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	return r.Cache.ExpireAt(ctx, key, timeUtils.Now().Add(time.Hour*24*time.Duration(r.svcCtx.Config.TokenRenewalDay))).Err()
}
