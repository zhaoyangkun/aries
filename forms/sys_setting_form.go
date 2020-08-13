package forms

// 网站设置表单
type SiteSettingForm struct {
	SysId         string `json:"sys_id" label:"设置 ID"`
	TypeName      string `json:"type_name" binding:"required,max=50" label:"设置类型名称"`
	SiteName      string `json:"site_name" binding:"required,max=50" label:"网站名称"`
	SiteDesc      string `json:"site_desc" label:"网站描述"`
	SiteUrl       string `json:"site_url" binding:"required,max=255" label:"网站地址"`
	SiteLogo      string `json:"site_logo" label:"Logo"`
	SeoKeyWords   string `json:"seo_key_words" label:"SEO 关键词"`
	HeadContent   string `json:"head_content" label:"全局 head"`
	FooterContent string `json:"footer_content" label:"全局 footer"`
}

// SMTP 配置表单
type EmailSettingForm struct {
	SysId    string `json:"sys_id" label:"设置 ID"`
	TypeName string `json:"type_name" binding:"required,max=50" label:"设置类型名称"`
	Address  string `json:"address" binding:"required,max=30" label:"SMTP 地址"`
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

// 图床设置表单
type PicBedSettingForm struct {
	SysId       string `json:"sys_id"`
	StorageType string `json:"storage_type"`
}

// sm.ms 表单
type SmmsForm struct {
	SysId       string `json:"sys_id" label:"设置 ID"`
	StorageType string `json:"storage_type" binding:"required,max=20" label:"设置类型名称"`
	Token       string `json:"token" binding:"required,max=100"`
}

// 腾讯云 COS 表单
type TencentCosForm struct {
	SysId       string `json:"sys_id" label:"设置 ID"`
	StorageType string `json:"storage_type" binding:"required,max=20" label:"设置类型名称"`
	Host        string `json:"host" binding:"required,max=255" label:"存储桶地址"`
	Scheme      string `json:"scheme" binding:"required,max=5" label:"传输协议"`
	Region      string `json:"region" binding:"required,max=20" label:"区域"`
	SecretId    string `json:"secret_id" binding:"required,max=255" label:"secret_id"`
	SecretKey   string `json:"secret_key" binding:"required,max=255" label:"secret_key"`
	FolderPath  string `json:"folder_path" binding:"required,max=255" label:"上传目录"`
	ImgProcess  string `json:"img_process" binding:"max=255" label:"图片处理"`
}
