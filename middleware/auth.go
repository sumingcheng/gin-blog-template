package middleware

import (
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

// GetLoginUid 从header里获取jwt, 从而得出uid
func GetLoginUid(ctx *gin.Context) int {
	// 尝试从request header获取"auth_token"
	authToken := ctx.Request.Header.Get("auth_token")
	if authToken != "" {
		return GetUidFromJwt(authToken)
	}

	// 若"auth_token"不存在，尝试从Cookie获取"refresh_token"
	refreshCookie, err := ctx.Request.Cookie("refresh_token")
	if err == nil && refreshCookie != nil {
		return GetUidFromJwt(refreshCookie.Value)
	}

	// 如果两者都没有找到有效的token，返回0
	return 0
}

// VerifyLogin 身份认证中间件，无授权则返回禁止状态
func VerifyLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginUid := GetLoginUid(ctx)
		if loginUid <= 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusForbidden,
				"msg":  "未登录或登录已过期",
			})

			ctx.Abort()
		} else {
			ctx.Set("uid", loginUid) // 把登录的uid放入ctx中
			ctx.Next()
		}
	}
}
