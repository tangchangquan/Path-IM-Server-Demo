# 用户关系模块

## model
### 好友
```go
package model

type Friendlist struct {
	UserId string `gorm:"column:user_id;primary_key;type:char(64);comment:主键 用户id;"`
	SelfId string `gorm:"column:self_id;type:char(64);comment:自己id,分表键;"`
	// 加好友时间
	CreatedAt int64 `gorm:"column:created_at;type:bigint(10);comment:加好友时间 秒级时间戳;"`
}

func (b *Friendlist) TableName() string {
	return "friendlist_" + b.SelfId
}

```
### 黑名单
```go
package model

type Blacklist struct {
	UserId string `gorm:"column:user_id;primary_key;type:char(64);comment:主键 用户id;"`
	SelfId string `gorm:"column:self_id;type:char(64);comment:自己id,分表键;"`
	// 拉黑时间
	CreatedAt int64 `gorm:"column:created_at;type:bigint(10);comment:拉黑时间 秒级时间戳;"`
}

func (b *Blacklist) TableName() string {
	return "blacklist_" + b.SelfId
}

```

### 好友申请
```go
package model

type FriendApplyRecord struct {
	Id     string `gorm:"column:id;primary_key;type:char(64);comment:主键"`
	UserId string `gorm:"column:user_id;index;type:char(64);comment:主键 用户id 申请人的id;"`
	SelfId string `gorm:"column:self_id;type:char(64);comment:自己id,分表键;"`
	// 申请时间
	CreatedAt int64 `gorm:"column:created_at;type:bigint(10);comment:加好友时间 秒级时间戳;"`
	// 申请状态
	Status int `gorm:"column:status;type:tinyint(1);comment:申请状态 0:申请中 1:同意 2:拒绝;"`
	// 申请消息
	Message string `gorm:"column:message;type:varchar(255);comment:申请消息;"`
}

func (f *FriendApplyRecord) TableName() string {
	return "friend_apply_record_" + f.SelfId
}

```

## proto
```protobuf
syntax = "proto3";

option go_package = "./pb";

package relation;
message PageReq {
  int32 page = 1;
  int32 page_size = 2;
}
message PageResp {
  int32 page = 1;
  int32 page_size = 2;
  int64 total = 3;
}
message RelationBaseResp {
  enum Errno {
    OK = 0;
    ERROR = 1;
  }
  Errno errCode = 1;
  string errMsg = 2;
}
message ApplyFriendReq {
  string applyUserId = 1; // 申请人
  string userId = 2; // 被申请人
  string message = 3; // 申请信息
}
message ApplyFriendResp {
  RelationBaseResp baseResp = 1;
}
message UpdateApplyFriendStatusReq {
  enum Status {
    AGREE = 0;
    REFUSE = 1;
  }
  Status status = 1;
  string id = 2; // 申请记录id
  string selfId = 3; // 自己的id
}
message UpdateApplyFriendStatusResp {
  RelationBaseResp baseResp = 1;
}
message GetApplyFriendListReq {
  string selfId = 1; // 用户id
  PageReq pageReq = 2;
}
message GetApplyFriendListResp {
  RelationBaseResp baseResp = 1;
  message ApplyFriend {
    string id = 1; // 申请记录id
    string applyUserId = 2; // 申请人
    string userId = 3; // 被申请人
    string message = 4; // 申请信息
    string createTime = 5; // 申请时间
    enum Status {
      Applying = 0;
      AGREE = 1;
      REFUSE = 2;
    }
    Status status = 6; // 申请状态
  }
  PageResp pageResp = 2;
  repeated ApplyFriend applyFriendList = 3;
}
message GetFriendListReq {
  string selfId = 1; // 用户id
}
message GetFriendListResp {
  RelationBaseResp baseResp = 1;
  message Friend {
    string userId = 1; // 用户id
    string selfId = 2; // 自己的id
    string remark = 3; // 备注
    string createTime = 4; // 加好友的时间
  }
  repeated Friend friendList = 2;
}
message DeleteFriendReq {
  string selfId = 1; // 用户id
  string userId = 2; // 好友id
}
message DeleteFriendResp {
  RelationBaseResp baseResp = 1;
}
message GetBlackListReq {
  string selfId = 1; // 用户id
  PageReq pageReq = 2;
}
message GetBlackListResp {
  RelationBaseResp baseResp = 1;
  message Black {
    string userId = 1; // 用户id
    string selfId = 2; // 自己的id
    string createTime = 3; // 加黑的时间 排序键
  }
  PageResp pageResp = 2;
  repeated Black blackList = 3;
}
message DeleteBlackReq {
  string selfId = 1; // 用户id
  string userId = 2; // 好友id
}
message DeleteBlackResp {
  RelationBaseResp baseResp = 1;
}
message BlackUserReq {
  string selfId = 1; // 用户id
  string userId = 2; // 好友id
}
message BlackUserResp {
  RelationBaseResp baseResp = 1;
}
service relationService {
  rpc ApplyFriend (ApplyFriendReq) returns (ApplyFriendResp);
  rpc UpdateApplyFriendStatus (UpdateApplyFriendStatusReq) returns (UpdateApplyFriendStatusResp);
  rpc GetApplyFriendList (GetApplyFriendListReq) returns (GetApplyFriendListResp);
  rpc GetFriendList (GetFriendListReq) returns (GetFriendListResp);
  rpc DeleteFriend (DeleteFriendReq) returns (DeleteFriendResp);
  rpc GetBlackList (GetBlackListReq) returns (GetBlackListResp);
  rpc DeleteBlack (DeleteBlackReq) returns (DeleteBlackResp);
  rpc BlackUser (BlackUserReq) returns (BlackUserResp);
}

```

## logic
### curd就不写了 主要是消息通知
#### 申请成为好友 发送通知
```go
package logic
func (l *ApplyFriendLogic) ApplyFriend(in *pb.ApplyFriendReq) (*pb.ApplyFriendResp, error) {
	m := &model.FriendApplyRecord{}
	m.SelfId = in.UserId
	count := int64(0)
	l.rep.Mysql.Model(m).Table(m.TableName()).
		Where("user_id = ?", in.ApplyUserId).
		Where("status = ?", 0).Count(&count)
	if count > 0 {
		return &pb.ApplyFriendResp{
			BaseResp: &pb.RelationBaseResp{},
		}, nil
	}
	record := &model.FriendApplyRecord{
		Id:        global.GetID(),
		UserId:    in.ApplyUserId,
		SelfId:    in.UserId,
		CreatedAt: timeUtils.Now().Unix(),
		Status:    0,
		Message:   in.Message,
	}
	err := xorm.Transaction(l.rep.Mysql,
		// 插入记录
		func(tx *gorm.DB) error {
			return xorm.Insert(tx, record)
		},
		// 发送通知
		func(_ *gorm.DB) error {
			_, err := l.svcCtx.MsgRpc().SendMsg(
				l.ctx,
				&chatpb.SendMsgReq{
					MsgData: types.NewSingleChatMsgNotification(
						in.ApplyUserId,
						in.UserId,
						types.NewApplyFriendContent(),
						types.ApplyFriendContentType,
						&chatpb.OfflinePushInfo{
							Title: "你有新的好友申请",
						},
					),
				},
			)
			if err != nil {
				return err
			}
			return nil
		},
	)
	return &pb.ApplyFriendResp{BaseResp: &pb.RelationBaseResp{
		ErrCode: 0,
		ErrMsg:  "",
	}}, err
}

```