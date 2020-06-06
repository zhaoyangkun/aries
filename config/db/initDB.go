package db

import (
	"aries/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//数据库对象
var Db *gorm.DB

//数据库配置信息
var username = config.DBConfig.UserName
var password = config.DBConfig.Password
var host = config.DBConfig.Host
var database = config.DBConfig.Database
var port = config.DBConfig.Port
var timeZone = config.DBConfig.TimeZone
var maxIdleConn = config.DBConfig.MaxIdleConn
var maxOpenConn = config.DBConfig.MaxOpenConn

//获取数据库连接
func getDataSource(username string, password string, host string,
	port string, database string, timeZone string) string {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		username, password, host, port, database, timeZone)
	return dataSource
}

//初始化数据库连接
func init() {
	var err error
	dataSource := getDataSource(username, password, host, port, database, timeZone)
	Db, err = gorm.Open("mysql", dataSource) //连接数据库
	if err != nil {
		log.Panicln("数据库连接错误：", err.Error())
	}
	Db.DB().SetMaxIdleConns(maxIdleConn)
	Db.DB().SetMaxOpenConns(maxOpenConn)
	Db.LogMode(false) //是否开启日志
	/*	defer func() {
		err = Db.Close() //关闭数据库
		if err != nil {
			log.Panicln("数据库无法关闭：", err.Error())
		}
	}()*/
}
