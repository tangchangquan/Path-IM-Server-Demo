package ctxdata

import (
	"context"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) string {
	var uid string
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(string); ok {
		uid = jsonUid
	}
	return uid
}

// CtxKeyPlatform get uid from ctx
var CtxKeyPlatform = "platform"

func GetPlatformFromCtx(ctx context.Context) string {
	var platform string
	if jsonPlatform, ok := ctx.Value(CtxKeyPlatform).(string); ok {
		platform = jsonPlatform
	}
	return platform
}
