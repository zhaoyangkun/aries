package config

import (
	"aries/config/model"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

// 项目运行配置
var AppConfig = model.GinConfig{
	//Mode:gin.ReleaseMode,
	Mode:            gin.DebugMode,                      // 运行模式
	Port:            "8088",                             // 端口
	TokenExpireTime: 3600,                               // JWT token 过期时间
	AllowedRefers:   []string{"localhost", "127.0.0.1"}, // 允许的 referer
}

// 数据库配置
var DBConfig = model.MysqlConfig{
	//Host:        "mysql",          // 主机地址
	Host:        "127.0.0.1", // 主机地址
	UserName:    "root",      // 用户名
	Password:    "19960331",  // 密码
	Database:    "aries",     // 数据库名
	Port:        "3306",      // 端口
	TimeZone:    "Local",     // 时区
	MaxIdleConn: 20,          // 最大空闲连接数
	MaxOpenConn: 100,         // 最大打开连接数
}

// 翻译器
var Trans ut.Translator
