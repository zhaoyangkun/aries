package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwaggerRouter(rootPath string, router *gin.Engine) {
	url := ginSwagger.URL("/swagger/doc.json")
	swaggerRouter := router.Group(rootPath)
	{
		swaggerRouter.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
