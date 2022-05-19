// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userservice

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseResp             = pb.BaseResp
	GetUserByIdReq       = pb.GetUserByIdReq
	GetUserByUsernameReq = pb.GetUserByUsernameReq
	GetUserResp          = pb.GetUserResp
	InsertUserReq        = pb.InsertUserReq
	InsertUserResp       = pb.InsertUserResp
	LoginByIdReq         = pb.LoginByIdReq
	LoginByPasswordReq   = pb.LoginByPasswordReq
	LoginResp            = pb.LoginResp
	UpdateUserReq        = pb.UpdateUserReq
	UpdateUserResp       = pb.UpdateUserResp
	User                 = pb.User

	UserService interface {
		GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserResp, error)
		GetUserByUsername(ctx context.Context, in *GetUserByUsernameReq, opts ...grpc.CallOption) (*GetUserResp, error)
		InsertUser(ctx context.Context, in *InsertUserReq, opts ...grpc.CallOption) (*InsertUserResp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
		LoginByPassword(ctx context.Context, in *LoginByPasswordReq, opts ...grpc.CallOption) (*LoginResp, error)
		LoginById(ctx context.Context, in *LoginByIdReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultUserService) GetUserByUsername(ctx context.Context, in *GetUserByUsernameReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.GetUserByUsername(ctx, in, opts...)
}

func (m *defaultUserService) InsertUser(ctx context.Context, in *InsertUserReq, opts ...grpc.CallOption) (*InsertUserResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.InsertUser(ctx, in, opts...)
}

func (m *defaultUserService) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUserService) LoginByPassword(ctx context.Context, in *LoginByPasswordReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.LoginByPassword(ctx, in, opts...)
}

func (m *defaultUserService) LoginById(ctx context.Context, in *LoginByIdReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.LoginById(ctx, in, opts...)
}
