package routers

import (
	"aries/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type TmplRouter struct {
}

func (t *TmplRouter) InitTemplateRouter(rootPath string, router *gin.Engine) {
	frontHandler := handlers.FrontHandler{}
	frontRouter := router.Group(rootPath)
	{
		url := ginSwagger.URL("/swagger/doc.json")
		frontRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		frontRouter.GET("", frontHandler.IndexHTML)
	}
}
