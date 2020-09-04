package handlers

import (
	"aries/config/setting"
	"aries/log"
	"aries/models"
	"aries/utils"
	"errors"
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
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("请求参数错误"))
			return
		}
		pagination.Page = uint(p)
	}

	paramSetting, _ := models.SysSettingItem{}.GetBySysSettingName("参数设置")
	if size, ok := paramSetting["index_page_size"]; ok {
		s, err := strconv.ParseUint(size, 10, 0)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("请求参数错误"))
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
		log.Logger.Debug("error 400")
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.HTML(http.StatusOK, "article.tmpl", gin.H{
		"blogVars": setting.BlogVars,
		"article":  article,
	})
}
