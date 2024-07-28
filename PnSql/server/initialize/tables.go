package initialize

import (
	"fmt"
	"github.com/xiaohongshu/PnSql/server/dao/model"
	"github.com/xiaohongshu/PnSql/server/global"
	"os"
)

func RegisterTables() {
	db := global.PVA_DB
	err := db.AutoMigrate(
		// 系统模块表
		model.Users{},
	)
	if err != nil {
		fmt.Println("register table err:", err)
		os.Exit(0)
	}
}
