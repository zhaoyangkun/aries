package util

// 状态码
const (
	Success      = 100 // 请求成功
	Redirect     = 101 // 重定向
	Forbidden    = 102 // 禁止访问
	RequestError = 103 // 请求数据缺失或者有误
	ServerError  = 104 // 服务器错误
)

// RestFulApi 返回信息结构
type Result struct {
	Code int         `json:"code" example:"000"` // 状态码
	Msg  string      `json:"msg" example:"信息"`   // 信息
	Data interface{} `json:"data"`               // 数据
}

// 令牌结构
type Token struct {
	Token string `json:"token"` // 令牌
	//RefreshToken string `json:"refresh_token"`
	Username string `json:"username"` // 用户名
	UserImg  string `json:"user_img"` // 头像
}

// 验证码结构
type Captcha struct {
	Id     string `json:"id"`
	ImgUrl string `json:"img_url"`
}
