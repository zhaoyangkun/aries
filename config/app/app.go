package app

import (
	"aries/config/db"
	"aries/config/migrate"
	"aries/config/setting"
	"aries/handlers"
	logger "aries/log"
	"aries/middlewares"
	"aries/models"
	"aries/routers"
	"aries/utils"
	"fmt"
	"html/template"
	"log"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
)

// InitApp 初始化 gin
func InitApp() *gin.Engine {
	// 加载配置
	s := setting.Setting{}
	s.InitSetting()
	s.InitLute()
	s.InitCache()
	db.InitDb()
	migrate.Migrate()
	gin.SetMode(setting.Config.Server.Mode)

	// 加载中间件
	router := gin.New()
	err := logger.InitLogger(
		setting.Config.Logger.FileName,
		setting.Config.Logger.Level,
		setting.Config.Logger.Format,
		setting.Config.Logger.MaxSize,
		setting.Config.Logger.MaxBackups,
		setting.Config.Logger.MaxAge,
	)
	if err != nil {
		log.Panicln("初始化日志失败：", err.Error())
	}
	middlewares.InitBucket(time.Second*time.Duration(setting.Config.Server.LimitTime), setting.Config.Server.LimitCap)
	router.Use(middlewares.Logger(logger.Logger), middlewares.Recover(logger.Logger, true), middlewares.Limiter())

	// 配置表单校验
	uni := ut.New(zh.New())
	setting.Trans, _ = uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = translations.RegisterDefaultTranslations(v, setting.Trans)
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("label")
			return name
		})
	}

	// 配置博客全局变量
	blogSetting, _ := models.SysSettingItem{}.GetBySysSettingName("网站设置")
	socialInfo, _ := models.SysSettingItem{}.GetBySysSettingName("社交信息")
	setting.BlogVars.InitBlogVars(blogSetting, socialInfo)

	// 初始化模板全局变量
	handlers.InitTmplVars()

	// 加载自定义模板函数
	router.SetFuncMap(template.FuncMap{
		"safe":     utils.SafeHtml,
		"add":      utils.AddUpTwoNum,
		"subtract": utils.SubtractTwoNum,
		"mod":      utils.Mod,
		"year":     utils.Year,
		"month":    utils.Month,
		"day":      utils.Day,
	})

	// 加载静态资源和模板
	router.Static("/static", fmt.Sprintf("./resources/themes/%s/static", setting.BlogVars.Theme))
	router.Static("/admin", "./resources/dist")
	router.LoadHTMLGlob(fmt.Sprintf("./resources/themes/%s/templates/**/*", setting.BlogVars.Theme))

	// 加载路由
	apiRouter := routers.ApiRouter{}
	tmplRouter := routers.TmplRouter{}
	tmplRouter.InitTemplateRouter("", router)
	apiRouter.InitApiRouter("/api/v1", router)

	return router
}
