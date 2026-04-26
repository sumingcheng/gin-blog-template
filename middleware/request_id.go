package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

const RequestIDKey = "X-Request-ID"

// RequestID 为每个请求生成或透传请求 ID
func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid := ctx.GetHeader(RequestIDKey)
		if rid == "" {
			rid = generateID()
		}
		ctx.Set(RequestIDKey, rid)
		ctx.Header(RequestIDKey, rid)
		ctx.Next()
	}
}

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
