package utils

import (
	"github.com/mojocn/base64Captcha"
)

// CaptchaConfig 验证码结构体
type CaptchaConfig struct {
	Id            string                       `json:"id"`
	CaptchaType   string                       `json:"captcha_type"`
	VerifyValue   string                       `json:"verify_value"`
	DriverAudio   *base64Captcha.DriverAudio   `json:"driver_audio"`
	DriverString  *base64Captcha.DriverString  `json:"driver_string"`
	DriverChinese *base64Captcha.DriverChinese `json:"driver_chinese"`
	DriverMath    *base64Captcha.DriverMath    `json:"driver_math"`
	DriverDigit   *base64Captcha.DriverDigit   `json:"driver_digit"`
}

var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 生成验证码
func GenerateCaptcha(captcha *CaptchaConfig) (string, error) {
	var driver base64Captcha.Driver

	// 根据验证码类型生成 base64 验证码
	switch captcha.CaptchaType {
	case "audio":
		driver = captcha.DriverAudio
	case "string":
		driver = captcha.DriverString.ConvertFonts()
	case "math":
		driver = captcha.DriverMath.ConvertFonts()
	case "chinese":
		driver = captcha.DriverChinese.ConvertFonts()
	default:
		captcha.DriverDigit = base64Captcha.NewDriverDigit(38, 120, 4, 0.6, 50)
		driver = captcha.DriverDigit
	}
	//初始化driver
	//captcha.DriverDigit = base64Captcha.NewDriverDigit(38, 120, 4, 0.7, 80)
	/*	captcha.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
		driver = captcha.DriverDigit*/
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	captcha.Id = id

	if err != nil {
		return "", err
	}

	return b64s, nil
}

// CaptchaVerify 校验验证码
func CaptchaVerify(captcha *CaptchaConfig) bool {
	return store.Verify(captcha.Id, captcha.VerifyValue, false)
}
