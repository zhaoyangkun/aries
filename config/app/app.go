package app

import (
	"aries/config/setting"
	normalRouter "aries/router"
	apiRouter "aries/router/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// 加载配置
func InitApp() *gin.Engine {
	// 设置运行模式
	gin.SetMode(setting.Config.Server.Mode)

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
		//根据提供的标记注册翻译
		//_ = v.RegisterTranslation("bookabledate", trans, func(ut ut.Translator) error {
		//	return ut.Add("bookabledate", "{0}不能早于当前时间或{1}格式错误!", true)
		//}, func(ut ut.Translator, fe validator.FieldError) string {
		//	t, _ := ut.T("bookabledate", fe.Field(), fe.Field())
		//	return t
		//})
	}

	// 加载静态资源
	//router.Static("/static", "./static")

	// 根据运行模式加载模板
	//if mode := gin.Mode(); mode == gin.TestMode {
	//	router.LoadHTMLGlob("../template/**/*")
	//} else {
	//	router.LoadHTMLGlob("template/**/*")
	//}

	// 路由分组
	normalRouter.InitSwaggerRouter("/swagger", router)
	apiRouter.InitCategoryApiRouter("/api/v1", router)
	apiRouter.InitAuthApiRouter("/api/v1", router)

	return router
}
