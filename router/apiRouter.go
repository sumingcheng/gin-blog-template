package router

import (
	_ "blog/docs"
	"blog/handler"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	// 接口
	apiRouter.POST("/login", handler.Login)
	apiRouter.POST("/token", handler.GetAuthToken)

	apiRouter.GET("/blog/belong", handler.BlogBelong)
	apiRouter.GET("/blog/list/:uid", handler.BlogList)
	apiRouter.GET("/blog/:bid", handler.BlogDetail)
	apiRouter.POST("/blog/update", middleware.Auth(), handler.BlogUpdate)
}
