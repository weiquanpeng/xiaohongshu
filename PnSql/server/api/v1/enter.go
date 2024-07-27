package v1

import "github.com/xiaohongshu/PnSql/server/api/v1/system"

type ApiGroup struct {
	SystemApiGroup system.SysApi
}

var ApiGroupApp = new(ApiGroup)
