package api

import (
	"aries/handler/api"
	"github.com/gin-gonic/gin"
)

func InitCategoryApiRouter(rootPath string, router *gin.Engine) {
	categoryApiRouter := router.Group(rootPath)
	//categoryApiRouter.Use(middleware.JWTAuthMiddleWare()) // 加载 JWT 权限校验中间件
	{
		categoryApiRouter.GET("/categories/all", api.GetAllCategories)
		categoryApiRouter.GET("/categories", api.GetCategoriesByPage)
		categoryApiRouter.POST("/categories", api.AddCategory)
		categoryApiRouter.PUT("/categories/:id", api.UpdateCategory)
		categoryApiRouter.DELETE("/categories/:id", api.DeleteCategory)
		categoryApiRouter.DELETE("/categories", api.MultiDeleteCategory)
	}
}
