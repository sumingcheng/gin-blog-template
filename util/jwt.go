package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	UID int `json:"uid"`
	jwt.RegisteredClaims
}

// GenAuthToken 生成 auth JWT，内置 exp
func GenAuthToken(uid int, secret string, expire time.Duration) (string, error) {
	claims := AuthClaims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "blog",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// VerifyAuthToken 解析并校验 JWT（含 exp 自动校验），返回 uid
func VerifyAuthToken(tokenStr, secret string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("签名算法不匹配: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*AuthClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("无效的 token")
	}
	return claims.UID, nil
}

// GenRefreshToken 生成 32 字节随机 hex 串作为 refresh token
func GenRefreshToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
