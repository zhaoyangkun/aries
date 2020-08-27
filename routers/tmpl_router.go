package routers

import (
	"aries/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type TmplRouter struct {
}

func (t *TmplRouter) InitTemplateRouter(rootPath string, router *gin.Engine) {
	tmplHandler := handlers.TmplHandler{}
	url := ginSwagger.URL("/swagger/doc.json")

	tmplRouter := router.Group(rootPath)
	{
		tmplRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		tmplRouter.GET("", tmplHandler.IndexTmpl)
		tmplRouter.GET("/page/:page", tmplHandler.IndexTmpl)
		tmplRouter.GET("/articles/:url", tmplHandler.ArticleTmpl)
		tmplRouter.GET("/error/403", tmplHandler.Error403Tmpl)
		tmplRouter.GET("/error/404", tmplHandler.Error404Tmpl)
		tmplRouter.GET("/error/500", tmplHandler.Error500Tmpl)
	}
}
