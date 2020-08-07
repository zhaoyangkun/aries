package handler

import (
	"aries/config/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页
func IndexHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"blogVars": setting.BlogVars,
	})
}
