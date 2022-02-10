package middlewares

import (
	"aries/config/setting"
	"aries/utils"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger 接收 gin 框架默认的日志
func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		status := ctx.Writer.Status()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		ip := ctx.ClientIP()
		if ip == "::1" {
			ip = "127.0.0.1"
		}
		query := ctx.Request.URL.RawQuery
		if query != "" {
			query = "?" + query
		}
		//userAgent := c.Request.UserAgent()
		//errors := c.Errors.ByType(gin.ErrorTypePrivate).String()

		ctx.Next()

		end := time.Now()
		cost := end.Sub(start)

		// 输出日志（终端 + 文件）
		//logger.Info(
		//	"",
		//	zap.Int("status", status),
		//	zap.String("method", method),
		//	zap.String("path", path),
		//	zap.String("ip", ip),
		//	zap.Duration("cost", cost),
		//)
		logger.Sugar().Infof("| %3d | %13v | %s | %s | %s%s |", status, cost, ip, method, path, query)

		// 错误页面跳转
		if !strings.Contains(path, "/api") && !strings.Contains(path, "/static") {
			switch status {
			case 400:
				ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
					"blogVars": setting.BlogVars,
					"code":     "400",
					"msg":      "请求错误",
				})
			case 403:
				ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
					"blogVars": setting.BlogVars,
					"code":     "403",
					"msg":      "您无权访问该页面",
				})
			case 404:
				ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
					"blogVars": setting.BlogVars,
					"code":     "404",
					"msg":      "您访问的页面不存在",
				})
			case 500:
				ctx.HTML(http.StatusOK, utils.GetTheme()+"error.tmpl", gin.H{
					"blogVars": setting.BlogVars,
					"code":     "500",
					"msg":      "服务器内部发生了错误",
				})
			}
		}
	}
}

// Recover recover 掉项目可能出现的 panic，并使用 zap 记录相关日志
func Recover(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					/*					logger.Error("[Recovery from panic]",
										zap.Any("error", err),
										zap.Stack("stack"),
										//zap.String("stack", string()),
									)*/
					logger.Sugar().Errorf("| %s | %s |", err, debug.Stack())
				} else {
					/*					logger.Error("[Recovery from panic]",
										zap.Any("error", err),
										zap.String("request", string(httpRequest)),
									)*/
					logger.Sugar().Errorf("| %s | %s |", err, httpRequest)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
