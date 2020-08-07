package router

import (
	"aries/handler"
	"github.com/gin-gonic/gin"
)

func InitFrontRouter(rootPath string, router *gin.Engine) {
	frontRouter := router.Group(rootPath)
	{
		frontRouter.GET("", handler.IndexHTML)
	}
}
