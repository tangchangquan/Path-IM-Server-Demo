package middleware

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/imuserservice"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"github.com/Path-IM/Path-IM-Server-Demo/common/ctxdata"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xerr"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xhttp"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
)

type JwtAuthMiddleware struct {
	imuserrpc func() imuserservice.ImUserService
	getLimit  func() *limit.PeriodLimit
}

func NewJwtAuthMiddleware(
	getImUserRpc func() imuserservice.ImUserService,
	getLimit func() *limit.PeriodLimit,
) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{getImUserRpc, getLimit}
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
				// 根据ip限流
				if !m.RateLimit(ctx, ip) {
					logx.WithContext(req.Context()).Error("[业务告警]", "用户频繁请求白名单接口：", "ip：", ip, "ua：", ua, "av：", av, "uri：", req.RequestURI)
					xhttp.HttpResult(req, resp, nil, xerr.New(xerr.IP_RATE_LIMIT, "您操作太过频繁"))
					return
				}
				next(resp, req.WithContext(ctx))
				return
			}
			xhttp.HttpResult(req, resp, nil, xerr.New(xerr.AUTH_ERROR, "登录已过期"))
			logx.WithContext(req.Context()).Error("[业务告警]", "用户破解接口访问失败：", "token：", token, "ip：", ip, "ua：", ua, "av：", av)
			return
		}
		response, err := m.imuserrpc().VerifyToken(req.Context(), &pb.VerifyTokenReq{
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

func (m *JwtAuthMiddleware) RateLimit(ctx context.Context, ip string) bool {
	takeRes, err := m.getLimit().Take(ip)
	if err != nil {
		logx.WithContext(ctx).Errorf("ip:%s, rate limit err:", err)
		return false
	}
	switch takeRes {
	case limit.OverQuota:
		logx.WithContext(ctx).Errorf("ip:%s, rate limit OverQuota", ip)
		return false
	case limit.Allowed:
		return true
	case limit.HitQuota:
		logx.WithContext(ctx).Errorf("ip:%s, rate limit HitQuota", ip)
		return false
	default:
		return true
	}
}
