package api

import (
	"aries/handler/api"
	"aries/middleware"
	"github.com/gin-gonic/gin"
)

func InitArticleApiRouter(rootPath string, router *gin.Engine) {
	ArticleApiRouter := router.Group(rootPath)
	ArticleApiRouter.Use(middleware.JWTAuthMiddleWare()) // 加载 JWT 权限校验中间件
	{
		ArticleApiRouter.GET("/all_articles", api.GetAllArticles)
		ArticleApiRouter.GET("/articles/:id", api.GetArticleById)
		ArticleApiRouter.GET("/articles", api.GetArticlesByPage)
		ArticleApiRouter.POST("/articles", api.AddArticle)
		ArticleApiRouter.PUT("/articles", api.UpdateArticle)
		ArticleApiRouter.DELETE("/articles/:id", api.DeleteArticle)
		ArticleApiRouter.DELETE("/articles", api.MultiDelArticles)
	}
}
