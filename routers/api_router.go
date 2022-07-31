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
	themeHandler := api.ThemeHandler{}

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

	ArticleApiRouter := router.Group(rootPath)
	{
		ArticleApiRouter.GET("/all_articles", middlewares.JWTAuth(), articleHandler.GetAllArticles)
		ArticleApiRouter.GET("/articles/:id", middlewares.JWTAuth(), articleHandler.GetArticleById)
		ArticleApiRouter.GET("/articles", middlewares.JWTAuth(), articleHandler.GetArticlesByPage)
		ArticleApiRouter.POST("/articles", middlewares.JWTAuth(), articleHandler.AddArticle)
		ArticleApiRouter.PUT("/articles", middlewares.JWTAuth(), articleHandler.UpdateArticle)
		ArticleApiRouter.DELETE("/articles/:id", middlewares.JWTAuth(), articleHandler.DeleteArticle)
		ArticleApiRouter.DELETE("/articles", middlewares.JWTAuth(), articleHandler.MultiDelArticles)
		ArticleApiRouter.POST("/articles/files", middlewares.JWTAuth(), articleHandler.ImportArticlesFromFiles)
		ArticleApiRouter.PATCH("/articles/recycle/:id", middlewares.JWTAuth(), articleHandler.RecycleOrRecoverArticle)
		ArticleApiRouter.PATCH("/articles/up", middlewares.JWTAuth(), articleHandler.MoveArticleUp)
		ArticleApiRouter.PATCH("/articles/down", middlewares.JWTAuth(), articleHandler.MoveArticleDown)
		ArticleApiRouter.POST("/articles/check", middlewares.Csrf(), articleHandler.CheckArticlePwd)
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
		commentRouter.PUT("/comments", middlewares.JWTAuth(), commentHandler.UpdateComment)
		commentRouter.DELETE("/comments/:id", middlewares.JWTAuth(), commentHandler.DeleteComment)
		commentRouter.DELETE("/comments", middlewares.JWTAuth(), commentHandler.MultiDelComments)
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
		sysSettingApiRouter.POST("/sys_setting/pic_bed/qubu", sysSettingHandler.SaveQubuSetting)
		sysSettingApiRouter.POST("/sys_setting/pic_bed/smms", sysSettingHandler.SaveSmmsSetting)
		sysSettingApiRouter.POST("/sys_setting/pic_bed/imgbb", sysSettingHandler.SaveImgbbSetting)
		sysSettingApiRouter.POST("/sys_setting/pic_bed/tencent_cos", sysSettingHandler.SaveTencentCosSetting)
		sysSettingApiRouter.POST("/sys_setting/comment/local", sysSettingHandler.SaveLocalCommentSetting)
		sysSettingApiRouter.POST("/sys_setting/comment/twikoo", sysSettingHandler.SaveTwikooSetting)
		sysSettingApiRouter.POST("/sys_setting/param", sysSettingHandler.SaveParamSetting)
		sysSettingApiRouter.POST("/sys_setting/social_info", sysSettingHandler.SaveSocialInfo)
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

	themeApiRouter := router.Group(rootPath)
	{
		themeApiRouter.GET("/all_themes", themeHandler.GetAllThemes)
		themeApiRouter.GET("/themes/:name", themeHandler.GetThemeByName)
		themeApiRouter.POST("/themes", middlewares.JWTAuth(), themeHandler.EnableTheme)
	}
}
