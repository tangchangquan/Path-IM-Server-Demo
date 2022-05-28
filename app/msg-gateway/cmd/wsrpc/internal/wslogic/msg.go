package wslogic

import (
	"context"
	imuserpb "github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/rpc/pb"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg-gateway/cmd/wsrpc/pb"
	"github.com/Path-IM/Path-IM-Server-Demo/app/msg/cmd/rpc/chat"
	chatpb "github.com/Path-IM/Path-IM-Server-Demo/app/msg/cmd/rpc/pb"
	"github.com/Path-IM/Path-IM-Server-Demo/common/types"
	"github.com/Path-IM/Path-IM-Server-Demo/common/utils"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xerr"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xtrace"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *MsggatewayLogic) msgParse(ctx context.Context, conn *UserConn, binaryMsg []byte) {
	m := &pb.Req{}
	err := proto.Unmarshal(binaryMsg, m)
	if err != nil {
		l.sendErrMsg(ctx, conn, 200, err.Error(), types.WSDataError, "")
		err = conn.Close()
		if err != nil {
			logx.WithContext(ctx).Error("ws close err", err.Error())
		}
		return
	}
	if err := validate.Struct(m); err != nil {
		logx.WithContext(ctx).Error("ws args validate  err", err.Error())
		l.sendErrMsg(ctx, conn, 201, err.Error(), xerr.NewErrCode(int(m.ReqIdentifier)), m.MsgIncr)
		return
	}
	if !l.svcCtx.Config.VerifyTokenOnce {
		tokenInvalid := false
		tokenInvalidMsg := ""
		xtrace.StartFuncSpan(ctx, utils.CallerFuncName(), func(ctx context.Context) {
			uid, platform := l.getUserUid(conn)
			tokenResp, err1 := l.svcCtx.ImUserService.VerifyToken(ctx, &imuserpb.VerifyTokenReq{
				Token:    m.Token,
				Platform: platform,
				SendID:   uid,
			})
			if err1 != nil {
				tokenInvalid = true
				tokenInvalidMsg = "server error"
				logx.WithContext(ctx).Error("ws verify token err", err1.Error())
				return
			}
			if !tokenResp.Success {
				tokenInvalid = true
				tokenInvalidMsg = tokenResp.ErrMsg
				logx.WithContext(ctx).Error("ws verify token err", "token invalid")
				return
			}
		})
		if tokenInvalid {
			l.sendErrMsg(ctx, conn, 202, tokenInvalidMsg, xerr.NewErrCode(int(m.ReqIdentifier)), m.MsgIncr)
			return
		}
	}
	switch m.ReqIdentifier {
	case types.WSGetNewestSeq:
		l.getSeqReq(ctx, conn, m)
	case types.WSGetNewestGroupSeq:
		l.getSuperGroupSeqReq(ctx, conn, m)
	case types.WSSendMsg:
		l.sendMsgReq(ctx, conn, m)
	case types.WSPullMsgBySeqList:
		l.pullMsgBySeqListReq(ctx, conn, m)
	case types.WSPullMsgByGroupSeqList:
		l.pullMsgBySuperGroupSeqListReq(ctx, conn, m)
	default:
	}
}

func (l *MsggatewayLogic) sendSignalMsgResp(ctx context.Context, conn *UserConn, errCode int32, errMsg string, m *Req) {
	mReply := Resp{
		ReqIdentifier: m.ReqIdentifier,
		MsgIncr:       m.MsgIncr,
		ErrCode:       errCode,
		ErrMsg:        errMsg,
		Data:          nil,
	}
	l.sendMsg(ctx, conn, mReply)
}

func (l *MsggatewayLogic) sendErrMsg(ctx context.Context, conn *UserConn, code int32, errMsg string, reqIdentifier *xerr.CodeError, msgIncr string) {
	mReply := Resp{
		ReqIdentifier: int32(reqIdentifier.GetErrCode()),
		MsgIncr:       msgIncr,
		ErrCode:       code,
		ErrMsg:        errMsg,
	}
	l.sendMsg(ctx, conn, mReply)
}

func (l *MsggatewayLogic) sendMsg(ctx context.Context, conn *UserConn, mReply Resp) {
	resp := &pb.Resp{
		ReqIdentifier: uint32(mReply.ReqIdentifier),
		MsgIncr:       mReply.MsgIncr,
		ErrCode:       uint32(mReply.ErrCode),
		ErrMsg:        mReply.ErrMsg,
		Data:          mReply.Data,
	}
	b, err := proto.Marshal(resp)
	if err != nil {
		uid, platform := l.getUserUid(conn)
		logx.WithContext(ctx).Error(mReply.ReqIdentifier, mReply.ErrCode, mReply.ErrMsg, "Encode Msg error", conn.RemoteAddr().String(), uid, platform, err.Error())
		return
	}
	err = l.writeMsg(conn, websocket.BinaryMessage, b)
	if err != nil {
		uid, platform := l.getUserUid(conn)
		logx.WithContext(ctx).Error(mReply.ReqIdentifier, mReply.ErrCode, mReply.ErrMsg, "WS WriteMsg error", conn.RemoteAddr().String(), uid, platform, err.Error())
	}
}

func (l *MsggatewayLogic) writeMsg(conn *UserConn, a int, msg []byte) error {
	conn.w.Lock()
	defer conn.w.Unlock()
	return conn.WriteMessage(a, msg)
}

func (l *MsggatewayLogic) sendMsgReq(ctx context.Context, conn *UserConn, m *pb.Req) {
	sendMsgAllCount++
	logx.WithContext(ctx).Info("Ws call success to sendMsgReq start", m.MsgIncr, m.ReqIdentifier, m.SendID, m.Data)
	nReply := new(chatpb.SendMsgResp)
	isPass, errCode, errMsg, pData := l.argsValidate(m, types.WSSendMsg)
	if isPass {
		data := pData.(chatpb.MsgData)
		pbData := chatpb.SendMsgReq{
			Token:   m.Token,
			MsgData: &data,
		}
		logx.WithContext(ctx).Info("Ws call success to sendMsgReq middle", m.ReqIdentifier, m.SendID, m.MsgIncr, data)

		reply, err := l.svcCtx.MsgRpc.SendMsg(ctx, &pbData)
		if err != nil {
			logx.WithContext(ctx).Error("UserSendMsg err ", err.Error())
			nReply.ErrCode = 200
			nReply.ErrMsg = err.Error()
			l.sendMsgResp(ctx, conn, m, nReply)
		} else {
			logx.WithContext(ctx).Info("rpc call success to sendMsgReq", reply.String())
			l.sendMsgResp(ctx, conn, m, reply)
		}
	} else {
		nReply.ErrCode = errCode
		nReply.ErrMsg = errMsg
		l.sendMsgResp(ctx, conn, m, nReply)
	}
}

func (l *MsggatewayLogic) sendMsgResp(ctx context.Context, conn *UserConn, m *pb.Req, pb *chat.SendMsgResp) {
	var mReplyData chatpb.UserSendMsgResp
	mReplyData.ClientMsgID = pb.GetClientMsgID()
	mReplyData.ServerMsgID = pb.GetServerMsgID()
	mReplyData.SendTime = pb.GetSendTime()
	b, _ := proto.Marshal(&mReplyData)
	mReply := Resp{
		ReqIdentifier: int32(m.ReqIdentifier),
		MsgIncr:       m.MsgIncr,
		ErrCode:       pb.GetErrCode(),
		ErrMsg:        pb.GetErrMsg(),
		Data:          b,
	}
	l.sendMsg(ctx, conn, mReply)
}

func (l *MsggatewayLogic) pullMsgBySeqListReq(ctx context.Context, conn *UserConn, m *pb.Req) {
	logx.WithContext(ctx).Info("Ws call success to pullMsgBySeqListReq start", m.SendID, m.ReqIdentifier, m.MsgIncr, m.Data)
	nReply := new(chatpb.PullMessageBySeqListResp)
	isPass, errCode, errMsg, data := l.argsValidate(m, types.WSPullMsgBySeqList)
	if isPass {
		rpcReq := chatpb.PullMessageBySeqListReq{}
		rpcReq.SeqList = data.(chatpb.PullMessageBySeqListReq).SeqList
		rpcReq.UserID = m.SendID
		logx.WithContext(ctx).Info("Ws call success to pullMsgBySeqListReq middle", m.SendID, m.ReqIdentifier, m.MsgIncr, data.(chatpb.PullMessageBySeqListReq).SeqList)
		reply, err := l.svcCtx.MsgRpc.PullMessageBySeqList(ctx, &chatpb.WrapPullMessageBySeqListReq{PullMessageBySeqListReq: &rpcReq})
		if err != nil {
			logx.WithContext(ctx).Errorf("pullMsgBySeqListReq err", err.Error())
			nReply.ErrCode = 200
			nReply.ErrMsg = err.Error()
			l.pullMsgBySeqListResp(ctx, conn, m, nReply)
		} else {
			logx.WithContext(ctx).Info("rpc call success to pullMsgBySeqListReq ", reply.String())
			l.pullMsgBySeqListResp(ctx, conn, m, reply.PullMessageBySeqListResp)
		}
	} else {
		nReply.ErrCode = errCode
		nReply.ErrMsg = errMsg
		l.pullMsgBySeqListResp(ctx, conn, m, nReply)
	}
}

func (l *MsggatewayLogic) pullMsgBySuperGroupSeqListReq(ctx context.Context, conn *UserConn, m *pb.Req) {
	logx.WithContext(ctx).Info("Ws call success to pullMsgBySuperGroupSeqListReq start", m.SendID, m.ReqIdentifier, m.MsgIncr, m.Data)
	nReply := new(chatpb.PullMessageBySeqListResp)
	isPass, errCode, errMsg, data := l.argsValidate(m, types.WSPullMsgByGroupSeqList)
	if isPass {
		rpcReq := chatpb.PullMessageBySuperGroupSeqListReq{}
		rpcReq.SeqList = data.(chatpb.PullMessageBySuperGroupSeqListReq).SeqList
		rpcReq.GroupID = data.(chatpb.PullMessageBySuperGroupSeqListReq).GroupID
		logx.WithContext(ctx).Info("Ws call success to pullMsgBySeqListReq middle", m.SendID, m.ReqIdentifier, m.MsgIncr, data.(chatpb.PullMessageBySuperGroupSeqListReq).SeqList)
		reply, err := l.svcCtx.MsgRpc.PullMessageBySuperGroupSeqList(ctx, &rpcReq)
		if err != nil {
			logx.WithContext(ctx).Errorf("pullMsgBySeqListReq err", err.Error())
			nReply.ErrCode = 200
			nReply.ErrMsg = err.Error()
			l.pullMsgBySuperGroupSeqListResp(ctx, conn, m, nReply)
		} else {
			logx.WithContext(ctx).Info("rpc call success to pullMsgBySeqListReq ", reply.String())
			l.pullMsgBySuperGroupSeqListResp(ctx, conn, m, reply.PullMessageBySeqListResp)
		}
	} else {
		nReply.ErrCode = errCode
		nReply.ErrMsg = errMsg
		l.pullMsgBySuperGroupSeqListResp(ctx, conn, m, nReply)
	}
}

func (l *MsggatewayLogic) pullMsgBySeqListResp(ctx context.Context, conn *UserConn, m *pb.Req, pb *chatpb.PullMessageBySeqListResp) {
	logx.WithContext(ctx).Info("pullMsgBySeqListResp come  here ", pb.String())
	c, _ := proto.Marshal(pb)
	mReply := Resp{
		ReqIdentifier: int32(m.ReqIdentifier),
		MsgIncr:       m.MsgIncr,
		ErrCode:       pb.GetErrCode(),
		ErrMsg:        pb.GetErrMsg(),
		Data:          c,
	}
	logx.WithContext(ctx).Info("pullMsgBySeqListResp all data  is ", mReply.ReqIdentifier, mReply.MsgIncr, mReply.ErrCode, mReply.ErrMsg,
		len(mReply.Data))

	l.sendMsg(ctx, conn, mReply)
}

func (l *MsggatewayLogic) pullMsgBySuperGroupSeqListResp(ctx context.Context, conn *UserConn, m *pb.Req, pb *chatpb.PullMessageBySeqListResp) {
	logx.WithContext(ctx).Info("pullMsgBySuperGroupSeqListResp come  here ", pb.String())
	c, _ := proto.Marshal(pb)
	mReply := Resp{
		ReqIdentifier: int32(m.ReqIdentifier),
		MsgIncr:       m.MsgIncr,
		ErrCode:       pb.GetErrCode(),
		ErrMsg:        pb.GetErrMsg(),
		Data:          c,
	}
	logx.WithContext(ctx).Info("pullMsgBySeqListResp all data  is ", mReply.ReqIdentifier, mReply.MsgIncr, mReply.ErrCode, mReply.ErrMsg,
		len(mReply.Data))

	l.sendMsg(ctx, conn, mReply)
}
