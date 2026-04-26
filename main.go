package main

import (
	"blog/middleware"
	"blog/router"
	"blog/util"
	"context"
	"embed"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	util.InitLog("log")
	util.InitTranslator("zh")
}

var (
	//go:embed web/dist/*
	buildFS embed.FS
	//go:embed web/dist/index.html
	indexPage []byte
	ginConfig = util.CreateConfig("gin")
)

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
	server := gin.Default()

	if err := server.SetTrustedProxies(ginConfig.GetStringSlice("trustedProxies")); err != nil {
		log.Fatalf("设置信任代理失败: %v", err)
	}

	server.Use(middleware.RequestID())
	server.Use(middleware.RequestLogger())
	server.Use(middleware.RateLimit(100, 200))
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.Metric())
	router.SetRouter(server, buildFS, indexPage)

	// 优雅关停
	addr := ginConfig.GetString("port")
	srv := &http.Server{Addr: addr, Handler: server}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		util.LogRus.Infof("服务启动: %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务启动失败: %v", err)
		}
	}()

	<-ctx.Done()
	util.LogRus.Info("收到关停信号，等待请求处理完成...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("服务强制关停: %v", err)
	}
	util.LogRus.Info("服务已关停")
}
