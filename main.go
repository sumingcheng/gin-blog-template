package main

import (
	"blog/handler"
	"blog/handler/middleware"
	"blog/util"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func init() {
	util.InitLog("log")
}

func main() {
	//gin.SetMode(gin.ReleaseMode)   // 设置为发布模式
	//gin.Defaultwriter = io.Discard // 关闭gin的日志输出

	router := gin.Default()

	router.Use(middleware.Metric())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/metrics", handler.Metric)

	// 静态资源和SPA配置
	router.Static("/static", "./dist")
	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html") // 托管SPA入口文件
	})

	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	router.POST("/login/submit", handler.Login)
	router.POST("/token", handler.GetAuthToken)

	router.GET("/blog/belong", handler.BlogBelong)
	router.GET("/blog/list/:uid", handler.BlogList)
	router.GET("/blog/:bid", handler.BlogDetail)
	router.POST("/blog/update", middleware.Auth(), handler.BlogUpdate)

	err := router.Run("localhost:5678")

	if err != nil {
		return
	}
}
