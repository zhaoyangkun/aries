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
		categoryApiRouter.POST("/categories", api.AddCategory)
		categoryApiRouter.PUT("/categories", api.UpdateCategory)
		categoryApiRouter.DELETE("/categories/:id", api.DeleteCategory)
		categoryApiRouter.DELETE("/categories", api.MultiDelCategories)
	}
}
