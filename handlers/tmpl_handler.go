package handlers

import (
	"aries/config/setting"
	"aries/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TmplHandler struct {
}

// 首页
func (t *TmplHandler) IndexHTML(ctx *gin.Context) {
	log.Logger.Sugar().Info("blogVars: ", setting.BlogVars)
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"blogVars": setting.BlogVars,
	})
}
