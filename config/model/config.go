package model

// 项目配置结构体
type GinConfig struct {
	Mode            string   // 运行模式
	Port            string   // 运行端口
	TokenExpireTime int      // JWT token 过期时间
	AllowedRefers   []string // 允许的 referer
}

// 数据库配置结构体
type MysqlConfig struct {
	Host        string // 主机地址
	UserName    string // 用户名
	Password    string // 密码
	Database    string // 数据库名
	Port        string // 端口
	TimeZone    string // 时区
	MaxIdleConn int    // 最大空闲连接数
	MaxOpenConn int    // 最大打开连接数
}
