package router

import (
	_ "blog/docs"
	"blog/handler"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	api := router.Group("/api")
	// API 接口
	api.POST("/login", handler.Login)
	api.POST("/token", handler.GetAuthToken)
	api.POST("/blog/belong", handler.BlogBelong)
	api.GET("/blog/list/:uid", handler.BlogList)
	api.GET("/blog/:bid", handler.BlogDetail)
	api.POST("/blog/update", middleware.VerifyLogin(), handler.BlogUpdate)
}
