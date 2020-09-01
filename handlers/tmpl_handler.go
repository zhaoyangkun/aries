package handlers

import (
	"aries/config/setting"
	"aries/models"
	"aries/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TmplHandler struct {
}

// 首页
func (t *TmplHandler) IndexTmpl(ctx *gin.Context) {
	page := ctx.Param("page")
	pagination := utils.Pagination{}
	var articles []models.Article

	pagination.Page = 1
	if page != "" {
		p, err := strconv.ParseUint(page, 10, 0)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/error/404")
			return
		}
		pagination.Page = uint(p)
	}

	paramSetting, _ := models.SysSettingItem{}.GetBySysSettingName("参数设置")
	if size, ok := paramSetting["index_page_size"]; ok {
		s, err := strconv.ParseUint(size, 10, 0)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/error/500")
			return
		}
		pagination.Size = uint(s)
	} else {
		pagination.Size = 12
	}

	articles, total, _ := models.Article{}.GetByPage(&pagination, "", 1, 0)
	pageResult := utils.GetPageData(articles, total, pagination)

	var pages []int
	totalPages := pageResult["total_pages"].(uint)
	for i := 1; i <= int(totalPages); i++ {
		pages = append(pages, i)
	}

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"blogVars":    setting.BlogVars,
		"currentPage": int(pagination.Page),
		"pages":       pages,
		"articles":    pageResult["data"],
	})
}

// 文章详情页
func (t *TmplHandler) ArticleTmpl(ctx *gin.Context) {
	url := ctx.Param("url")
	article, _ := models.Article{}.GetByUrl(url)
	if article.Title == "" {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error/404")
		return
	}
	ctx.HTML(http.StatusOK, "article.tmpl", gin.H{
		"blogVars": setting.BlogVars,
		"article":  article,
	})
}

// 403 错误页
func (t *TmplHandler) Error403Tmpl(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "403.tmpl", nil)
}

// 404 错误页
func (t *TmplHandler) Error404Tmpl(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "404.tmpl", nil)
}

// 500 错误页
func (t *TmplHandler) Error500Tmpl(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "500.tmpl", nil)
}
