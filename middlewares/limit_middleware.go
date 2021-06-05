package middlewares

import (
	"aries/utils"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

var bucket *ratelimit.Bucket

// InitBucket 初始化 Bucket
func InitBucket(fillInternal time.Duration, capacity int64) {
	bucket = ratelimit.NewBucket(fillInternal, capacity)
}

// Limiter 限流中间件
func Limiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.Forbidden,
				Msg:  "您的访问过于频繁！",
				Data: nil,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
