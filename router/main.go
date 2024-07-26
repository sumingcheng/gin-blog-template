package router

import (
	"embed"
	"github.com/gin-gonic/gin"
)

func SetRouter(
	router *gin.Engine,
	buildFS embed.FS,
	indexPage []byte,
) {
	// Other
	SetOtherRouter(router)
	// API
	SetApiRouter(router)
	// Web
	SetWebRouter(router, buildFS, indexPage)
}
