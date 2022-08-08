package db

import (
	"aries/config/setting"
	"fmt"
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // mysql驱动
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // gorm mysql
)

// Db 数据库对象
var Db *gorm.DB

// 获取数据库连接
func getDataSource() string {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		setting.Config.Database.UserName,
		setting.Config.Database.Password,
		setting.Config.Database.Host,
		setting.Config.Database.Port,
		setting.Config.Database.Database,
		url.QueryEscape(setting.Config.Database.TimeZone), // 对时区进行 Url 编码
	)
	return dataSource
}

// InitDb 初始化数据库连接
func InitDb() {
	var err error

	//连接数据库
	Db, err = gorm.Open("mysql", getDataSource())
	if err != nil {
		log.Panic("数据库连接错误：", err.Error())
	}

	// 设置连接池参数
	Db.DB().SetMaxIdleConns(setting.Config.Database.MaxIdleConn)
	Db.DB().SetMaxOpenConns(setting.Config.Database.MaxOpenConn)

	// 开发环境下开启 sql 日志
	if setting.Config.Server.Mode == gin.DebugMode {
		Db.LogMode(true)
	}
}
