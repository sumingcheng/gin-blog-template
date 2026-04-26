package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimit 全局令牌桶限流，r 为每秒允许的请求数，b 为突发容量
func RateLimit(r rate.Limit, b int) gin.HandlerFunc {
	limiter := rate.NewLimiter(r, b)
	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"code": 42900,
				"msg":  "请求过于频繁，请稍后再试",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
