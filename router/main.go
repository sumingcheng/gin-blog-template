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
	//其他
	SetOtherRouter(router)
	// 接口
	SetApiRouter(router)
	// web
	SetWebRouter(router, buildFS, indexPage)
}
