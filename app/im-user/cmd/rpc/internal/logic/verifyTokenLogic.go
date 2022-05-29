package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/repository"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	jwtUtils "github.com/Path-IM/Path-IM-Server-Demo/common/utils/jwt"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xerr"
	"github.com/zeromicro/go-zero/core/limit"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

// 检查token
func (l *VerifyTokenLogic) VerifyToken(in *pb.VerifyTokenReq) (*pb.VerifyTokenResp, error) {
	// 检查限流
	{
		takeRes, err := l.rep.RateLimiter.Take(in.Token)
		if err != nil {
			l.Errorf("check rate limiter error: %v", err)
		} else {
			switch takeRes {
			case limit.OverQuota:
				l.Errorf("token:%s id:%s, rate limit OverQuota", in.Token, in.SendID)
				return &pb.VerifyTokenResp{
					ErrMsg: "rate limit OverQuota",
				}, xerr.New(xerr.TOKEN_RATE_LIMIT, "您的操作过于频繁，请稍后再试")
			case limit.HitQuota:
				l.Errorf("token:%s id:%s, rate limit HitQuota", in.Token, in.SendID)
				return &pb.VerifyTokenResp{
					ErrMsg: "rate limit HitQuota",
				}, xerr.New(xerr.TOKEN_RATE_LIMIT, "您的操作过于频繁，请稍后再试")
			}
		}
	}
	return &pb.VerifyTokenResp{
		Uid:     in.SendID,
		Success: true,
		ErrMsg:  "",
	}, nil
	claim, err := jwtUtils.GetClaimFromToken(in.Token, l.svcCtx.Config.TokenSecret)
	if err != nil {
		return nil, err
	}
	tokenMap, err := l.rep.GetTokenMap(l.ctx, claim.UID, claim.Platform)
	if err != nil {
		return nil, err
	}
	if ex, ok := tokenMap[in.Token]; ok {
		if time.Now().UnixMilli() > ex {
			e := l.rep.DeleteToken(l.ctx, claim.UID, claim.Platform, in.Token)
			if e != nil {
				l.Errorf("delete token error: %s", e.Error())
			}
			return &pb.VerifyTokenResp{
				Uid:     "",
				Success: false,
				ErrMsg:  "token expired",
			}, nil
		}
		// 有的话就更新过期时间
		go func() {
			_ = l.rep.RenewalToken(l.ctx, claim.UID, claim.Platform, in.Token)
		}()
		return &pb.VerifyTokenResp{
			Uid:     claim.UID,
			Success: true,
			ErrMsg:  "",
		}, nil
	} else {
		// 没有 token 提示
		return &pb.VerifyTokenResp{
			Uid:     "",
			Success: false,
			ErrMsg:  "token is not exist",
		}, nil
	}
}
