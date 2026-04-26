package router

import (
	"blog/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetOtherRouter(router *gin.Engine) {
	router.GET("/health", handler.Health)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/metrics", handler.Metric)
}
