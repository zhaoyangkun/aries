package api

import (
	"aries/handler/api"
	"aries/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysSettingApiRouter(rootPath string, router *gin.Engine) {
	sysSettingApiRouter := router.Group(rootPath, middleware.JWTAuthMiddleWare())
	{
		sysSettingApiRouter.GET("/sys_setting/items", api.GetSysSettingItem)
		sysSettingApiRouter.POST("/sys_setting/site", api.SaveSiteSetting)
		sysSettingApiRouter.POST("/sys_setting/smtp", api.SaveSMTPSetting)
		sysSettingApiRouter.POST("/sys_setting/pic_bed", api.SavePicBedSetting)
		sysSettingApiRouter.POST("/sys_setting/email/test", api.SendTestEmail)
		sysSettingApiRouter.GET("/sys_setting/index_info", api.GetAdminIndexData)
	}
}
