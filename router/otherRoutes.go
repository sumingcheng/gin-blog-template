package router

import (
	"blog/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetOtherRouter(router *gin.Engine) {
	// 文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 仪表盘
	router.GET("/metrics", handler.Metric)
}
