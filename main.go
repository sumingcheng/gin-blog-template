package main

import (
	"blog/database"
	"blog/middleware"
	"blog/router"
	"blog/util"
	"embed"
	"github.com/gin-gonic/gin"
)

func init() {
	util.InitLog("log")
}

var (
	//go:embed web/dist/*
	buildFS embed.FS
	//go:embed web/dist/index.html
	indexPage []byte
	ginConfig = util.CreateConfig("gin")
)

// 添加注释以描述 server 信息
// @title           Swagger Example API
// @version         2.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	//gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	//gin.Defaultwriter = io.Discard // 关闭gin的日志输出,所有的日志都会被丢弃
	database.AutoMigrate()

	server := gin.Default()
	err := server.SetTrustedProxies(ginConfig.GetStringSlice("trustedProxies"))
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.Metric())
	// Router
	router.SetRouter(server, buildFS, indexPage)

	err = server.Run(ginConfig.GetString("port"))
	if err != nil {
		return
	}
}
