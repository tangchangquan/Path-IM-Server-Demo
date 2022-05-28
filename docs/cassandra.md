# 使用 cassandra 代替 mongodb
## 一、app/msg-transfer/cmd/history-cassandra 代替 app/msg-transfer/cmd/history-mongo
```go

func (r *Rep) SaveUserChatCassandra2(spanCtx context.Context, uid string, sendTime int64, m *chatpb.MsgDataToDB) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(r.svcCtx.Config.Cassandra.TimeoutSecond)*time.Second)
	seqUid := getSeqUid(uid, m.MsgData.Seq)
	{
		err := r.CassandraClient.Query(
			fmt.Sprintf(
				`UPDATE %s.%s SET msgs = msgs + ? WHERE uid = ?`,
				r.svcCtx.Config.Cassandra.Keyspace,
				r.svcCtx.Config.Cassandra.SingleChatMsgTableName,
			),
			[]map[int64][]byte{{
				sendTime: m.MsgData.Bytes(),
			}}, seqUid,
		).WithContext(ctx).Exec()
		if err != nil {
			logx.WithContext(spanCtx).Errorf("SaveUserChatCassandra2 error: %v", err)
			return err
		}
	}
	return nil
}

func (r *Rep) SaveSuperGroupChatCassandra2(spanCtx context.Context, groupId string, sendTime int64, m *chatpb.MsgDataToDB) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(r.svcCtx.Config.Cassandra.TimeoutSecond)*time.Second)
	seqGroupId := getSeqGroupId(groupId, m.MsgData.Seq)
	{
		err := r.CassandraClient.Query(
			fmt.Sprintf(
				`UPDATE %s.%s SET msgs = msgs + ? WHERE groupid = ?`,
				r.svcCtx.Config.Cassandra.Keyspace,
				r.svcCtx.Config.Cassandra.SuperGroupChatMsgTableName,
			),
			[]map[int64][]byte{{
				sendTime: m.MsgData.Bytes(),
			}}, seqGroupId,
		).WithContext(ctx).Exec()
		if err != nil {
			logx.WithContext(spanCtx).Errorf("SaveSuperGroupChatCassandra2 error: %v", err)
			return err
		}
	}
	return nil
}

```

> cassandra update时 不需要验证是否存在，因为cassandra会自动创建
 
## 二、msg 新增cassandra 实现 interface
```go
package rep
type IPullHistoryMsg interface {
	GetMsgBySeqList(uid string, seqList []uint32) (seqMsg []*pb.MsgData, err error)
	GetMsgBySuperGroupSeqList(groupId string, seqList []uint32) (seqMsg []*pb.MsgData, err error)
}

type CassandraHistory struct {
	svcCtx          *svc.ServiceContext
	CassandraClient *gocql.Session
}


func (r *CassandraHistory) GetMsgBySeqList(uid string, seqList []uint32) (seqMsg []*pb.MsgData, err error) {
	var hasSeqList []uint32
	singleCount := 0
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(r.svcCtx.Config.Cassandra.TimeoutSecond)*time.Second)
	m := func(uid string, seqList []uint32) map[string][]uint32 {
		t := make(map[string][]uint32)
		for i := 0; i < len(seqList); i++ {
			seqUid := getSeqUid(uid, seqList[i])
			if value, ok := t[seqUid]; !ok {
				var temp []uint32
				t[seqUid] = append(temp, seqList[i])
			} else {
				t[seqUid] = append(value, seqList[i])
			}
		}
		return t
	}(uid, seqList)

	for seqUid, value := range m {
		sChat := &model.CassUserChat{}
		if err = r.CassandraClient.Query(fmt.Sprintf(
			`SELECT uid, msgs FROM %s.%s WHERE uid = ?`,
			r.svcCtx.Config.Cassandra.Keyspace,
			r.svcCtx.Config.Cassandra.SingleChatMsgTableName,
		), seqUid).WithContext(ctx).Scan(&sChat.Uid, &sChat.Msgs); err != nil {
			logx.Error("findone uid failed:", err.Error())
			continue
		}
		singleCount = 0
		for i := 0; i < len(sChat.Msgs); i++ {
			msg := new(pb.MsgData)
			msgMap := sChat.Msgs[i]
			var bytes []byte
			for _, bytes = range msgMap {
				break
			}
			err = proto.Unmarshal(bytes, msg)
			if err != nil {
				logx.Error("Unmarshal failed:", err.Error())
				return nil, err
			}
			if numUtils.IsContainUInt32(msg.Seq, value) {
				seqMsg = append(seqMsg, msg)
				hasSeqList = append(hasSeqList, msg.Seq)
				singleCount++
				if singleCount == len(value) {
					break
				}
			}
		}
	}
	if len(hasSeqList) != len(seqList) {
		var diff []uint32
		diff = numUtils.DifferenceUInt32(hasSeqList, seqList)
		exceptionMSg := genExceptionMessageBySeqList(diff)
		seqMsg = append(seqMsg, exceptionMSg...)

	}
	return seqMsg, nil
}

func (r *CassandraHistory) GetMsgBySuperGroupSeqList(groupId string, seqList []uint32) (seqMsg []*pb.MsgData, err error) {
	var hasSeqList []uint32
	singleCount := 0
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(r.svcCtx.Config.Cassandra.TimeoutSecond)*time.Second)
	m := func(groupid string, seqList []uint32) map[string][]uint32 {
		t := make(map[string][]uint32)
		for i := 0; i < len(seqList); i++ {
			seqGroupId := getSeqGroupId(groupid, seqList[i])
			if value, ok := t[seqGroupId]; !ok {
				var temp []uint32
				t[seqGroupId] = append(temp, seqList[i])
			} else {
				t[seqGroupId] = append(value, seqList[i])
			}
		}
		return t
	}(groupId, seqList)

	for seqUid, value := range m {
		sChat := &model.CassGroupChat{}
		if err = r.CassandraClient.Query(fmt.Sprintf(
			`SELECT groupid, msgs FROM %s.%s WHERE groupid = ?`,
			r.svcCtx.Config.Cassandra.Keyspace,
			r.svcCtx.Config.Cassandra.SuperGroupChatMsgTableName,
		), seqUid).WithContext(ctx).Scan(&sChat.Groupid, &sChat.Msgs); err != nil {
			logx.Error("findone groupid failed:", err.Error())
			continue
		}
		singleCount = 0
		for i := 0; i < len(sChat.Msgs); i++ {
			msg := new(pb.MsgData)
			msgMap := sChat.Msgs[i]
			var bytes []byte
			for _, bytes = range msgMap {
				break
			}
			err = proto.Unmarshal(bytes, msg)
			if err != nil {
				logx.Error("Unmarshal failed:", err.Error())
				return nil, err
			}
			if numUtils.IsContainUInt32(msg.Seq, value) {
				seqMsg = append(seqMsg, msg)
				hasSeqList = append(hasSeqList, msg.Seq)
				singleCount++
				if singleCount == len(value) {
					break
				}
			}
		}
	}
	if len(hasSeqList) != len(seqList) {
		var diff []uint32
		diff = numUtils.DifferenceUInt32(hasSeqList, seqList)
		exceptionMSg := genExceptionMessageBySeqList(diff)
		seqMsg = append(seqMsg, exceptionMSg...)

	}
	return seqMsg, nil
}

```