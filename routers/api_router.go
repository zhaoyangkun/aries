package routers

import (
	"aries/handlers/api"
	"aries/middlewares"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (a *ApiRouter) InitApiRouter(rootPath string, router *gin.Engine) {
	authHandler := api.AuthHandler{}
	userHandler := api.UserHandler{}
	articleHandler := api.ArticleHandler{}
	categoryHandler := api.CategoryHandler{}
	tagHandler := api.TagHandler{}
	commentHandler := api.CommentHandler{}
	linkHandler := api.LinkHandler{}
	navHandler := api.NavHandler{}
	sysSettingHandler := api.SysSettingHandler{}
	pictureHandler := api.PictureHandler{}
	journalHandler := api.JournalHandler{}
	galleryHandler := api.GalleryHandler{}
	pageHandler := api.PageHandler{}

	authApiRouter := router.Group(rootPath)
	{
		authApiRouter.POST("/auth/login", authHandler.Login)
		authApiRouter.POST("/auth/register", authHandler.Register)
		authApiRouter.GET("/auth/captcha", authHandler.CreateCaptcha)
		authApiRouter.POST("/auth/pwd/forget", authHandler.ForgetPwd)
		authApiRouter.POST("/auth/pwd/reset", authHandler.ResetPwd)
	}

	userApiRouter := router.Group(rootPath)
	{
		userApiRouter.GET("/all_users", userHandler.GetAllUsers)
		userApiRouter.PUT("/users", userHandler.UpdateUser)
		userApiRouter.PUT("/users/pwd", userHandler.UpdateUserPwd)
	}

	ArticleApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		ArticleApiRouter.GET("/all_articles", articleHandler.GetAllArticles)
		ArticleApiRouter.GET("/articles/:id", articleHandler.GetArticleById)
		ArticleApiRouter.GET("/articles", articleHandler.GetArticlesByPage)
		ArticleApiRouter.POST("/articles", articleHandler.AddArticle)
		ArticleApiRouter.PUT("/articles", articleHandler.UpdateArticle)
		ArticleApiRouter.DELETE("/articles/:id", articleHandler.DeleteArticle)
		ArticleApiRouter.DELETE("/articles", articleHandler.MultiDelArticles)
		ArticleApiRouter.POST("/articles/files", articleHandler.ImportArticlesFromFiles)
		ArticleApiRouter.PATCH("/articles/up", articleHandler.MoveArticleUp)
		ArticleApiRouter.PATCH("/articles/down", articleHandler.MoveArticleDown)
	}

	categoryApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		categoryApiRouter.GET("/all_categories", categoryHandler.GetAllCategories)
		categoryApiRouter.GET("/parent_categories", categoryHandler.GetAllParentCategories)
		categoryApiRouter.GET("/categories", categoryHandler.GetCategoriesByPage)
		categoryApiRouter.POST("/categories/article", categoryHandler.AddArticleCategory)
		categoryApiRouter.PUT("/categories/article", categoryHandler.UpdateArticleCategory)
		categoryApiRouter.POST("/categories/link", categoryHandler.AddLinkCategory)
		categoryApiRouter.PUT("/categories/link", categoryHandler.UpdateLinkCategory)
		categoryApiRouter.POST("/categories/gallery", categoryHandler.AddGalleryCategory)
		categoryApiRouter.PUT("/categories/gallery", categoryHandler.UpdateGalleryCategory)
		categoryApiRouter.DELETE("/categories/:id", categoryHandler.DeleteCategory)
		categoryApiRouter.DELETE("/categories", categoryHandler.MultiDelCategories)
	}

	tagApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		tagApiRouter.GET("/all_tags", tagHandler.GetAllTags)
		tagApiRouter.GET("/tags", tagHandler.GetTagsByPage)
		tagApiRouter.GET("/tags/:id", tagHandler.GetTagById)
		tagApiRouter.POST("/tags", tagHandler.AddTag)
		tagApiRouter.PUT("/tags", tagHandler.UpdateTag)
		tagApiRouter.DELETE("/tags/:id", tagHandler.DeleteTag)
		tagApiRouter.DELETE("/tags", tagHandler.MultiDelTags)
	}

	commentRouter := router.Group(rootPath)
	{
		commentRouter.GET("/all_comments", commentHandler.GetAllComments)
		commentRouter.GET("/comments", commentHandler.GetCommentsByPage)
		commentRouter.POST("/comments", commentHandler.AddComment)
		commentRouter.PUT("/comments", commentHandler.UpdateComment, middlewares.JWTAuth())
		commentRouter.DELETE("/comments/:id", commentHandler.DeleteComment, middlewares.JWTAuth())
		commentRouter.DELETE("/comments", commentHandler.MultiDelComments, middlewares.JWTAuth())
	}

	linkApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		linkApiRouter.GET("/all_links", linkHandler.GetAllLinks)
		linkApiRouter.GET("/links", linkHandler.GetLinksByPage)
		linkApiRouter.POST("/links", linkHandler.CreateLink)
		linkApiRouter.PUT("/links", linkHandler.UpdateLink)
		linkApiRouter.DELETE("/links/:id", linkHandler.DeleteLink)
		linkApiRouter.DELETE("/links", linkHandler.MultiDelLinks)
	}

	navApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		navApiRouter.GET("/navs", navHandler.GetAllNavs)
		navApiRouter.POST("/navs", navHandler.CreateNav)
		navApiRouter.PUT("/navs", navHandler.UpdateNav)
		navApiRouter.PATCH("/navs/:type/up/:order_id", navHandler.MoveNavUp)
		navApiRouter.PATCH("/navs/:type/down/:order_id", navHandler.MoveNavDown)
		navApiRouter.DELETE("/navs/:id", navHandler.DeleteNav)
		navApiRouter.DELETE("/navs", navHandler.MultiDelNavs)
	}

	sysSettingApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		sysSettingApiRouter.GET("/sys_setting/blog_vars", sysSettingHandler.GetBlogVars)
		sysSettingApiRouter.GET("/sys_setting/items", sysSettingHandler.GetSysSettingItem)
		sysSettingApiRouter.GET("/sys_setting/index_info", sysSettingHandler.GetAdminIndexData)
		sysSettingApiRouter.POST("/sys_setting/site", sysSettingHandler.SaveSiteSetting)
		sysSettingApiRouter.POST("/sys_setting/smtp", sysSettingHandler.SaveSMTPSetting)
		sysSettingApiRouter.POST("/sys_setting/email/test", sysSettingHandler.SendTestEmail)
		sysSettingApiRouter.POST("/sys_setting/pic_bed/smms", sysSettingHandler.SaveSmmsSetting)
		sysSettingApiRouter.POST("/sys_setting/pic_bed/imgbb", sysSettingHandler.SaveImgbbSetting)
		sysSettingApiRouter.POST("/sys_setting/pic_bed/tencent_cos", sysSettingHandler.SaveTencentCosSetting)
		sysSettingApiRouter.POST("/sys_setting/comment", sysSettingHandler.SaveCommentSetting)
		sysSettingApiRouter.POST("/sys_setting/param", sysSettingHandler.SaveParamSetting)
	}

	imgApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		imgApiRouter.GET("/images", pictureHandler.GetPicturesByPage)
		imgApiRouter.POST("/images/attachment/upload", pictureHandler.UploadImgToAttachment)
		imgApiRouter.DELETE("/images", pictureHandler.MultiDelPictures)
	}

	journalApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		journalApiRouter.GET("/all_journals", journalHandler.GetAllJournals)
		journalApiRouter.GET("/journals/:id", journalHandler.GetJournalById)
		journalApiRouter.GET("/journals", journalHandler.GetJournalsByPage)
		journalApiRouter.POST("/journals", journalHandler.CreateJournal)
		journalApiRouter.PUT("/journals", journalHandler.UpdateJournal)
		journalApiRouter.DELETE("/journals", journalHandler.MultiDelJournals)
	}

	galleryApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		galleryApiRouter.GET("/all_galleries", galleryHandler.GetAllGalleries)
		galleryApiRouter.GET("/galleries/:id", galleryHandler.GetGalleryById)
		galleryApiRouter.GET("/galleries", galleryHandler.GetGalleriesByPage)
		galleryApiRouter.POST("/galleries", galleryHandler.CreateGallery)
		galleryApiRouter.PUT("/galleries", galleryHandler.UpdateGallery)
		galleryApiRouter.DELETE("/galleries", galleryHandler.MultiDelGalleries)
	}

	pageApiRouter := router.Group(rootPath, middlewares.JWTAuth())
	{
		pageApiRouter.GET("/all_pages", pageHandler.GetAllPages)
		pageApiRouter.GET("/pages", pageHandler.GetPagesByPage)
		pageApiRouter.POST("/pages", pageHandler.CreatePage)
		pageApiRouter.PUT("/pages", pageHandler.UpdatePage)
		pageApiRouter.DELETE("/pages", pageHandler.MultiDelPages)
	}
}
