package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware
// Access-Control-Allow-Origin: 允许所有域进行跨域请求（*）。在生产环境中，你可能希望将其设置为特定的域名。
// Access-Control-Allow-Credentials: 允许携带证书（如 cookies）。
// Access-Control-Allow-Headers: 指定了浏览器允许访问的头部。
// Access-Control-Allow-Methods: 定义允许的 HTTP 方法。
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
