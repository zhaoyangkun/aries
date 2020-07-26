package main

import (
	"aries/config/app"
	"aries/config/setting"
	_ "aries/docs"
	log "github.com/sirupsen/logrus"
)

// @title Gin Swagger
// @version 1.0
// @description Aries 开源博客项目 API 接口文档

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8088
func main() {
	engine := app.InitApp()                             // 初始化
	err := engine.Run(":" + setting.Config.Server.Port) // 监听端口并运行
	if err != nil {
		log.Panicln("项目启动失败: ", err.Error())
	}
}
