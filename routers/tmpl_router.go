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
		tmplRouter.GET("/p/:page", tmplHandler.IndexTmpl)
		tmplRouter.GET("/articles/:url", tmplHandler.ArticleTmpl)
		tmplRouter.GET("/categories", tmplHandler.CategoryTmpl)
		tmplRouter.GET("/category/p/:page", tmplHandler.CategoryTmpl)
		tmplRouter.GET("/categories/:url", tmplHandler.CategoryTmpl)
		tmplRouter.GET("/categories/:url/p/:page", tmplHandler.CategoryTmpl)
		tmplRouter.GET("/tags", tmplHandler.TagTmpl)
		tmplRouter.GET("/tag/p/:page", tmplHandler.TagTmpl)
		tmplRouter.GET("/tags/:name", tmplHandler.TagTmpl)
		tmplRouter.GET("/tags/:name/p/:page", tmplHandler.TagTmpl)
		tmplRouter.GET("/archives", tmplHandler.ArchiveTmpl)
		tmplRouter.GET("/links", tmplHandler.LinkTmpl)
		tmplRouter.GET("/journals", tmplHandler.JournalTmpl)
		tmplRouter.GET("/galleries", tmplHandler.GalleryTmpl)
		tmplRouter.GET("/custom/:url", tmplHandler.CustomTmpl)
		tmplRouter.GET("/search", tmplHandler.SearchTmpl)
		tmplRouter.GET("/search/p/:page", tmplHandler.SearchTmpl)
		tmplRouter.GET("/sitemap.xml", tmplHandler.SiteMapXml)
	}
}
