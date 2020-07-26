package api

import (
	"aries/handler/api"
	"github.com/gin-gonic/gin"
)

func InitUserApiRouter(rootPath string, router *gin.Engine) {
	userApiRouter := router.Group(rootPath)
	{
		userApiRouter.GET("/all_users", api.GetAllUsers)
	}
}
