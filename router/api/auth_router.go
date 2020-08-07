package api

import (
	"aries/handler/api"
	"github.com/gin-gonic/gin"
)

func InitAuthApiRouter(rootPath string, router *gin.Engine) {
	authApiRouter := router.Group(rootPath)
	{
		authApiRouter.POST("/auth/login", api.Login)
		authApiRouter.POST("/auth/register", api.Register)
		authApiRouter.GET("/auth/captcha", api.CreateCaptcha)
		authApiRouter.POST("/auth/pwd/forget", api.ForgetPwd)
		authApiRouter.POST("/auth/pwd/reset", api.ResetPwd)
	}
}
