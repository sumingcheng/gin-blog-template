package router

import (
	_ "blog/docs"
	"blog/handler"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	api := router.Group("/api")

	// 公开接口
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)
	api.POST("/logout", handler.Logout)
	api.GET("/token", handler.GetAuthToken)

	// 博客公开接口
	api.GET("/blog/list", handler.BlogList)
	api.GET("/blog/:bid", handler.BlogDetail)

	// 需要登录的接口
	auth := api.Group("", middleware.AuthUid())
	{
		auth.POST("/user/password", handler.ChangePassword)
		auth.GET("/user/profile", handler.GetProfile)
		auth.POST("/blog/create", handler.BlogCreate)
		auth.POST("/blog/update", handler.BlogUpdate)
		auth.DELETE("/blog/:bid", handler.BlogDelete)
	}

	// 需要 refresh token 的接口
	api.POST("/blog/belong", middleware.CheckRefreshToken(), handler.BlogBelong)
}
