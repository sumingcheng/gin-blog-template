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
	// API
	SetApiRouter(router)
	// Other
	SetOtherRouter(router)
	// Web
	SetWebRouter(router, buildFS, indexPage)
}
