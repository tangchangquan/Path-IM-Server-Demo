package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/repository"
	onlinemessagerelayservice "github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/onlineMessageRelayService"
	gatewaypb "github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/pb"
	jwtUtils "github.com/showurl/Zero-IM-Server/common/utils/jwt"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"

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
	for _, endpoint := range l.svcCtx.Config.MsgGatewayRpc.Endpoints {
		services = append(services, onlinemessagerelayservice.NewOnlineMessageRelayService(
			zrpc.MustNewClient(zrpc.RpcClientConf{
				Endpoints: []string{endpoint},
				Target:    l.svcCtx.Config.MsgGatewayRpc.Target,
				App:       l.svcCtx.Config.MsgGatewayRpc.App,
				Token:     l.svcCtx.Config.MsgGatewayRpc.Token,
				NonBlock:  true,
				Timeout:   0,
			}),
		))
	}
	return
}
