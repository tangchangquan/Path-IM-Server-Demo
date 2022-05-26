package xconf

import (
	"github.com/showurl/Path-IM-Server/common/xenv"
	"github.com/zeromicro/go-zero/core/conf"
	"os"
)

func MustLoad(path string, v interface{}) {
	// 判断文件是否存在
	_, e := os.ReadFile(path)
	if e != nil {
		// 读系统环境变量 解析到 value
		err := xenv.Parse(v)
		if err != nil {
			panic(err)
		}
	} else {
		conf.MustLoad(path, v)
	}
}
