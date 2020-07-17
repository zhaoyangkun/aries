package api

import (
	"aries/handler/api"
	"aries/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserApiRouter(rootPath string, router *gin.Engine) {
	userApiRouter := router.Group(rootPath, middleware.JWTAuthMiddleWare())
	{
		userApiRouter.GET("/all_users", api.GetAllUsers)
	}
}
