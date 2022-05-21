# 写http api
> logic 基本就是调用rpc 没什么逻辑 所以放上一个示例

## 鉴权中间件
```go
package middleware

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/imuserservice"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"github.com/showurl/Zero-IM-Server/common/ctxdata"
	"github.com/showurl/Zero-IM-Server/common/xerr"
	"github.com/showurl/Zero-IM-Server/common/xhttp"
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

```

## 密码登录
```go
package logic

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 获取用户信息
	rpcResp, err := l.svcCtx.UserService().GetUserByUsername(l.ctx, &pb.GetUserByUsernameReq{
		Username: req.Username,
	})
	if err != nil {
		l.Errorf("GetSelfInfo rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("GetSelfInfo rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	// 判断密码是否一致
	l.Infof("rpcResp user: %+v", rpcResp.User.String())
	if encrypt.Md5(req.Password) != rpcResp.User.Password {
		// 密码错误
		err = fmt.Errorf("密码错误")
		return
	}
	loginResp, err := l.svcCtx.UserService().LoginById(l.ctx, &pb.LoginByIdReq{
		UserId:   rpcResp.User.Id,
		Platform: ctxdata.GetPlatformFromCtx(l.ctx),
	})
	if err != nil {
		l.Errorf("LoginById rpc error: %v", err)
		return
	}
	if rpcResp.BaseResp.ErrCode != 0 {
		l.Errorf("LoginById rpc error: %v", rpcResp.BaseResp.ErrMsg)
		err = fmt.Errorf("%v", rpcResp.BaseResp.ErrMsg)
		return
	}
	resp = &types.LoginResp{
		UserModel: types.UserModel{
			ID:           rpcResp.User.Id,
			Username:     rpcResp.User.Username,
			Nickname:     rpcResp.User.Nickname,
			Sign:         rpcResp.User.Sign,
			Avatar:       rpcResp.User.Avatar,
			Province:     rpcResp.User.Province,
			City:         rpcResp.User.City,
			District:     rpcResp.User.District,
			Birthday:     rpcResp.User.Birthday,
			RegisterTime: rpcResp.User.RegisterTime,
			IsMale:       rpcResp.User.IsMale,
		},
		Token: loginResp.Token,
	}
	return
}

```

## 用户注册 jaeger
![jaeger.png](https://raw.githubusercontent.com/showurl/Zero-IM-Docs/main/images/202205)