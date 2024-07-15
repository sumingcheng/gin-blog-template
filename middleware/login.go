package middleware

import (
	"blog/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckIfAlreadyLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, err := ctx.Cookie("auth_token")
		if err == nil {
			authToken, valid := database.VerifyRefreshToken(refreshToken)
			if valid {
				ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "repeatedLogin", "token": authToken})
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
