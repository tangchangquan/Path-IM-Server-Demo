package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/repository"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	onlinemessagerelayservice "github.com/Path-IM/Path-IM-Server-Demo/app/msg-gateway/cmd/wsrpc/onlineMessageRelayService"
	gatewaypb "github.com/Path-IM/Path-IM-Server-Demo/app/msg-gateway/cmd/wsrpc/pb"
	jwtUtils "github.com/Path-IM/Path-IM-Server-Demo/common/utils/jwt"
	"github.com/golang-jwt/jwt/v4"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewLoginByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByIdLogic {
	return &LoginByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *LoginByIdLogic) LoginById(in *pb.LoginByIdReq) (*pb.LoginResp, error) {
	claim := jwtUtils.BuildClaims(in.UserId, in.Platform)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(l.svcCtx.Config.TokenSecret))
	if err != nil {
		l.Errorf("jwt signed string error: %v", err)
		return nil, err
	}
	// 获取tokenmap
	tokenMap, err := l.rep.GetTokenMap(l.ctx, claim.UID, claim.Platform)
	if err != nil {
		return nil, err
	}
	if len(tokenMap) > 0 {
		// 删除tokenmap
		err = l.rep.DelTokenMap(l.ctx, claim.UID, claim.Platform)
		if err != nil {
			l.Errorf("del token map error: %v", err)
			return nil, err
		}
		// 断开用户之前的连接
		services, err := l.getAllMsgGatewayService()
		if err != nil {
			return nil, err
		}
		for _, service := range services {
			_, err := service.KickUserConns(l.ctx, &gatewaypb.KickUserConnsReq{
				UserID:      claim.UID,
				PlatformIDs: []string{claim.Platform},
			})
			if err != nil {
				l.Errorf("kick user conns error: %v", err)
				return nil, err
			}
		}
	}
	// 写入tokenmap
	err = l.rep.SetTokenMap(l.ctx, claim.UID, claim.Platform, signedString)
	if err != nil {
		l.Errorf("set token map error: %v", err)
		return nil, err
	}
	return &pb.LoginResp{
		User:  nil,
		Token: signedString,
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}

func (l *LoginByIdLogic) getAllMsgGatewayService() (services []onlinemessagerelayservice.OnlineMessageRelayService, err error) {
	if l.svcCtx.Config.MsgGatewayRpcK8sTarget == "" {
		return onlinemessagerelayservice.GetAllByEtcd(l.ctx, l.svcCtx.Config.MsgGatewayRpc, l.svcCtx.Config.MsgGatewayRpc.Key)
	} else {
		return onlinemessagerelayservice.GetAllByK8s(l.svcCtx.Config.MsgGatewayRpcK8sTarget)
	}
}
