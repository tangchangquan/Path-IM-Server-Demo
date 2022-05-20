// Code generated by goctl. DO NOT EDIT!
// Source: relation.proto

package relationservice

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApplyFriendReq                     = pb.ApplyFriendReq
	ApplyFriendResp                    = pb.ApplyFriendResp
	BlackUserReq                       = pb.BlackUserReq
	BlackUserResp                      = pb.BlackUserResp
	DeleteBlackReq                     = pb.DeleteBlackReq
	DeleteBlackResp                    = pb.DeleteBlackResp
	DeleteFriendReq                    = pb.DeleteFriendReq
	DeleteFriendResp                   = pb.DeleteFriendResp
	GetApplyFriendListReq              = pb.GetApplyFriendListReq
	GetApplyFriendListResp             = pb.GetApplyFriendListResp
	GetApplyFriendListResp_ApplyFriend = pb.GetApplyFriendListResp_ApplyFriend
	GetBlackListReq                    = pb.GetBlackListReq
	GetBlackListResp                   = pb.GetBlackListResp
	GetBlackListResp_Black             = pb.GetBlackListResp_Black
	GetFriendListReq                   = pb.GetFriendListReq
	GetFriendListResp                  = pb.GetFriendListResp
	GetFriendListResp_Friend           = pb.GetFriendListResp_Friend
	PageReq                            = pb.PageReq
	PageResp                           = pb.PageResp
	RelationBaseResp                   = pb.RelationBaseResp
	UpdateApplyFriendStatusReq         = pb.UpdateApplyFriendStatusReq
	UpdateApplyFriendStatusResp        = pb.UpdateApplyFriendStatusResp

	RelationService interface {
		ApplyFriend(ctx context.Context, in *ApplyFriendReq, opts ...grpc.CallOption) (*ApplyFriendResp, error)
		UpdateApplyFriendStatus(ctx context.Context, in *UpdateApplyFriendStatusReq, opts ...grpc.CallOption) (*UpdateApplyFriendStatusResp, error)
		GetApplyFriendList(ctx context.Context, in *GetApplyFriendListReq, opts ...grpc.CallOption) (*GetApplyFriendListResp, error)
		GetFriendList(ctx context.Context, in *GetFriendListReq, opts ...grpc.CallOption) (*GetFriendListResp, error)
		DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...grpc.CallOption) (*DeleteFriendResp, error)
		GetBlackList(ctx context.Context, in *GetBlackListReq, opts ...grpc.CallOption) (*GetBlackListResp, error)
		DeleteBlack(ctx context.Context, in *DeleteBlackReq, opts ...grpc.CallOption) (*DeleteBlackResp, error)
		BlackUser(ctx context.Context, in *BlackUserReq, opts ...grpc.CallOption) (*BlackUserResp, error)
	}

	defaultRelationService struct {
		cli zrpc.Client
	}
)

func NewRelationService(cli zrpc.Client) RelationService {
	return &defaultRelationService{
		cli: cli,
	}
}

func (m *defaultRelationService) ApplyFriend(ctx context.Context, in *ApplyFriendReq, opts ...grpc.CallOption) (*ApplyFriendResp, error) {
	client := pb.NewRelationServiceClient(m.cli.Conn())
	return client.ApplyFriend(ctx, in, opts...)
}

func (m *defaultRelationService) UpdateApplyFriendStatus(ctx context.Context, in *UpdateApplyFriendStatusReq, opts ...grpc.CallOption) (*UpdateApplyFriendStatusResp, error) {
	client := pb.NewRelationServiceClient(m.cli.Conn())
	return client.UpdateApplyFriendStatus(ctx, in, opts...)
}

func (m *defaultRelationService) GetApplyFriendList(ctx context.Context, in *GetApplyFriendListReq, opts ...grpc.CallOption) (*GetApplyFriendListResp, error) {
	client := pb.NewRelationServiceClient(m.cli.Conn())
	return client.GetApplyFriendList(ctx, in, opts...)
}

func (m *defaultRelationService) GetFriendList(ctx context.Context, in *GetFriendListReq, opts ...grpc.CallOption) (*GetFriendListResp, error) {
	client := pb.NewRelationServiceClient(m.cli.Conn())
	return client.GetFriendList(ctx, in, opts...)
}

func (m *defaultRelationService) DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...grpc.CallOption) (*DeleteFriendResp, error) {
	client := pb.NewRelationServiceClient(m.cli.Conn())
	return client.DeleteFriend(ctx, in, opts...)
}

func (m *defaultRelationService) GetBlackList(ctx context.Context, in *GetBlackListReq, opts ...grpc.CallOption) (*GetBlackListResp, error) {
	client := pb.NewRelationServiceClient(m.cli.Conn())
	return client.GetBlackList(ctx, in, opts...)
}

func (m *defaultRelationService) DeleteBlack(ctx context.Context, in *DeleteBlackReq, opts ...grpc.CallOption) (*DeleteBlackResp, error) {
	client := pb.NewRelationServiceClient(m.cli.Conn())
	return client.DeleteBlack(ctx, in, opts...)
}

func (m *defaultRelationService) BlackUser(ctx context.Context, in *BlackUserReq, opts ...grpc.CallOption) (*BlackUserResp, error) {
	client := pb.NewRelationServiceClient(m.cli.Conn())
	return client.BlackUser(ctx, in, opts...)
}