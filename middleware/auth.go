package middleware

import (
	"blog/database"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	UidInToken = "uid"
)

var (
	KeyConfig = util.CreateConfig("key")
)

// GetUidFromJwt 从jwt里取出uid
func GetUidFromJwt(jwt string) int {
	_, payload, err := util.VerifyJwt(jwt, KeyConfig.GetString("jwt"))
	if err != nil {
		return 0
	}
	for k, v := range payload.UserDefined {
		if k == UidInToken {
			if val, err := strconv.ParseFloat(v, 64); err == nil {
				return int(val)
			}
		}
	}
	return 0
}

// AuthUid 身份认证中间件，无授权则返回禁止状态
func AuthUid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var uid int

		// 尝试从 Header 中获取 "auth_token"
		authToken := ctx.Request.Header.Get("auth_token")
		if authToken != "" {
			uid = GetUidFromJwt(authToken)
		}

		// 验证 uid，如果不大于 0 则认为用户未登录或登录已过期
		if uid <= 0 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "未登录或登录已过期",
			})
			ctx.Abort()
			return
		}

		// 验证通过，把登录的 uid 放入 ctx 中供后续中间件或处理函数使用
		ctx.Set("uid", uid)
		ctx.Next()
	}
}

// CheckRefreshToken 检查 refresh_token 是否有效
func CheckRefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, err := ctx.Cookie("refresh_token")
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "未找到 Token 请重新登录",
			})
			ctx.Abort()
			return
		}

		_, valid := database.VerifyRefreshToken(refreshToken)
		if !valid {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Token 已失效，请重新登录",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
