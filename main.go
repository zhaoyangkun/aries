package main

import (
	"aries/config"
	"aries/config/app"
	"log"
)

func main() {
	engine := app.InitApp()                        // 加载相关配置
	err := engine.Run(":" + config.AppConfig.Port) // 监听端口
	if err != nil {
		log.Panicln("err: ", err.Error())
	}
}
