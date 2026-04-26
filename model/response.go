package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一 API 响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{Code: 0, Msg: "success", Data: data})
}

func Fail(ctx *gin.Context, err *BizError) {
	ctx.JSON(err.HTTPStatus, Response{Code: err.Code, Msg: err.Msg})
}

func FailMsg(ctx *gin.Context, err *BizError, msg string) {
	ctx.JSON(err.HTTPStatus, Response{Code: err.Code, Msg: msg})
}
