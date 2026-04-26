package middleware

import (
	"blog/database"
	"blog/model"
	"blog/service"

	"github.com/gin-gonic/gin"
)

// AuthUid 身份认证中间件
func AuthUid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.Request.Header.Get("auth_token")
		uid := service.GetUidFromToken(authToken)
		if uid <= 0 {
			model.Fail(ctx, model.ErrUnauthorized)
			ctx.Abort()
			return
		}
		ctx.Set("uid", uid)
		ctx.Next()
	}
}

// CheckRefreshToken 检查 refresh_token 是否有效
func CheckRefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, err := ctx.Cookie("refresh_token")
		if err != nil {
			model.Fail(ctx, model.ErrUnauthorized)
			ctx.Abort()
			return
		}
		_, valid := database.VerifyRefreshToken(refreshToken)
		if !valid {
			model.Fail(ctx, model.ErrRefreshInvalid)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
