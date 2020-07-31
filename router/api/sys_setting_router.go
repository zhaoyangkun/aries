package api

import (
	"aries/handler/api"
	"github.com/gin-gonic/gin"
)

func InitSysSettingApiRouter(rootPath string, router *gin.Engine) {
	sysSettingApiRouter := router.Group(rootPath)
	{
		sysSettingApiRouter.GET("/sys_setting_items", api.GetSysSettingItem)
		sysSettingApiRouter.POST("/smtp", api.SaveSMTP)
		sysSettingApiRouter.POST("/test_email", api.SendTestEmail)
	}
}
