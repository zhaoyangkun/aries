package middlewares

import (
	"aries/config/setting"
	"aries/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 无需 csrfToken 验证的请求方法
var allowedMethods = []string{"GET", "HEAD", "OPTIONS", "TRACE"}

// Csrf 校验 csrf 中间件
func Csrf() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if verifyCsrfToken(ctx) {
			ctx.Next()
		} else {
			result := utils.Result{
				Code: utils.Forbidden,
				Msg:  "Csrf Forbidden",
				Data: nil,
			}

			ctx.JSON(http.StatusOK, result)
			ctx.Abort()
		}
	}
}

// CreateCsrfToken 创建 csrfToken
func CreateCsrfToken(ctx *gin.Context) (csrfToken string) {
	session := sessions.Default(ctx)
	csrfToken = utils.GetSessionStr(session, "csrfToken")

	if csrfToken == "" {
		nowTime := time.Now().String()           //获取当前时间
		csrfToken, _ = utils.EncryptPwd(nowTime) //根据当前时间生成 csrfToken
		session.Set("csrfToken", csrfToken)      //将 csrfToken 保存到 session 中
		_ = session.Save()
	}

	return
}

// 校验 csrfToken
func verifyCsrfToken(ctx *gin.Context) bool {
	// 请求方法无需校验 csrfToken
	if utils.IsContain(allowedMethods, ctx.Request.Method) {
		return true
	}

	referer := ctx.Request.Referer()
	if strings.Contains(referer, "localhost") || strings.Contains(referer, "127.0.0.1") {
		referer = "127.0.0.1"
	} else {
		referer = strings.Replace(referer, "//", "/", 1)
		referer = strings.Split(referer, "/")[1]
	}

	// 校验 Referer
	if !utils.IsContain(setting.Config.Server.AllowedRefers, referer) {
		return false
	}

	// 从请求头部获取 csrfToken
	csrfToken := ctx.Request.Header.Get("csrfToken")

	// 头部中无，则从 form 中获取 csrfToken
	if csrfToken == "" {
		csrfToken = ctx.Request.FormValue("csrfToken")
	}

	// 获取 session 中的 csrfToken
	correctCsrfToken := utils.GetSessionStr(sessions.Default(ctx), "csrfToken")

	// 校验 csrfToken
	if (csrfToken == "") || (csrfToken != correctCsrfToken) {
		return false
	}

	return true
}
