# 用户登录功能
## 增加user.proto
### 使用用户名和密码注册
### 使用用户名和密码登录
### 获取用户信息
### 修改用户信息
### 最终
```protobuf
syntax = "proto3";

option go_package = "./pb";

package user;

message User {
  string Id = 1;
  string Username = 2;
  string Password = 3;
  string Nickname = 4;
  string Sign = 5;
  string Avatar = 6;
  string Province = 7;
  string City = 8;
  string District = 9;
  string Birthday = 10;
  string RegisterTime = 11;
  bool IsMale = 12;
}
message BaseResp {
  enum Errno {
    OK = 0;
    ERROR = 1;
  }
  Errno errCode = 1;
  string errMsg = 2;
}
message GetUserByIdReq {
  string userId = 1;
}
message GetUserByUsernameReq {
  string username = 1;
}
message GetUserResp {
  User user = 1;
  BaseResp baseResp = 2;
}
message InsertUserReq {
  User user = 1;
}
message InsertUserResp {
  BaseResp baseResp = 1;
}
message UpdateUserReq {
  User user = 1;
  repeated string fields = 2;
}
message UpdateUserResp {
  BaseResp baseResp = 1;
}
message LoginByPasswordReq {
  string username = 2;
  string password = 3;
}
message LoginByIdReq {
  string userId = 1;
}
message LoginResp {
  User user = 1;
  string token = 2;
  BaseResp baseResp = 3;
}

service userService {
  rpc GetUserById(GetUserByIdReq) returns (GetUserResp);
  rpc GetUserByUsername(GetUserByUsernameReq) returns (GetUserResp);
  rpc InsertUser(InsertUserReq) returns (InsertUserResp);
  rpc UpdateUser(UpdateUserReq) returns (UpdateUserResp);
  rpc LoginByPassword(LoginByPasswordReq) returns (LoginResp);
  rpc LoginById(LoginByIdReq) returns (LoginResp);
}

```
### 执行命令
```shell
goctl rpc protoc user.proto -v --go_out=.. --go-grpc_out=..  --zrpc_out=.. --style=goZero --home ../../../../../goctl/home
```

## 注册
### logic
```go
package logic

func (l *InsertUserLogic) InsertUser(in *pb.InsertUserReq) (*pb.InsertUserResp, error) {
	user := &model.User{
		Id:           in.User.Id,
		Username:     in.User.Username,
		Password:     in.User.Password,
		Nickname:     in.User.Nickname,
		Sign:         in.User.Sign,
		Avatar:       in.User.Avatar,
		Province:     in.User.Province,
		City:         in.User.City,
		District:     in.User.District,
		Birthday:     in.User.Birthday,
		RegisterTime: in.User.RegisterTime,
		IsMale:       in.User.IsMale,
	}
	err := xorm.Transaction(l.rep.Mysql,
		// 插入数据
		func(tx *gorm.DB) error {
			return xorm.Insert(tx, user)
		},
		// 清除缓存
		l.rep.FuncDelUserCache(l.ctx, user),
	)
	if err != nil {
		return &pb.InsertUserResp{
			BaseResp: &pb.BaseResp{
				ErrCode: -1,
				ErrMsg:  err.Error(),
			},
		}, err
	}
	// 预热数据
	go l.rep.WarmUpUser(l.ctx, user)
	return &pb.InsertUserResp{
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}

```
### repository
```go
package repository

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

func (r *Rep) WarmUpUser(ctx context.Context, user *model.User) {
	var err error
	err = r.DetailCache.FirstById(user, user.Id)
	if err != nil {
		logx.WithContext(ctx).Errorf("WarmUpUser: %s", err)
	}
	err = r.DetailCache.FirstByWhere(user, map[string][]interface{}{
		"username = ?": {user.Username},
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("WarmUpUser: %s", err)
	}
}
func (r *Rep) FuncDelUserCache(ctx context.Context, users ...*model.User) func(tx *gorm.DB) error {
	if len(users) == 0 {
		return func(tx *gorm.DB) error {
			return nil
		}
	}
	return func(_ *gorm.DB) error {
		var keys []string
		for _, user := range users {
			idK := r.DetailCache.KeyById(user, user.Id)
			unameK := r.DetailCache.Key(user,
				map[string][]interface{}{
					"username=?": {user.Username},
				})
			logx.WithContext(ctx).Infof("del user cache keys: %+v", []string{idK, unameK})
			keys = append(keys, idK, unameK)
		}
		err := r.Cache.Del(ctx, keys...).Err()
		if err != nil {
			return err
		}
		return nil
	}
}

```

## 获取用户信息
### logic
```go
package logic

func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserResp, error) {
	user := &model.User{}
	err := l.rep.DetailCache.FirstById(user, in.UserId)
	if err != nil {
		if xormerr.RecordNotFound(err) {
			err = nil
		}
		if err == global.RedisErrorNotExists {
			err = nil
		}
		if err != nil {
			l.Errorf("get user by id error: %s", err.Error())
			return &pb.GetUserResp{
				User: nil,
				BaseResp: &pb.BaseResp{
					ErrCode: -1,
					ErrMsg:  err.Error(),
				},
			}, err
		}
	}
	var userInfo *pb.User
	if user != nil {
		userInfo = &pb.User{
			Id:           user.Id,
			Username:     user.Username,
			Password:     user.Password,
			Nickname:     user.Nickname,
			Sign:         user.Sign,
			Avatar:       user.Avatar,
			Province:     user.Province,
			City:         user.City,
			District:     user.District,
			Birthday:     user.Birthday,
			RegisterTime: user.RegisterTime,
			IsMale:       user.IsMale,
		}
	}
	return &pb.GetUserResp{
		User: userInfo,
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}

func (l *GetUserByUsernameLogic) GetUserByUsername(in *pb.GetUserByUsernameReq) (*pb.GetUserResp, error) {
	user := &model.User{}
	user.Username = in.Username
	err := l.rep.DetailCache.FirstByWhere(
		user,
		map[string][]interface{}{
			"username=?": {in.Username},
		},
	)
	if err != nil {
		if xormerr.RecordNotFound(err) {
			err = nil
		}
		if err == global.RedisErrorNotExists {
			err = nil
		}
		if err != nil {
			l.Errorf("get user by username error: %s", err.Error())
			return &pb.GetUserResp{
				User: nil,
				BaseResp: &pb.BaseResp{
					ErrCode: -1,
					ErrMsg:  err.Error(),
				},
			}, err
		}
	}
	var userInfo *pb.User
	if user != nil {
		userInfo = &pb.User{
			Id:           user.Id,
			Username:     user.Username,
			Password:     user.Password,
			Nickname:     user.Nickname,
			Sign:         user.Sign,
			Avatar:       user.Avatar,
			Province:     user.Province,
			City:         user.City,
			District:     user.District,
			Birthday:     user.Birthday,
			RegisterTime: user.RegisterTime,
			IsMale:       user.IsMale,
		}
	}
	return &pb.GetUserResp{
		User: userInfo,
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}

```

## 修改用户信息
### logic
```go
package logic
func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	oldUser, err := NewGetUserByIdLogic(l.ctx, l.svcCtx).GetUserById(&pb.GetUserByIdReq{UserId: in.User.Id})
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Id:           in.User.Id,
		Username:     in.User.Username,
		Password:     in.User.Password,
		Nickname:     in.User.Nickname,
		Sign:         in.User.Sign,
		Avatar:       in.User.Avatar,
		Province:     in.User.Province,
		City:         in.User.City,
		District:     in.User.District,
		Birthday:     in.User.Birthday,
		RegisterTime: in.User.RegisterTime,
		IsMale:       in.User.IsMale,
	}
	updateMap := make(map[string]interface{})
	for _, field := range in.Fields {
		updateMap[field] = user.Value(field)
	}
	err = xorm.Transaction(l.rep.Mysql,
		// 插入数据
		func(tx *gorm.DB) error {
			return tx.Model(&model.User{}).Table(user.TableName()).
				Where("id = ?", user.Id).
				Updates(updateMap).Error
		},
		// 清除缓存
		l.rep.FuncDelUserCache(l.ctx, &model.User{
			Id:       oldUser.User.Id,
			Username: oldUser.User.Username,
		}, user),
	)
	if err != nil {
		return &pb.UpdateUserResp{
			BaseResp: &pb.BaseResp{
				ErrCode: -1,
				ErrMsg:  err.Error(),
			},
		}, err
	}
	// 预热数据
	go l.rep.WarmUpUser(l.ctx, user)
	return &pb.UpdateUserResp{
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}
```

## 登录
### logic
```go
package logic
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
```
### repository
```go
package repository

import (
	"context"
	"github.com/showurl/Zero-IM-Server/common/types"
	timeUtils "github.com/showurl/Zero-IM-Server/common/utils/time"
	"strconv"
	"time"
)

func (r *Rep) GetTokenMap(ctx context.Context, uid string, platform string) (map[string]int64, error) {
	key := types.RedisKeyToken + uid + ":" + platform
	result, err := r.Cache.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]int64{}, err
	}
	tokenMap := make(map[string]int64)
	for k, v := range result {
		tokenMap[k], _ = strconv.ParseInt(v, 10, 64)
	}
	return tokenMap, nil
}
func (r *Rep) SetTokenMap(ctx context.Context, uid string, platform string, token string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	expiredAt := timeUtils.Now().Add(time.Hour * 24 * time.Duration(r.svcCtx.Config.TokenRenewalDay))
	return r.Cache.HSet(ctx, key, token, expiredAt.UnixMilli()).Err()
}
func (r *Rep) DelTokenMap(ctx context.Context, uid string, platform string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	return r.Cache.Del(ctx, key).Err()
}

func (r *Rep) RenewalToken(ctx context.Context, uid string, platform string, token string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	expiredAt := timeUtils.Now().Add(time.Hour * 24 * time.Duration(r.svcCtx.Config.TokenRenewalDay))
	return r.Cache.HSet(ctx, key, token, expiredAt.UnixMilli()).Err()
}

func (r *Rep) DeleteToken(ctx context.Context, uid string, platform string, token string) error {
	key := types.RedisKeyToken + uid + ":" + platform
	return r.Cache.HDel(ctx, key, token).Err()
}

```
