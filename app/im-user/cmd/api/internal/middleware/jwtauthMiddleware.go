package middleware

import (
	"context"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/imuserservice"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/rpc/pb"
	"github.com/showurl/Path-IM-Server/common/ctxdata"
	"github.com/showurl/Path-IM-Server/common/xerr"
	"github.com/showurl/Path-IM-Server/common/xhttp"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
)

type JwtAuthMiddleware struct {
	f func() imuserservice.ImUserService
}

func NewJwtAuthMiddleware(f func() imuserservice.ImUserService) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{f}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("token")
		ip := req.Header.Get("x-real-ip")
		ua := req.Header.Get("user-agent")
		av := req.Header.Get("app-version")
		platform := req.Header.Get("platform")
		userId := req.Header.Get("User_id")
		if len(strings.Split(token, ".")) != 3 {
			// 判断是否是白名单接口
			if isWhite(req) {
				ctx := context.WithValue(req.Context(), ctxdata.CtxKeyJwtUserId, userId)
				ctx = context.WithValue(ctx, ctxdata.CtxKeyPlatform, platform)
				next(resp, req.WithContext(ctx))
				return
			}
			xhttp.HttpResult(req, resp, nil, xerr.New(xerr.AUTH_ERROR, "登录已过期"))
			logx.WithContext(req.Context()).Error("[业务告警]", "用户破解接口访问失败：", "token：", token, "ip：", ip, "ua：", ua, "av：", av)
			return
		}
		response, err := m.f().VerifyToken(req.Context(), &pb.VerifyTokenReq{
			Token:    token,
			Platform: platform,
			SendID:   userId,
		})
		if err != nil {
			xhttp.ParamErrorResult(req, resp, err)
			return
		}
		{
			req.Header.Set("user-id", response.Uid)
		}
		if !response.Success {
			// 判断是否是白名单接口
			if isWhite(req) {
				ctx := context.WithValue(req.Context(), ctxdata.CtxKeyJwtUserId, userId)
				ctx = context.WithValue(ctx, ctxdata.CtxKeyPlatform, platform)
				next(resp, req)
				return
			}
			xhttp.HttpResult(req, resp, nil, xerr.New(xerr.AUTH_ERROR, "登录已过期"))
			return
		} else {
			ctx := context.WithValue(req.Context(), ctxdata.CtxKeyJwtUserId, response.Uid)
			ctx = context.WithValue(ctx, ctxdata.CtxKeyPlatform, platform)
			next(resp, req.WithContext(ctx))
			return
		}
	}
}

func isWhite(req *http.Request) bool {
	return strings.Contains(req.URL.Path, "/white/")
}
