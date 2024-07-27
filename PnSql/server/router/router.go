package router

import "github.com/xiaohongshu/PnSql/server/router/system"

type EnterGroup struct {
	Sys_Group system.SysGroup
}

var AllRouter = new(EnterGroup)
