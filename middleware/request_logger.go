package middleware

import (
	"blog/util"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger 记录每个请求的关键信息
func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		latency := time.Since(start)
		rid, _ := ctx.Get(RequestIDKey)

		util.LogRus.Infof("[%s] %s %s %d %v",
			rid, ctx.Request.Method, ctx.Request.URL.Path,
			ctx.Writer.Status(), latency,
		)
	}
}
