package middleware

import (
	"blog/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var AllowArr = util.CreateConfig("gin").GetStringSlice("allowOrigins")

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     AllowArr,                                                                                     // 允许所有源
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},                                 // 允许的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Auth_token", "Authorization", "Refresh_token"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                                                                   // 允许暴露给浏览器的响应头
		AllowCredentials: true,                                                                                         // 允许凭证
		MaxAge:           12 * time.Hour,                                                                               // 预检请求的缓存时间
	})
}
