package main

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

func main() {
	type CodeApi struct{}

	type configJsonBody struct {
		Id            string
		CaptchaType   string
		VerifyValue   string
		DriverAudio   *base64Captcha.DriverAudio
		DriverString  *base64Captcha.DriverString
		DriverChinese *base64Captcha.DriverChinese
		DriverMath    *base64Captcha.DriverMath
		DriverDigit   *base64Captcha.DriverDigit
	}

	// 定义验证存储
	var store = base64Captcha.DefaultMemStore
	var digit = base64Captcha.DefaultDriverDigit
	captcha := base64Captcha.NewCaptcha(digit, store)
	id, baseURL, _, err := captcha.Generate()
	if err != nil {
		fmt.Println(err)
	}
	//captcha := base64Captcha.NewCaptcha(digit, store)
	fmt.Println("-------------")
	fmt.Println(id, baseURL)
	fmt.Println(digit.Length)

}
