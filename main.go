package main

import (
	"aries/config/app"
	"aries/config/setting"
	"log"
)

func main() {
	engine := app.InitApp()                             // 加载相关配置
	err := engine.Run(":" + setting.Config.Server.Port) // 监听端口
	if err != nil {
		log.Panicln("err: ", err.Error())
	}
}
