package api

import (
	"aries/handler/api"
	"github.com/gin-gonic/gin"
)

func InitCategoryApiRouter(rootPath string, router *gin.Engine) {
	categoryApiRouter := router.Group(rootPath)
	{
		categoryApiRouter.GET("/categories", api.GetAllCategoriesHandler)
	}
}
