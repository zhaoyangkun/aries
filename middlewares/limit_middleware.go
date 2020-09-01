package middlewares

import "github.com/gin-gonic/gin"

// 限流中间件
func Limit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
