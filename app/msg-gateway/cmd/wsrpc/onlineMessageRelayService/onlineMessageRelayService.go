// Code generated by goctl. DO NOT EDIT!
// Source: msg-gateway.proto

package onlinemessagerelayservice

import (
	"context"

	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-gateway/cmd/wsrpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUsersOnlineStatusReq                = pb.GetUsersOnlineStatusReq
	GetUsersOnlineStatusResp               = pb.GetUsersOnlineStatusResp
	GetUsersOnlineStatusResp_FailedDetail  = pb.GetUsersOnlineStatusResp_FailedDetail
	GetUsersOnlineStatusResp_SuccessDetail = pb.GetUsersOnlineStatusResp_SuccessDetail
	GetUsersOnlineStatusResp_SuccessResult = pb.GetUsersOnlineStatusResp_SuccessResult
	KickUserConnsReq                       = pb.KickUserConnsReq
	KickUserConnsResp                      = pb.KickUserConnsResp
	OnlinePushMsgReq                       = pb.OnlinePushMsgReq
	OnlinePushMsgResp                      = pb.OnlinePushMsgResp
	SingleMsgToUser                        = pb.SingleMsgToUser

	OnlineMessageRelayService interface {
		OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error)
		GetUsersOnlineStatus(ctx context.Context, in *GetUsersOnlineStatusReq, opts ...grpc.CallOption) (*GetUsersOnlineStatusResp, error)
		KickUserConns(ctx context.Context, in *KickUserConnsReq, opts ...grpc.CallOption) (*KickUserConnsResp, error)
	}

	defaultOnlineMessageRelayService struct {
		cli zrpc.Client
	}
)

func NewOnlineMessageRelayService(cli zrpc.Client) OnlineMessageRelayService {
	return &defaultOnlineMessageRelayService{
		cli: cli,
	}
}

func (m *defaultOnlineMessageRelayService) OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error) {
	client := pb.NewOnlineMessageRelayServiceClient(m.cli.Conn())
	return client.OnlinePushMsg(ctx, in, opts...)
}

func (m *defaultOnlineMessageRelayService) GetUsersOnlineStatus(ctx context.Context, in *GetUsersOnlineStatusReq, opts ...grpc.CallOption) (*GetUsersOnlineStatusResp, error) {
	client := pb.NewOnlineMessageRelayServiceClient(m.cli.Conn())
	return client.GetUsersOnlineStatus(ctx, in, opts...)
}

func (m *defaultOnlineMessageRelayService) KickUserConns(ctx context.Context, in *KickUserConnsReq, opts ...grpc.CallOption) (*KickUserConnsResp, error) {
	client := pb.NewOnlineMessageRelayServiceClient(m.cli.Conn())
	return client.KickUserConns(ctx, in, opts...)
}
