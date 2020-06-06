package app

import (
	"aries/config"
	"aries/router/api"
	"github.com/gin-gonic/gin"
)

// 加载配置
func InitApp() *gin.Engine {
	// 设置运行模式
	gin.SetMode(config.AppConfig.Mode)

	// 获取 engine
	router := gin.Default()

	// 加载静态资源
	//router.Static("/static", "./static")

	// 根据运行模式加载模板
	//if mode := gin.Mode(); mode == gin.TestMode {
	//	router.LoadHTMLGlob("../template/**/*")
	//} else {
	//	router.LoadHTMLGlob("template/**/*")
	//}

	// 路由分组
	api.InitCategoryApiRouter("/api/v1", router)

	return router
}
