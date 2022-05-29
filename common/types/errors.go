package types

import (
	"fmt"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xerr"
)

var WSDataError = xerr.New(3001, "ws data error")
var (
	ErrTokenExpired     = fmt.Errorf("token 过期")
	ErrTokenInvalid     = fmt.Errorf("token 无效")
	ErrTokenMalformed   = fmt.Errorf("token 格式错误")
	ErrTokenNotValidYet = fmt.Errorf("token 还未生效")
	ErrTokenUnknown     = fmt.Errorf("未知错误")
	ErrTokenKicked      = fmt.Errorf("被踢出")
)

// error code
const (
	ErrCodeOK     = iota // 成功
	ErrCodeFailed        // 失败
	ErrCodeLimit         // 限流

	ErrCodeProtoUnmarshal = 400 // proto解析错误
	ErrCodeParams         = 401 // 参数错误
)
