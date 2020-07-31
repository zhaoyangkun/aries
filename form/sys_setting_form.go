package form

// SMTP 配置表单
type EmailForm struct {
	SysId    string `json:"sys_id" label:"设置 ID"`
	TypeName string `json:"type_name" binding:"required,max=50" label:"设置类型名称"`
	Address  string `json:"address" binding:"required,max=30" label:"SMTP 地址"`
	Protocol string `json:"protocol" binding:"required,max=5" label:"协议"`
	Port     string `json:"port" binding:"required,max=3" label:"端口"`
	Account  string `json:"account" binding:"required,max=30,email" label:"邮箱帐号"`
	Pwd      string `json:"pwd" binding:"required,max=30" label:"密码"`
	Sender   string `json:"sender" binding:"required,max=30" label:"发送人"`
}

// 发送邮件测试表单
type EmailSendForm struct {
	Sender       string `json:"sender" binding:"required,max=30" label:"发送人"`
	ReceiveEmail string `json:"receive_email" binding:"required,max=30,email" label:"接收邮箱"`
	Title        string `json:"title" binding:"required,max=100" label:"邮箱标题"`
	Content      string `json:"content" binding:"required,max=1200" label:"邮件内容"`
}
