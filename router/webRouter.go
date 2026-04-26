package router

import (
	"blog/common"
	"blog/middleware"
	"embed"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func SetNoCacheHeaders(c *gin.Context) {
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
}

func SetWebRouter(
	router *gin.Engine,
	buildFS embed.FS,
	indexPage []byte,
) {
	// Cache 只对静态资源生效，不影响 API
	staticGroup := router.Group("/", middleware.Cache())
	staticGroup.Use(static.Serve("/", common.EmbedFolder(buildFS, "web/dist")))

	router.NoRoute(func(c *gin.Context) {
		SetNoCacheHeaders(c)
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPage)
	})
}
