package logic

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	jwtUtils "github.com/showurl/Zero-IM-Server/common/utils/jwt"

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
	claim, err := jwtUtils.GetClaimFromToken(in.Token, l.svcCtx.Config.TokenSecret)
	if err != nil {
		return nil, err
	}
	tokenMap, err := l.rep.GetTokenMap(l.ctx, claim.UID, claim.Platform)
	if err != nil {
		return nil, err
	}
	if _, ok := tokenMap[in.Token]; ok {
		// 有的话就更新过期时间
		go func() {
			_ = l.rep.RenewalToken(l.ctx, claim.UID, claim.Platform)
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
