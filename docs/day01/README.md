# 1、克隆Zero-IM-Server

```shell
git clone https://github.com/showurl/Zero-IM-Server.git -b main --depth 1
```

# 2、搜索 todo

## 2.1 im-user-rpc中 的 todo

### getGroupMemberIDListFromCacheLogic.go

> 从缓存取群成员列表

```go
// getGroupMemberIDListFromCacheLogic.go
func (l *GetGroupMemberIDListFromCacheLogic) GetGroupMemberIDListFromCache(in *pb.GetGroupMemberIDListFromCacheReq) (*pb.GetGroupMemberIDListFromCacheResp, error) {
// 如果你使用 Open-IM 的群聊功能 此处需要你实现
// 如果你仅仅使用 Zero-IM 的超级大群功能 你需要实现 GetUserListFromSuperGroupWithOpt rpc接口
// [🐶]我就不实现了 我准备用自己写的超级大群功能
}
```

### getSingleConversationRecvMsgOptsLogic.go

> 获取单聊会话的消息接收选项

#### 先把仓库加进来

##### 新增package repository

```go
// rep.go
package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/common/xcache"
	"github.com/showurl/Zero-IM-Server/common/xcache/dc"
	"github.com/showurl/Zero-IM-Server/common/xcache/global"
	"github.com/showurl/Zero-IM-Server/common/xorm"
	"gorm.io/gorm"
)

type Rep struct {
	svcCtx      *svc.ServiceContext
	Cache       redis.UniversalClient
	Mysql       *gorm.DB
	DetailCache *dc.DbMapping
}

var rep *Rep

func NewRep(svcCtx *svc.ServiceContext) *Rep {
	if rep != nil {
		return rep
	}
	rep = &Rep{
		svcCtx: svcCtx,
		Cache:  xcache.GetClient(svcCtx.Config.RedisConfig.RedisConf, global.DB(svcCtx.Config.RedisConfig.DB)),
		Mysql:  xorm.GetClient(svcCtx.Config.MysqlConfig),
	}
	rep.DetailCache = dc.GetDbMapping(rep.Cache, rep.Mysql)
	return rep
}

```

##### 新增 redis mongodb config

```go
// config/config.go
type Config struct {
zrpc.RpcServerConf
RedisConfig RedisConfig
MysqlConfig global.MysqlConfig
}
type RedisConfig struct {
redis.RedisConf
DB int
}
```

##### etc imuser.yaml新增配置

```yaml
RedisConfig:
  Host: 192.168.2.77:6379
  Pass: "123456"
  Type: node
  DB: 1

MysqlConfig:
  Addr: "root:123456@tcp(127.0.0.1:3306)/zeroim?charset=utf8mb4&parseTime=True&loc=Local&timeout=20s&readTimeout=20s&writeTimeout=20s"
  MaxIdleConns: 10
  MaxOpenConns: 10
  LogLevel: INFO
```

```go
// getSingleConversationRecvMsgOptsLogic.go
type GetSingleConversationRecvMsgOptsLogic struct {
ctx    context.Context
svcCtx *svc.ServiceContext
logx.Logger
rep *repository.Rep
}

func NewGetSingleConversationRecvMsgOptsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSingleConversationRecvMsgOptsLogic {
return &GetSingleConversationRecvMsgOptsLogic{
ctx:    ctx,
svcCtx: svcCtx,
Logger: logx.WithContext(ctx),
rep:    repository.NewRep(svcCtx),
}
}
```

#### model

```go
// model/single_conversation_record.go
package model

type SingleConversationRecord struct {
	ConversationId string `gorm:"column:conversation_id;primary_key;type:varchar(127);comment:主键 会话id;"`
	UserId         string `gorm:"column:user_id;type:char(64);comment:用户id;"`
	RecvMsgOpt     int8   `gorm:"column:recv_msg_opt;index;type:tinyint(1);comment:接收消息选项，0接收并提醒 1屏蔽消息 2接收但不提醒 默认0;default:0;"`
	Remark         string `gorm:"column:remark;type:varchar(255);comment:备注;default:'';"`
}

func (s *SingleConversationRecord) GetIdString() string {
	return s.ConversationId
}

func (s *SingleConversationRecord) TableName() string {
	// 分表
	return "single_conversation_record_" + s.UserId
}

```

#### logic

```go
package logic

//  获取单聊会话的消息接收选项
func (l *GetSingleConversationRecvMsgOptsLogic) GetSingleConversationRecvMsgOpts(in *pb.GetSingleConversationRecvMsgOptsReq) (*pb.GetSingleConversationRecvMsgOptsResp, error) {
	record := &model.SingleConversationRecord{}
	record.UserId = in.UserID
	record.ConversationId = in.ConversationID
	err := l.rep.DetailCache.FirstByID(
		record,
		dc.WithFieldId("conversation_id"),
	)
	if err != nil {
		if xormerr.RecordNotFound(err) || err == global.RedisErrorNotExists {
			err = nil
		} else if xormerr.TableNotFound(err) {
			l.rep.Mysql.Table(record.TableName()).AutoMigrate(record)
			return nil, err
		} else {
			return nil, err
		}
	}
	return &pb.GetSingleConversationRecvMsgOptsResp{
		CommonResp: &pb.CommonResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
		Opts: pb.RecvMsgOpt(record.RecvMsgOpt),
	}, nil
}

```

#### 单元测试

```go
package rpc

import (
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/imuserservice"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	conf = zrpc.RpcClientConf{
		Endpoints: []string{
			"192.168.2.77:10240",
		},
		Target:   "",
		App:      "",
		Token:    "",
		NonBlock: true,
		Timeout:  0,
	}
	imUserService = imuserservice.NewImUserService(zrpc.MustNewClient(conf))
)

func TestGetSingleConversationRecvMsgOptsLogic(t *testing.T) {
	resp, err := imUserService.GetSingleConversationRecvMsgOpts(
		context.Background(),
		&pb.GetSingleConversationRecvMsgOptsReq{
			UserID:         "1",
			ConversationID: "supergroup_0",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

```

### getUserListFromSuperGroupWithOptLogic.go

> 使用接收消息选项获取群聊中的用户列表

#### model

```go
package model

type SuperGroupConversationRecord struct {
	UserId     string `gorm:"column:user_id;primary_key;type:char(64);comment:主键 用户id;"`
	GroupId    string `gorm:"column:group_id;type:char(64);comment:群组id;"`
	RecvMsgOpt int8   `gorm:"column:recv_msg_opt;index;type:tinyint(1);comment:接收消息选项，0接收并提醒 1屏蔽消息 2接收但不提醒 默认0;default:0;"`
	Remark     string `gorm:"column:remark;type:varchar(255);comment:备注;default:'';"`
}

func (s *SuperGroupConversationRecord) GetIdString() string {
	return s.UserId
}

func (s *SuperGroupConversationRecord) TableName() string {
	return "supergroup_conversation_record_" + s.GroupId
}
```

#### repository.go

```go
rep.RelationCache = rc.NewRelationMapping(rep.Mysql, rep.Cache)
```

#### getUserListFromSuperGroupWithOptLogic.go

```go
package logic

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Zero-IM-Server/app/im-user/model"
	"github.com/showurl/Zero-IM-Server/common/xcache/global"
	"github.com/showurl/Zero-IM-Server/common/xcache/rc"
	xormerr "github.com/showurl/Zero-IM-Server/common/xorm/err"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListFromSuperGroupWithOptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetUserListFromSuperGroupWithOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListFromSuperGroupWithOptLogic {
	return &GetUserListFromSuperGroupWithOptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

//  获取超级群成员列表 通过消息接收选项
func (l *GetUserListFromSuperGroupWithOptLogic) GetUserListFromSuperGroupWithOpt(in *pb.GetUserListFromSuperGroupWithOptReq) (*pb.GetUserListFromSuperGroupWithOptResp, error) {
	resp := &pb.GetUserListFromSuperGroupWithOptResp{
		CommonResp:    &pb.CommonResp{},
		UserIDOptList: nil,
	}
	record := &model.SuperGroupConversationRecord{}
	record.GroupId = in.SuperGroupID
	for _, opt := range in.Opts {
		var userIds []string
		err := l.rep.RelationCache.List(
			&userIds,
			0,
			-1,
			record,
			"user_id",
			map[string]interface{}{
				"recv_msg_opt": opt,
			},
			rc.Order("user_id"),
		)
		if err != nil {
			if xormerr.TableNotFound(err) {
				l.rep.Mysql.Table(record.TableName()).AutoMigrate(record)
			}
			if global.RedisErrorNotExists == err {
				continue
			}
			return nil, err
		}
		for _, id := range userIds {
			resp.UserIDOptList = append(resp.UserIDOptList, &pb.UserIDOpt{
				UserID: id,
				Opts:   opt,
			})
		}
	}
	return resp, nil
}

```

#### 单元测试

```go
package rpc

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestGetUserListFromSuperGroupWithOpt(t *testing.T) {
	resp, err := imUserService.GetUserListFromSuperGroupWithOpt(
		context.Background(),
		&pb.GetUserListFromSuperGroupWithOptReq{
			SuperGroupID: "supergroup_0",
			Opts: []pb.RecvMsgOpt{
				pb.RecvMsgOpt_ReceiveMessage,
				pb.RecvMsgOpt_ReceiveNotNotifyMessage,
			},
			//UserIDList:   []string{"user_0", "user_1"},
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

```

### ifAInBBlacklistLogic.go

> 判断A是否在B的黑名单中

#### model

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

#### logic

```go
package logic

//  判断用户A是否在B黑名单中
func (l *IfAInBBlacklistLogic) IfAInBBlacklist(in *pb.IfAInBBlacklistReq) (*pb.IfAInBBlacklistResp, error) {
	blacklist := &model.Blacklist{}
	blacklist.SelfId = in.BUserID
	exist, err := l.rep.RelationCache.Exist(in.AUserID, blacklist, "user_id", map[string]interface{}{})
	if err != nil {
		if xormerr.TableNotFound(err) {
			_ = l.rep.Mysql.Table(blacklist.TableName()).AutoMigrate(blacklist)
		}
		return nil, err
	}
	return &pb.IfAInBBlacklistResp{
		CommonResp: &pb.CommonResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
		IsInBlacklist: exist,
	}, nil
}

```

#### 单元测试

```go
package rpc

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestIfAInBBlacklist(t *testing.T) {
	resp, err := imUserService.IfAInBBlacklist(
		context.Background(),
		&pb.IfAInBBlacklistReq{
			AUserID: "a",
			BUserID: "b",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
```

### ifAInBFriendListLogic.go 和 ifAInBBlacklistLogic.go 差不多 复制粘贴改改即可

### verifyTokenLogic.go
> 验证token

#### token 组成部分 允许同平台单账号同时登录
- redis key：`userId` + `platform`
- redis hash key：`token`
- redis hash value：`int(暂时无意义)`

#### repository
```go
package repository

import (
	"context"
	"github.com/showurl/Zero-IM-Server/common/types"
	timeUtils "github.com/showurl/Zero-IM-Server/common/utils/time"
	"strconv"
	"time"
)

func (r *Rep) GetTokenMap(ctx context.Context, uid string, platform string) (map[string]int, error) {
	key := types.RedisKeyToken + uid + ":" + platform
	result, err := r.Cache.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]int{}, err
	}
	tokenMap := make(map[string]int)
	for k, v := range result {
		tokenMap[k], _ = strconv.Atoi(v)
	}
	return tokenMap, nil
}

func (r *Rep) RenewalToken(ctx context.Context, uid string, platform string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	return r.Cache.ExpireAt(ctx, key, timeUtils.Now().Add(time.Hour*24*time.Duration(r.svcCtx.Config.TokenRenewalDay))).Err()
}

```

#### logic
```go
package logic

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/repository"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	jwtUtils "github.com/showurl/Zero-IM-Server/common/utils/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

// 检查token
func (l *VerifyTokenLogic) VerifyToken(in *pb.VerifyTokenReq) (*pb.VerifyTokenResp, error) {
	claim, err := jwtUtils.GetClaimFromToken(in.Token, l.svcCtx.Config.TokenSecret)
	if err != nil {
		return nil, err
	}
	tokenMap, err := l.rep.GetTokenMap(l.ctx, claim.UID, claim.Platform)
	if err != nil {
		return nil, err
	}
	if _, ok := tokenMap[in.Token]; ok {
		// 有的话就更新过期时间
		go func() {
			_ = l.rep.RenewalToken(l.ctx, claim.UID, claim.Platform)
		}()
		return &pb.VerifyTokenResp{
			Uid:     claim.UID,
			Success: true,
			ErrMsg:  "",
		}, nil
	} else {
		// 没有 token 提示
		return &pb.VerifyTokenResp{
			Uid:     "",
			Success: false,
			ErrMsg:  "token is not exist",
		}, nil
	}
}

```

## 2.2 msg-callback 中的 todo 
> 消息发送前后的回调 

> before sendmsg 可以返回错误信息 使msg发送失败

> after sendmsg 可以记录消息内容 达到双写的目的

> before send text msg 可以过滤消息文本 进行违禁词替换