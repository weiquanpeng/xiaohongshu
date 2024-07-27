package main

import (
	"github.com/xiaohongshu/PnSql/server/core"
	"github.com/xiaohongshu/PnSql/server/initialize"
)

func main() {
	core.Viper()
	initialize.InitRouter()
}
