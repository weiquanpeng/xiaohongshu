package system

import (
	"github.com/xiaohongshu/PnSql/server/dao/server"
)

type SysApi struct {
	Login
}

var (
	// 创建实例，保存帖子
	ToLogin = new(server.AuthLogin)
)
