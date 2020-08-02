package api

import (
	"aries/handler/api"
	"github.com/gin-gonic/gin"
)

func InitSysSettingApiRouter(rootPath string, router *gin.Engine) {
	sysSettingApiRouter := router.Group(rootPath)
	{
		sysSettingApiRouter.GET("/sys_setting_items", api.GetSysSettingItem)
		sysSettingApiRouter.POST("/site_setting", api.SaveSiteSetting)
		sysSettingApiRouter.POST("/smtp_setting", api.SaveSMTPSetting)
		sysSettingApiRouter.POST("/test_send_email", api.SendTestEmail)
	}
}
