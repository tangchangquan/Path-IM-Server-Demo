// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: chat.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	GetMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSeqResp, error)
	GetSuperGroupMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSuperGroupSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSuperGroupSeqResp, error)
	PullMessageBySeqList(ctx context.Context, in *WrapPullMessageBySeqListReq, opts ...grpc.CallOption) (*WrapPullMessageBySeqListResp, error)
	PullMessageBySuperGroupSeqList(ctx context.Context, in *PullMessageBySuperGroupSeqListReq, opts ...grpc.CallOption) (*WrapPullMessageBySeqListResp, error)
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) GetMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSeqResp, error) {
	out := new(GetMaxAndMinSeqResp)
	err := c.cc.Invoke(ctx, "/pbChat.Chat/GetMaxAndMinSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetSuperGroupMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSuperGroupSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSuperGroupSeqResp, error) {
	out := new(GetMaxAndMinSuperGroupSeqResp)
	err := c.cc.Invoke(ctx, "/pbChat.Chat/GetSuperGroupMaxAndMinSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) PullMessageBySeqList(ctx context.Context, in *WrapPullMessageBySeqListReq, opts ...grpc.CallOption) (*WrapPullMessageBySeqListResp, error) {
	out := new(WrapPullMessageBySeqListResp)
	err := c.cc.Invoke(ctx, "/pbChat.Chat/PullMessageBySeqList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) PullMessageBySuperGroupSeqList(ctx context.Context, in *PullMessageBySuperGroupSeqListReq, opts ...grpc.CallOption) (*WrapPullMessageBySeqListResp, error) {
	out := new(WrapPullMessageBySeqListResp)
	err := c.cc.Invoke(ctx, "/pbChat.Chat/PullMessageBySuperGroupSeqList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	out := new(SendMsgResp)
	err := c.cc.Invoke(ctx, "/pbChat.Chat/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	GetMaxAndMinSeq(context.Context, *GetMaxAndMinSeqReq) (*GetMaxAndMinSeqResp, error)
	GetSuperGroupMaxAndMinSeq(context.Context, *GetMaxAndMinSuperGroupSeqReq) (*GetMaxAndMinSuperGroupSeqResp, error)
	PullMessageBySeqList(context.Context, *WrapPullMessageBySeqListReq) (*WrapPullMessageBySeqListResp, error)
	PullMessageBySuperGroupSeqList(context.Context, *PullMessageBySuperGroupSeqListReq) (*WrapPullMessageBySeqListResp, error)
	SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) GetMaxAndMinSeq(context.Context, *GetMaxAndMinSeqReq) (*GetMaxAndMinSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaxAndMinSeq not implemented")
}
func (UnimplementedChatServer) GetSuperGroupMaxAndMinSeq(context.Context, *GetMaxAndMinSuperGroupSeqReq) (*GetMaxAndMinSuperGroupSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuperGroupMaxAndMinSeq not implemented")
}
func (UnimplementedChatServer) PullMessageBySeqList(context.Context, *WrapPullMessageBySeqListReq) (*WrapPullMessageBySeqListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullMessageBySeqList not implemented")
}
func (UnimplementedChatServer) PullMessageBySuperGroupSeqList(context.Context, *PullMessageBySuperGroupSeqListReq) (*WrapPullMessageBySeqListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullMessageBySuperGroupSeqList not implemented")
}
func (UnimplementedChatServer) SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_GetMaxAndMinSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMaxAndMinSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMaxAndMinSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/GetMaxAndMinSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMaxAndMinSeq(ctx, req.(*GetMaxAndMinSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetSuperGroupMaxAndMinSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMaxAndMinSuperGroupSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetSuperGroupMaxAndMinSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/GetSuperGroupMaxAndMinSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetSuperGroupMaxAndMinSeq(ctx, req.(*GetMaxAndMinSuperGroupSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_PullMessageBySeqList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrapPullMessageBySeqListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).PullMessageBySeqList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/PullMessageBySeqList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).PullMessageBySeqList(ctx, req.(*WrapPullMessageBySeqListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_PullMessageBySuperGroupSeqList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullMessageBySuperGroupSeqListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).PullMessageBySuperGroupSeqList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/PullMessageBySuperGroupSeqList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).PullMessageBySuperGroupSeqList(ctx, req.(*PullMessageBySuperGroupSeqListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbChat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMaxAndMinSeq",
			Handler:    _Chat_GetMaxAndMinSeq_Handler,
		},
		{
			MethodName: "GetSuperGroupMaxAndMinSeq",
			Handler:    _Chat_GetSuperGroupMaxAndMinSeq_Handler,
		},
		{
			MethodName: "PullMessageBySeqList",
			Handler:    _Chat_PullMessageBySeqList_Handler,
		},
		{
			MethodName: "PullMessageBySuperGroupSeqList",
			Handler:    _Chat_PullMessageBySuperGroupSeqList_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _Chat_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
