package router

import (
	"blog/common"
	"blog/middleware"
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setWebRouter(
	router *gin.Engine,
	buildFS embed.FS,
	indexPage []byte,
) {
	router.Use(middleware.Cache())
	router.Use(static.Serve("/", common.EmbedFolder(buildFS, "web/dist")))
	router.NoRoute(func(c *gin.Context) {
		middleware.SetNoCacheHeaders(c)
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPage)
	})
}
