package forms

// SiteSettingForm 网站设置表单
type SiteSettingForm struct {
	SysId         string `json:"sys_id" label:"设置 ID"`
	TypeName      string `json:"type_name" binding:"required,max=50" label:"设置类型名称"`
	SiteName      string `json:"site_name" binding:"required,max=50" label:"网站名称"`
	SiteDesc      string `json:"site_desc" binding:"max=255" label:"网站描述"`
	SiteUrl       string `json:"site_url" binding:"required,max=255,url" label:"网站地址"`
	StaticRoot    string `json:"static_root" binding:"max=255,url" label:"静态路径"`
	SiteLogo      string `json:"site_logo" binding:"max=255" label:"Logo"`
	SeoKeyWords   string `json:"seo_key_words" binding:"max=255" label:"SEO 关键词"`
	RecordNumber  string `json:"record_number" binding:"max=255" label:"备案号"`
	HeadContent   string `json:"head_content" binding:"max=1000" label:"全局 head"`
	FooterContent string `json:"footer_content" binding:"max=1000" label:"全局 footer"`
}

// EmailSettingForm SMTP 配置表单
type EmailSettingForm struct {
	SysId    string `json:"sys_id" label:"设置 ID"`
	TypeName string `json:"type_name" binding:"required,max=50" label:"设置类型名称"`
	Address  string `json:"address" binding:"required,max=30" label:"SMTP 地址"`
	Port     string `json:"port" binding:"required,max=3" label:"端口"`
	Account  string `json:"account" binding:"required,max=30,email" label:"邮箱帐号"`
	Pwd      string `json:"pwd" binding:"required,max=30" label:"密码"`
	Sender   string `json:"sender" binding:"required,max=30" label:"发送人"`
}

// EmailSendForm 发送邮件测试表单
type EmailSendForm struct {
	Sender       string `json:"sender" binding:"required,max=30" label:"发送人"`
	ReceiveEmail string `json:"receive_email" binding:"required,max=30,email" label:"接收邮箱"`
	Title        string `json:"title" binding:"required,max=100" label:"邮箱标题"`
	Content      string `json:"content" binding:"required,max=1000" label:"邮件内容"`
}

// PicBedSettingForm 图床设置表单
type PicBedSettingForm struct {
	SysId       string `json:"sys_id"`
	StorageType string `json:"storage_type"`
}

// QubuForm 去不图床表单
type QubuForm struct {
	SysId       string `json:"sys_id" label:"设置 ID"`
	StorageType string `json:"storage_type" binding:"required,max=20" label:"设置类型名称"`
	Token       string `json:"token" binding:"required,max=100"`
}

// SmmsForm sm.ms 表单
type SmmsForm struct {
	SysId       string `json:"sys_id" label:"设置 ID"`
	StorageType string `json:"storage_type" binding:"required,max=20" label:"设置类型名称"`
	Token       string `json:"token" binding:"required,max=100"`
}

// ImgbbForm imgbb 表单
type ImgbbForm struct {
	SysId       string `json:"sys_id" label:"设置 ID"`
	StorageType string `json:"storage_type" binding:"required,max=20" label:"设置类型名称"`
	Token       string `json:"token" binding:"required,max=100"`
}

// TencentCosForm 腾讯云 COS 表单
type TencentCosForm struct {
	SysId       string `json:"sys_id" label:"设置 ID"`
	StorageType string `json:"storage_type" binding:"required,max=20" label:"设置类型名称"`
	Host        string `json:"host" binding:"required,max=255" label:"存储桶域名"`
	Scheme      string `json:"scheme" binding:"required,max=5" label:"传输协议"`
	Region      string `json:"region" binding:"required,max=20" label:"区域"`
	SecretId    string `json:"secret_id" binding:"required,max=255" label:"secret_id"`
	SecretKey   string `json:"secret_key" binding:"required,max=255" label:"secret_key"`
	FolderPath  string `json:"folder_path" binding:"required,max=255" label:"上传目录"`
	ImgProcess  string `json:"img_process" binding:"max=255" label:"图片处理"`
}

// CommentPlugInForm 评论组件表单
type CommentPlugInForm struct {
	SysId  string `json:"sys_id" label:"设置 ID"`
	PlugIn string `json:"plug_in" label:"评论组件"`
}

// LocalCommentSettingForm 本地评论设置表单
type LocalCommentSettingForm struct {
	SysId      string `json:"sys_id" label:"设置 ID"`
	PlugIn     string `json:"plug_in" label:"评论组件"`
	IsOn       string `json:"is_on" binding:"required" label:"是否开启评论"`
	IsReviewOn string `json:"is_review_on" binding:"required" label:"是否开启评论审核"`
	IsReplyOn  string `json:"is_reply_on" label:"是否开启邮箱回复"`
	PageSize   string `json:"page_size" label:"每页评论条数"`
}

// TwikooSettingForm twikoo 评论组件设置表单
type TwikooSettingForm struct {
	SysId  string `json:"sys_id" label:"设置 ID"`
	PlugIn string `json:"plug_in" label:"评论组件"`
	EnvId  string `json:"env_id" binding:"required" label:"环境 ID"`
	Region string `json:"region" label:"区域"`
	Path   string `json:"path" label:"文章 URL 路径"`
	Lang   string `json:"lang" label:"语言"`
}

// ParamSettingForm 参数设置表单
type ParamSettingForm struct {
	SysId           string `json:"sys_id" label:"设置 ID"`
	TypeName        string `json:"type_name" label:"设置类型名称"`
	IndexPageSize   string `json:"index_page_size" label:"首页每页条数"`
	ArchivePageSize string `json:"archive_page_size" label:"归档页每页条数"`
	SiteMapPageSize string `json:"site_map_page_size" label:"站点地图每页条数"`
}

// SocialInfoForm 社交信息
type SocialInfoForm struct {
	SysId    string `json:"sys_id" label:"设置 ID"`
	TypeName string `json:"type_name" label:"设置类型名称"`
	QQ       string `json:"qq" label:"qq"`
	Wechat   string `json:"wechat" label:"微信"`
	Github   string `json:"github" label:"github"`
	Weibo    string `json:"weibo" label:"微博"`
	Zhihu    string `json:"zhihu" label:"知乎"`
}
