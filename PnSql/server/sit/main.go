package main

import (
	"github.com/xiaohongshu/PnSql/server/core"
	"github.com/xiaohongshu/PnSql/server/global"
)

func main() {
	core.Viper()
	core.GormMysql()
	global.PVA_DB = core.GormMysql()

}
