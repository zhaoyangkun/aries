package api

import (
	"aries/handler/api"
	"aries/middleware"
	"github.com/gin-gonic/gin"
)

func InitCommentApiRouter(rootPath string, router *gin.Engine) {
	commentRouter := router.Group(rootPath, middleware.JWTAuthMiddleWare())
	{
		commentRouter.GET("/all_comments", api.GetAllComments)
		commentRouter.GET("/comments", api.GetCommentsByPage)
		commentRouter.POST("/comments", api.AddComment)
		commentRouter.PUT("/comments", api.UpdateComment)
		commentRouter.DELETE("/comments/:id", api.DeleteComment)
		commentRouter.DELETE("/comments", api.MultiDelComments)
	}

}
