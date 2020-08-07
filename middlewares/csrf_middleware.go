package middlewares

import (
	"aries/config/setting"
	"aries/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//无需 csrfToken 验证的请求方法
var allowedMethods = []string{"GET", "HEAD", "OPTIONS", "TRACE"}

// csrf 校验中间件
func CsrfMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if verifyCsrfToken(ctx) {
			ctx.Next()
		} else {
			result := utils.Result{
				Code: http.StatusForbidden,
				Msg:  "Csrf Forbidden",
				Data: nil,
			}
			ctx.JSON(result.Code, result)
			ctx.Abort()
		}
	}
}

//创建 csrfToken
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

//校验 csrfToken
func verifyCsrfToken(ctx *gin.Context) bool {
	// 请求方法无需校验 csrfToken
	if utils.IsContain(allowedMethods, ctx.Request.Method) {
		return true
	}
	referer := ctx.Request.Referer()
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
