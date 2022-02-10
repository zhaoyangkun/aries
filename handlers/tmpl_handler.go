package handlers

import (
	"aries/config/setting"
	"aries/log"
	"aries/middlewares"
	"aries/models"
	"aries/utils"
	"github.com/douyacun/gositemap"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TmplHandler struct {
}

// 模板全局变量
var (
	navs       []models.Nav
	categories []models.Category
	tags       []models.Tag
)

// InitTmplVars 初始化模板全局变量
func InitTmplVars() {
	navs, _ = models.Nav{}.GetAll()
	categories, _ = models.Category{}.GetAllByType(0)
	tags, _ = models.Tag{}.GetAllWithNoArticle()
}

// IndexTmpl 首页
func (t *TmplHandler) IndexTmpl(ctx *gin.Context) {
	page := ctx.Param("page")
	pagination := utils.Pagination{}
	var articles []models.Article

	pagination.Page = 1
	if page != "" {
		p, err := strconv.ParseUint(page, 10, 0)
		if err != nil {
			ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
				"blogVars": setting.BlogVars,
				"code":     "400",
				"msg":      "请求错误",
			})
			return
		}
		pagination.Page = uint(p)
	}

	paramSetting, _ := models.SysSettingItem{}.GetBySysSettingName("参数设置")
	if size, ok := paramSetting["index_page_size"]; ok {
		s, err := strconv.ParseUint(size, 10, 0)
		if err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
				"blogVars": setting.BlogVars,
				"code":     "500",
				"msg":      "服务器内部发生了错误",
			})
			return
		}
		pagination.Size = uint(s)
	} else {
		pagination.Size = 10
	}

	articles, total, _ := models.Article{}.GetByPage(&pagination, "", 1, 0)
	pageData := utils.GetPageData(articles, total, pagination)

	var pages []int
	totalPages := pageData["total_pages"].(uint)
	for i := 1; i <= int(totalPages); i++ {
		pages = append(pages, i)
	}

	ctx.HTML(http.StatusOK, utils.GetTheme()+"index.tmpl", gin.H{
		"blogVars":    setting.BlogVars,
		"navs":        navs,
		"categories":  categories,
		"tags":        tags,
		"articles":    articles,
		"currentPage": int(pagination.Page),
		"pages":       pages,
		"subTitle":    "",
		"pageSize":    pagination.Size,
		"totalCount":  total,
		"totalPages":  totalPages,
	})
}

// ArticleTmpl 文章详情页
func (t *TmplHandler) ArticleTmpl(ctx *gin.Context) {
	url := ctx.Param("url")

	article, _ := models.Article{}.GetByUrl(url)
	if article.Title == "" {
		ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
			"blogVars": setting.BlogVars,
			"code":     "400",
			"msg":      "请求错误",
		})
		return
	}
	if article.Pwd != "" {
		session := sessions.Default(ctx)
		flag := session.Get("article-" + strconv.Itoa(int(article.ID)))
		if flag == nil { // 未输入过密码，跳转到密码页面
			ctx.HTML(http.StatusOK, utils.GetTheme()+"pwd.tmpl", gin.H{
				"articleId": article.ID,
				"csrfToken": middlewares.CreateCsrfToken(ctx),
				"blogVars":  setting.BlogVars,
				"subTitle":  "私密文章",
			})
			return
		}
	}

	_ = article.UpdateVisitCount()
	preArticle, _ := models.Article{}.GetPrevious(article.OrderId, article.IsTop, true)
	nextArticle, _ := models.Article{}.GetNext(article.OrderId, article.IsTop, true)
	users, _ := models.User{}.GetAll()

	commentPlugInSetting, _ := models.SysSettingItem{}.GetBySysSettingName("评论组件设置")
	commentSetting, _ := models.SysSettingItem{}.GetBySysSettingName(commentPlugInSetting["plug_in"])

	ctx.HTML(http.StatusOK, utils.GetTheme()+"article.tmpl", gin.H{
		"blogVars":             setting.BlogVars,
		"navs":                 navs,
		"categories":           categories,
		"tags":                 tags,
		"article":              article,
		"preArticle":           preArticle,
		"nextArticle":          nextArticle,
		"user":                 users[0],
		"subTitle":             article.Title,
		"articleID":            article.ID,
		"pageID":               0,
		"commentType":          1,
		"commentPlugInSetting": commentPlugInSetting,
		"commentSetting":       commentSetting,
	})
}

// CategoryTmpl 分类页
func (t *TmplHandler) CategoryTmpl(ctx *gin.Context) {
	url := ctx.Param("url")
	page := ctx.Param("page")
	pagination := utils.Pagination{}
	var p = uint64(1)
	var err error

	//// 解决路由冲突
	//matchResult, _ := regexp.MatchString(`p/\d+`, url)
	//if matchResult {
	//	page = url[2:]
	//}
	if page != "" {
		p, err = strconv.ParseUint(page, 10, 0)
		if err != nil {
			ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
				"blogVars": setting.BlogVars,
				"code":     "400",
				"msg":      "请求错误",
			})
			return
		}
	}
	pagination.Page = uint(p)

	paramSetting, _ := models.SysSettingItem{}.GetBySysSettingName("参数设置")
	if size, ok := paramSetting["index_page_size"]; ok {
		s, err := strconv.ParseUint(size, 10, 0)
		if err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
				"blogVars": setting.BlogVars,
				"code":     "500",
				"msg":      "服务器内部发生了错误",
			})
			return
		}
		pagination.Size = uint(s)
	} else {
		pagination.Size = 3
	}

	if url == "" {
		articles, total, _ := models.Article{}.GetByPage(&pagination, "", 1, 0)
		pageData := utils.GetPageData(articles, total, pagination)

		var pages []int
		totalPages := pageData["total_pages"].(uint)
		for i := 1; i <= int(totalPages); i++ {
			pages = append(pages, i)
		}

		ctx.HTML(http.StatusOK, utils.GetTheme()+"category-list.tmpl", gin.H{
			"blogVars":     setting.BlogVars,
			"navs":         navs,
			"categories":   categories,
			"tags":         tags,
			"articles":     articles,
			"categoryName": "",
			"categoryUrl":  url,
			"currentPage":  int(pagination.Page),
			"pages":        pages,
			"subTitle":     "分类列表",
			"pageSize":     pagination.Size,
			"totalCount":   total,
			"totalPages":   totalPages,
		})
	} else {
		articles, name, total, _ := models.Article{}.GetByCategoryUrl(&pagination, url)
		pageData := utils.GetPageData(articles, total, pagination)

		var pages []int
		totalPages := pageData["total_pages"].(uint)
		for i := 1; i <= int(totalPages); i++ {
			pages = append(pages, i)
		}

		ctx.HTML(http.StatusOK, utils.GetTheme()+"category.tmpl", gin.H{
			"blogVars":     setting.BlogVars,
			"navs":         navs,
			"categories":   categories,
			"tags":         tags,
			"articles":     articles,
			"categoryName": name,
			"categoryUrl":  url,
			"currentPage":  int(pagination.Page),
			"pages":        pages,
			"subTitle":     name,
			"pageSize":     pagination.Size,
			"totalCount":   total,
			"totalPages":   totalPages,
		})
	}
}

// TagTmpl 标签页
func (t *TmplHandler) TagTmpl(ctx *gin.Context) {
	name := ctx.Param("name")
	page := ctx.Param("page")
	pagination := utils.Pagination{}
	var p = uint64(1)
	var err error

	if page != "" {
		p, err = strconv.ParseUint(page, 10, 0)
		if err != nil {
			ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
				"blogVars": setting.BlogVars,
				"code":     "400",
				"msg":      "请求错误",
			})
			return
		}
	}
	pagination.Page = uint(p)

	paramSetting, _ := models.SysSettingItem{}.GetBySysSettingName("参数设置")
	if size, ok := paramSetting["index_page_size"]; ok {
		s, err := strconv.ParseUint(size, 10, 0)
		if err != nil {
			log.Logger.Sugar().Error("error: ", err.Error())
			ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
				"blogVars": setting.BlogVars,
				"code":     "500",
				"msg":      "服务器内部发生了错误",
			})
			return
		}
		pagination.Size = uint(s)
	} else {
		pagination.Size = 10
	}

	if name == "" {
		articles, total, _ := models.Article{}.GetByPage(&pagination, "", 1, 0)
		pageData := utils.GetPageData(articles, total, pagination)

		var pages []int
		totalPages := pageData["total_pages"].(uint)
		for i := 1; i <= int(totalPages); i++ {
			pages = append(pages, i)
		}

		ctx.HTML(http.StatusOK, utils.GetTheme()+"tag-list.tmpl", gin.H{
			"blogVars":    setting.BlogVars,
			"navs":        navs,
			"categories":  categories,
			"tags":        tags,
			"articles":    articles,
			"tagName":     name,
			"currentPage": int(pagination.Page),
			"pages":       pages,
			"subTitle":    "标签列表",
			"pageSize":    pagination.Size,
			"totalCount":  total,
			"totalPages":  totalPages,
		})
	} else {
		articles, total, _ := models.Article{}.GetByTagName(&pagination, name)
		pageData := utils.GetPageData(articles, total, pagination)

		var pages []int
		totalPages := pageData["total_pages"].(uint)
		for i := 1; i <= int(totalPages); i++ {
			pages = append(pages, i)
		}

		ctx.HTML(http.StatusOK, utils.GetTheme()+"tag.tmpl", gin.H{
			"blogVars":    setting.BlogVars,
			"navs":        navs,
			"categories":  categories,
			"tags":        tags,
			"articles":    articles,
			"tagName":     name,
			"currentPage": int(pagination.Page),
			"pages":       pages,
			"subTitle":    name,
			"pageSize":    pagination.Size,
			"totalCount":  total,
			"totalPages":  totalPages,
		})
	}
}

// ArchiveTmpl 归档页
func (t *TmplHandler) ArchiveTmpl(ctx *gin.Context) {
	archives, _ := models.Archive{}.GetAll()
	articles, _ := models.Article{}.GetAll()

	ctx.HTML(http.StatusOK, utils.GetTheme()+"archive.tmpl", gin.H{
		"blogVars":   setting.BlogVars,
		"navs":       navs,
		"categories": categories,
		"tags":       tags,
		"archives":   archives,
		"articles":   articles,
		"subTitle":   "归档",
	})
}

// LinkTmpl 友链页
func (t *TmplHandler) LinkTmpl(ctx *gin.Context) {
	linkCategories, _ := models.Category{}.GetAllByType(1)
	links, _ := models.Link{}.GetAll()

	ctx.HTML(http.StatusOK, utils.GetTheme()+"link.tmpl", gin.H{
		"blogVars":       setting.BlogVars,
		"navs":           navs,
		"categories":     categories,
		"tags":           tags,
		"linkCategories": linkCategories,
		"links":          links,
		"subTitle":       "友链",
	})
}

// JournalTmpl 日志页
func (t *TmplHandler) JournalTmpl(ctx *gin.Context) {
	journals, _ := models.Journal{}.GetAll()
	users, _ := models.User{}.GetAll()

	ctx.HTML(http.StatusOK, utils.GetTheme()+"journal.tmpl", gin.H{
		"blogVars":   setting.BlogVars,
		"navs":       navs,
		"categories": categories,
		"tags":       tags,
		"journals":   journals,
		"user":       users[0],
		"subTitle":   "日志",
	})
}

// GalleryTmpl 图库
func (t *TmplHandler) GalleryTmpl(ctx *gin.Context) {
	photoCategories, _ := models.Category{}.GetGalleryCategories()
	photos, _ := models.Gallery{}.GetAll()

	ctx.HTML(http.StatusOK, utils.GetTheme()+"photo.tmpl", gin.H{
		"blogVars":        setting.BlogVars,
		"navs":            navs,
		"categories":      categories,
		"photoCategories": photoCategories,
		"photos":          photos,
		"subTitle":        "图库",
	})
}

// CustomTmpl 自定义页
func (t *TmplHandler) CustomTmpl(ctx *gin.Context) {
	url := ctx.Param("url")

	page, _ := models.Page{}.GetByUrl(url)
	if page.Title == "" {
		ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
			"blogVars": setting.BlogVars,
			"code":     "400",
			"msg":      "请求错误",
		})
		return
	}

	commentPlugInSetting, _ := models.SysSettingItem{}.GetBySysSettingName("评论组件设置")
	commentSetting, _ := models.SysSettingItem{}.GetBySysSettingName(commentPlugInSetting["plug_in"])

	ctx.HTML(http.StatusOK, utils.GetTheme()+"custom.tmpl", gin.H{
		"blogVars":             setting.BlogVars,
		"navs":                 navs,
		"categories":           categories,
		"tags":                 tags,
		"page":                 page,
		"subTitle":             page.Title,
		"articleID":            0,
		"pageID":               page.ID,
		"commentType":          4,
		"commentPlugInSetting": commentSetting,
		"commentSetting":       commentSetting,
	})
}

// SiteMapXml 站点地图
func (t *TmplHandler) SiteMapXml(ctx *gin.Context) {
	st := gositemap.NewSiteMap()
	st.SetPretty(true)

	articles, _ := models.Article{}.GetAll()
	for _, article := range articles {
		url := gositemap.NewUrl()
		url.SetLoc(setting.BlogVars.ContextPath + "/articles/" + article.URL)
		url.SetLastmod(article.UpdatedAt)
		url.SetChangefreq("daily")
		url.SetPriority(1)
		st.AppendUrl(url)
	}

	ctx.XML(http.StatusOK, st)
}

// SearchTmpl 搜索
func (t *TmplHandler) SearchTmpl(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	page := ctx.Param("page")

	pagination := utils.Pagination{Page: 1}
	if page != "" {
		p, err := strconv.ParseUint(page, 10, 0)
		if err != nil {
			ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
				"blogVars": setting.BlogVars,
				"code":     "400",
				"msg":      "请求错误",
			})
			return
		}
		pagination.Page = uint(p)
	}
	articles, total, _ := models.Article{}.GetByPage(&pagination, keyword, 5, 0)
	pageData := utils.GetPageData(articles, total, pagination)

	var pages []int
	totalPages := pageData["total_pages"].(uint)
	for i := 1; i <= int(totalPages); i++ {
		pages = append(pages, i)
	}

	ctx.HTML(http.StatusOK, utils.GetTheme()+"search.tmpl", gin.H{
		"blogVars":    setting.BlogVars,
		"navs":        navs,
		"categories":  categories,
		"tags":        tags,
		"articles":    articles,
		"currentPage": int(pagination.Page),
		"pages":       pages,
		"subTitle":    "关于「" + keyword + "」的搜索结果",
		"keyword":     keyword,
		"pageSize":    pagination.Size,
		"totalCount":  total,
		"totalPages":  totalPages,
	})
}
