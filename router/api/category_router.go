package api

import (
	"aries/handler/api"
	"aries/middleware"
	"github.com/gin-gonic/gin"
)

func InitCategoryApiRouter(rootPath string, router *gin.Engine) {
	categoryApiRouter := router.Group(rootPath, middleware.JWTAuthMiddleWare())
	{
		categoryApiRouter.GET("/all_categories", api.GetAllCategories)
		categoryApiRouter.GET("/parent_categories", api.GetAllParentCategories)
		categoryApiRouter.GET("/categories", api.GetCategoriesByPage)
		categoryApiRouter.POST("/categories/article", api.AddArticleCategory)
		categoryApiRouter.PUT("/categories/article", api.UpdateArticleCategory)
		categoryApiRouter.POST("/categories/link", api.AddLinkCategory)
		categoryApiRouter.PUT("/categories/link", api.UpdateLinkCategory)
		categoryApiRouter.DELETE("/categories/:id", api.DeleteCategory)
		categoryApiRouter.DELETE("/categories", api.MultiDelCategories)
	}
}
