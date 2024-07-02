package router

import (
	"blog/common"
	"blog/middleware"
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.Use(middleware.Cache())
	router.Use(middleware.CORSMiddleware())
	router.Use(static.Serve("/", common.EmbedFolder(buildFS, "web/dist")))
	router.NoRoute(func(c *gin.Context) {
		SetNoCacheHeaders(c)
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPage)
	})
}
