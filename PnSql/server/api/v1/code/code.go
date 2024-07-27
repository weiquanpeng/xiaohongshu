package code

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/xiaohongshu/PnSql/server/router/common/response"
)

type CodeApi struct{}

type codeJsonBody struct {
	Id   string `form:"id"`
	Code string `form:"code"`
}

// 创建验证码
func (codeApi *CodeApi) CreateCaptcha(c *gin.Context) {
	// 定义验证存储
	var store = base64Captcha.DefaultMemStore
	var digit = base64Captcha.DefaultDriverDigit
	digit.Length = 6
	captcha := base64Captcha.NewCaptcha(digit, store)
	id, baseURL, _, err := captcha.Generate()
	if err != nil {
		response.FailWithMessage("验证生成错误", c)
		return
	}
	response.OkWithData(map[string]any{"id": id, "baseURL": baseURL}, c)
}

func (codeApi *CodeApi) VerifyCaptcha(c *gin.Context) {
	newcode := codeJsonBody{}
	err := c.ShouldBind(&newcode)
	if err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}

}
