package handlers

import (
	"aries/config/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FrontHandler struct {
}

// 首页
func (f *FrontHandler) IndexHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"blogVars": setting.BlogVars,
	})
}
