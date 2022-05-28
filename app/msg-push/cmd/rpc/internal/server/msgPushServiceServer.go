// Code generated by goctl. DO NOT EDIT!
// Source: msgpush.proto

package server

import (
	"context"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-push/cmd/rpc/internal/logic"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-push/cmd/rpc/internal/svc"
	pushpb "github.com/Path-IM/Path-IM-Server-Demo/app/msg-push/cmd/rpc/pb"
	chatpb "github.com/Path-IM/Path-IM-Server-Demo/app/msg/cmd/rpc/pb"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xkafka"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xtrace"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/attribute"
)

type fcb func(msg []byte, msgKey string)
type Cmd2Value struct {
	Cmd   int
	Value interface{}
}

type MsgPushServiceServer struct {
	svcCtx *svc.ServiceContext
	pushpb.UnimplementedMsgPushServiceServer
	SingleConsumerGroup     *xkafka.MConsumerGroup
	SuperGroupConsumerGroup *xkafka.MConsumerGroup
}

func (s *MsgPushServiceServer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (s *MsgPushServiceServer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (s *MsgPushServiceServer) ConsumeSingle(sess sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage) {
	msgFromMQ := chatpb.PushMsgDataToMQ{}
	if err := proto.Unmarshal(msg.Value, &msgFromMQ); err != nil {
		logx.Errorf("unmarshal msg error: %v", err)
		return
	}
	xtrace.RunWithTrace(msgFromMQ.OperationID, func(ctx context.Context) {
		xtrace.StartFuncSpan(ctx, "MsgPushServiceServer.ConsumeSingle.PushMsg", func(ctx context.Context) {
			_, _ = s.PushMsg(ctx, &pushpb.PushMsgReq{
				MsgData:      msgFromMQ.MsgData,
				PushToUserID: msgFromMQ.PushToUserID,
			})
			sess.MarkMessage(msg, "")
		})

	}, attribute.String("msg.key", string(msg.Key)))
}

func (s *MsgPushServiceServer) ConsumeSuperGroup(sess sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage) {
	msgFromMQ := &chatpb.PushMsgToSuperGroupDataToMQ{}
	if err := proto.Unmarshal(msg.Value, msgFromMQ); err != nil {
		logx.Errorf("unmarshal msg error: %v", err)
		return
	}
	xtrace.RunWithTrace(msgFromMQ.OperationID, func(ctx context.Context) {
		xtrace.StartFuncSpan(ctx, "MsgPushServiceServer.ConsumeSuperGroup", func(ctx context.Context) {
			_, err := s.PushSuperGroupMsg(ctx, msgFromMQ)
			if err == nil {
				sess.MarkMessage(msg, "")
			} else {
				logx.WithContext(ctx).Errorf("push super group msg error: %v", err)
			}
		})

	}, attribute.String("msg.key", string(msg.Key)))
}

func (s *MsgPushServiceServer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if msg.Topic == s.svcCtx.Config.SinglePushConsumer.Topic {
			s.ConsumeSingle(sess, msg)
		} else if msg.Topic == s.svcCtx.Config.SuperGroupPushConsumer.Topic {
			s.ConsumeSuperGroup(sess, msg)
		}
	}
	return nil
}

func NewMsgPushServiceServer(svcCtx *svc.ServiceContext) *MsgPushServiceServer {
	m := &MsgPushServiceServer{
		svcCtx: svcCtx,
		SingleConsumerGroup: xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{
			KafkaVersion:   sarama.V0_10_2_0,
			OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false,
		}, []string{svcCtx.Config.SinglePushConsumer.Topic},
			svcCtx.Config.SinglePushConsumer.Brokers, svcCtx.Config.SinglePushConsumer.SinglePushGroupID),
		SuperGroupConsumerGroup: xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{
			KafkaVersion:   sarama.V0_10_2_0,
			OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false,
		}, []string{svcCtx.Config.SuperGroupPushConsumer.Topic},
			svcCtx.Config.SuperGroupPushConsumer.Brokers, svcCtx.Config.SuperGroupPushConsumer.SuperGroupPushGroupID),
	}
	m.Subscribe()
	return m

}

func (s *MsgPushServiceServer) PushMsg(ctx context.Context, in *pushpb.PushMsgReq) (*pushpb.PushMsgResp, error) {
	l := logic.NewPushMsgLogic(ctx, s.svcCtx)
	return l.PushMsg(in)
}
func (s *MsgPushServiceServer) PushSuperGroupMsg(ctx context.Context, in *chatpb.PushMsgToSuperGroupDataToMQ) (*pushpb.PushMsgResp, error) {
	l := logic.NewPushMsgLogic(ctx, s.svcCtx)
	return l.PushSuperGroupMsg(in)
}

func (s *MsgPushServiceServer) Subscribe() {
	go s.SingleConsumerGroup.RegisterHandleAndConsumer(s)
	go s.SuperGroupConsumerGroup.RegisterHandleAndConsumer(s)
}
