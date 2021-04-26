package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

var bucket *ratelimit.Bucket

func InitBucket(fillInternal time.Duration, capacity int64) {
	bucket = ratelimit.NewBucket(fillInternal, capacity)
}

func Limiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			ctx.JSON(http.StatusOK, "您的访问过于频繁！")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
