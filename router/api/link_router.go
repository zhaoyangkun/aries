package api

import (
	"aries/handler/api"
	"github.com/gin-gonic/gin"
)

func InitLinkApiRouter(rootPath string, router *gin.Engine) {
	linkApiRouter := router.Group(rootPath)
	{
		linkApiRouter.GET("/all_links", api.GetAllLinks)
		linkApiRouter.GET("/links", api.GetLinksByPage)
		linkApiRouter.POST("/links", api.CreateLink)
		linkApiRouter.PUT("/links", api.UpdateLink)
		linkApiRouter.DELETE("/links/:id", api.DeleteLink)
		linkApiRouter.DELETE("/links", api.MultiDelLinks)
	}
}
