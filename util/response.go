package util

// RestFulApi 返回信息结构体
type Result struct {
	Code int         `json:"code" example:"000"` // 状态码，参考 http　状态码
	Msg  string      `json:"msg" example:"信息"`   // 信息
	Data interface{} `json:"data"`               // 数据
}

// 令牌结构体
type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Username     string `json:"username"`
	UserImg      string `json:"user_img"`
}

// 验证码结构体
type Captcha struct {
	Id     string `json:"id"`
	ImgUrl string `json:"img_url"`
}
