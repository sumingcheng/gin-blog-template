package main

import (
	"blog/middleware"
	"blog/router"
	"blog/util"
	"embed"
	"github.com/gin-gonic/gin"
)

func init() {
	util.InitLog("log")
}

//go:embed web/dist/*
var buildFS embed.FS

//go:embed web/dist/index.html
var indexPage []byte

func main() {
	//gin.SetMode(gin.ReleaseMode)   // 设置为发布模式
	//gin.Defaultwriter = io.Discard // 关闭gin的日志输出

	server := gin.Default()
	trustedProxies := []string{"0.0.0.0/0"}
	err := server.SetTrustedProxies(trustedProxies)
	server.Use(middleware.Metric())

	router.SetRouter(server, buildFS, indexPage)

	err = server.Run("localhost:5678")
	if err != nil {
		return
	}
}
