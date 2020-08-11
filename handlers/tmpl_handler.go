package handlers

import (
	"aries/config/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TmplHandler struct {
}

// 首页
func (f *TmplHandler) IndexHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"blogVars": setting.BlogVars,
	})
}
