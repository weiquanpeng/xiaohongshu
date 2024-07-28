package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/xiaohongshu/PnSql/server/common/response"
	"github.com/xiaohongshu/PnSql/server/dao/model"
)

type Login struct{}

var store = base64Captcha.DefaultMemStore

// 生成验证码
func (l *Login) Login(c *gin.Context) {
	digit := &base64Captcha.DriverDigit{Height: 80, Width: 240, Length: 6, MaxSkew: 0.8, DotCount: 120}
	captcha := base64Captcha.NewCaptcha(digit, store)
	id, baseURL, _, err := captcha.Generate()
	if err != nil {
		response.FailWithMessage("验证生成错误", c)
		return
	}
	response.OkWithData(map[string]any{"id": id, "baseURL": baseURL}, c)
}

// 登录接口
func (l *Login) ToLogin(c *gin.Context) {
	type LoginParam struct {
		Username string `form:"username"`
		Password string `form:"password"`
		Id       string `form:"id"`
		Code     string `form:"code"`
	}
	param := LoginParam{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithDetailed(err.Error(), "参数格式错误", c)
		return
	}

	//校验验证码
	verify := store.Verify(param.Id, param.Code, true)
	if !verify {
		response.FailWithDecide(602, "验证码错误", c)
		return
	}

	//校验用户、密码
	var judge bool
	judge = ToLogin.GetUsers(param.Username, param.Password)
	if !judge {
		response.FailWithDecide(603, "用户或密码错误", c)
		return
	}
	response.OkWithData("登录成功", c)

}

// 注册用户
func (l *Login) RegisterUser(c *gin.Context) {
	var user model.Users
	err := c.ShouldBindJSON(&user)
	fmt.Println(user)
	if err != nil {
		response.FailWithDetailed(err.Error(), "参数格式错误", c)
		return
	}
	err = ToLogin.CreateUsers(&user)
	if err != nil {
		response.FailWithMessage("注册失败", c)
		return
	}
	response.OkWithMessage("注册成功", c)
}
