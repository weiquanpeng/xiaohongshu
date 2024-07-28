package initialize

import (
	"github.com/xiaohongshu/PnSql/server/core"
	"github.com/xiaohongshu/PnSql/server/global"
)

func InitDB() {
	global.PVA_DB = core.GormMysql()
	RegisterTables()
}
