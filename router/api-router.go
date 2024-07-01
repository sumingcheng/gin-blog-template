package router

import (
	"blog/handler"
	"blog/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	// 文档
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 仪表盘
	apiRouter.GET("/metrics", handler.Metric)
	// 接口
	apiRouter.GET("/login", handler.Login)
	apiRouter.POST("/token", handler.GetAuthToken)
	apiRouter.GET("/blog/belong", handler.BlogBelong)
	apiRouter.GET("/blog/list/:uid", handler.BlogList)
	apiRouter.GET("/blog/:bid", handler.BlogDetail)
	apiRouter.POST("/blog/update", middleware.Auth(), handler.BlogUpdate)
}
