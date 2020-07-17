package api

import (
	"aries/handler/api"
	"aries/middleware"
	"github.com/gin-gonic/gin"
)

func InitTagApiRouter(rootPath string, router *gin.Engine) {
	tagApiRouter := router.Group(rootPath, middleware.JWTAuthMiddleWare())
	{
		tagApiRouter.GET("/all_tags", api.GetAllTags)
		tagApiRouter.GET("/tags", api.GetTagsByPage)
		tagApiRouter.GET("/tags/:id", api.GetTagById)
		tagApiRouter.POST("/tags", api.AddTag)
		tagApiRouter.PUT("/tags", api.UpdateTag)
		tagApiRouter.DELETE("/tags/:id", api.DeleteTag)
		tagApiRouter.DELETE("/tags", api.MultiDelTags)
	}
}
