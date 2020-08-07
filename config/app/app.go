package app

import (
	"aries/config/db"
	"aries/config/migrate"
	"aries/config/setting"
	"aries/middleware"
	"aries/model"
	routers "aries/router"
	"fmt"
	"github.com/88250/lute"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"time"
)

// 初始化 gin
func InitApp() *gin.Engine {
	// 加载配置
	setting.InitSetting()
	db.InitDb()
	migrate.Migrate()
	setting.LuteEngine = lute.New()
	setting.Cache = persistence.NewInMemoryStore(time.Hour * 1)
	gin.SetMode(setting.Config.Server.Mode)

	// 加载中间件
	router := gin.New()
	router.Use(middleware.LoggerMiddleWare(), gin.Recovery())

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
	blogSetting, _ := model.SysSettingItem{}.GetBySysSettingName("网站设置")
	setting.BlogVars.InitBlogVars(blogSetting)

	// 加载静态资源和模板
	router.Static("/static", fmt.Sprintf("theme/%s/static", setting.BlogVars.Theme))
	router.LoadHTMLGlob(fmt.Sprintf("theme/%s/template/**", setting.BlogVars.Theme))

	// 加载路由
	apiRouter := routers.ApiRouter{}
	tmplRouter := routers.TmplRouter{}
	tmplRouter.InitTemplateRouter("", router)
	apiRouter.InitApiRouter("/api/v1", router)

	return router
}
