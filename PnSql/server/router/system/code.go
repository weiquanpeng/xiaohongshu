package system

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/api/v1/code"
)

type CodeRouter struct{}

func (c *CodeRouter) InitCodeRouter(Router *gin.RouterGroup) {
	codeApi := code.CodeApi{}
	course := Router.Group("code")
	{
		course.GET("get", codeApi.CreateCaptcha)
	}
}
