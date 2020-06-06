package config

import (
	"aries/config/model"
	"github.com/gin-gonic/gin"
)

// 项目运行配置
var AppConfig = model.GinConfig{
	Mode: gin.DebugMode, //运行模式
	//Mode:gin.ReleaseMode,
	Port: "8088", //端口
}

// 数据库配置
var DBConfig = model.MysqlConfig{
	//Host:        "mysql",          //主机地址
	Host:        "127.0.0.1", //主机地址
	UserName:    "root",      //用户名
	Password:    "19960331",  //密码
	Database:    "aries",     //数据库名
	Port:        "3306",      //端口
	TimeZone:    "Local",     //时区
	MaxIdleConn: 20,          //最大空闲连接数
	MaxOpenConn: 100,         //最大打开连接数
}
