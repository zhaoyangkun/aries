package app

import (
	"aries/config/db"
	"aries/config/migrate"
	"aries/config/setting"
	templRouter "aries/router"
	apiRouter "aries/router/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// 初始化 gin
func InitApp() *gin.Engine {
	// 加载配置文件
	setting.InitSetting()
	// 连接数据库
	db.InitDb()
	// 反向生成数据表
	migrate.Migrate()

	// 设置运行模式
	gin.SetMode(setting.Config.Server.Mode)
	// 开启日志颜色
	gin.ForceConsoleColor()
	// 获取 engine
	router := gin.Default()

	// 表单翻译参数
	uni := ut.New(zh.New())
	setting.Trans, _ = uni.GetTranslator("zh")
	// 表单校验配置
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册翻译器
		_ = translations.RegisterDefaultTranslations(v, setting.Trans)
		// 注册一个函数，获取 struct tag 里自定义的 label 作为字段名
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("label")
			return name
		})
	}

	// 加载静态资源
	//router.Static("/static", "./static")

	// 根据运行模式加载模板
	//if mode := gin.Mode(); mode == gin.TestMode {
	//	router.LoadHTMLGlob("../template/**/*")
	//} else {
	//	router.LoadHTMLGlob("template/**/*")
	//}

	// 模板路由
	templRouter.InitSwaggerRouter("/swagger", router)
	// api路由
	apiRouter.InitCategoryApiRouter("/api/v1", router)
	apiRouter.InitAuthApiRouter("/api/v1", router)
	apiRouter.InitTagApiRouter("/api/v1", router)

	return router
}
