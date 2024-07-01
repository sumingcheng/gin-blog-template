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

//go:embed views/dist/*
var buildFS embed.FS

//go:embed views/dist/index.html
var indexPage []byte

func main() {
	//gin.SetMode(gin.ReleaseMode)   // 设置为发布模式
	//gin.Defaultwriter = io.Discard // 关闭gin的日志输出

	server := gin.Default()
	server.Use(middleware.Metric())

	router.SetRouter(server, buildFS, indexPage)

	err := server.Run("localhost:5678")
	if err != nil {
		return
	}
}
