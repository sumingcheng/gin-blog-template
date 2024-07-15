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
	// 优先尝试从request header里获取名为"auth_token"的token
	token := ctx.Request.Header.Get("auth_token")
	// 如果头部没有"auth_token"，尝试从Cookie里的"refresh_token"获取，通常用于刷新会话
	if token == "" {
		for _, cookie := range ctx.Request.Cookies() {
			if cookie.Name == "refresh_token" {
				token = cookie.Value
				break // 找到refresh_token后停止循环
			}
		}
	}

	if token == "" {
		return 0
	}

	return GetUidFromJwt(token)
}

// Auth 身份认证中间件，无授权则返回禁止状态
func Auth() gin.HandlerFunc {
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
