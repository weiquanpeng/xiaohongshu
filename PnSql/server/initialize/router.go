package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/global"
	"github.com/xiaohongshu/PnSql/server/router"
)

func InitRouter() {
	r := gin.New()
	r.Use(gin.Recovery())

	systemRouter := router.AllRouter.Sys_Group

	PublicGroup := r.Group("/api/v1/public")
	{
		systemRouter.InitLoginRouter(PublicGroup)
		systemRouter.InitCodeRouter(PublicGroup)
	}
	//PrivateGroup := r.Group("/api/v1/")
	//{
	//
	//}

	stPort := global.P_cfg.System.Port
	if stPort == "" {
		stPort = "2379"
	}
	err := r.Run(fmt.Sprintf(":%s", stPort))
	if err != nil {
		panic(fmt.Sprintf("start server error: %s", err.Error()))
	}
}
