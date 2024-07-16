package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type JwtHeader struct {
	Algo string `json:"alg"` // 使用的算法, 比如HMAC SHA256
	Type string `json:"typ"` // Token的类型，这里是JWT
}

type JwtPayload struct {
	ID          string            `json:"jti"` // Token的ID
	Issue       string            `json:"iss"` // 发行者
	Audience    string            `json:"aud"` // 接收方
	Subject     string            `json:"sub"` // 主题
	IssueAt     int64             `json:"iat"` // 发行时间
	NotBefore   int64             `json:"nbf"` // 生效时间
	Expiration  int64             `json:"exp"` // 过期时间
	UserDefined map[string]string `json:"ud"`  // 用户自定义的数据
}

type CustomClaims struct {
	UserID int `json:"UserID"`
	jwt.RegisteredClaims
}

var (
	DefaultHeader = JwtHeader{
		Algo: "HS256",
		Type: "JWT",
	}
	key = []byte("ssss")
)

func GetRefreshToken() (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Blog_RefreshToken",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func GenJWT(header JwtHeader, payload JwtPayload, secret string) (string, error) {
	part1, err := MarshalAndEncode(header)
	if err != nil {
		return "", err
	}
	part2, err := MarshalAndEncode(payload)
	if err != nil {
		return "", err
	}
	signature := GenerateSignature(part1+"."+part2, secret)
	return part1 + "." + part2 + "." + signature, nil
}

func VerifyJwt(token, secret string) (*JwtHeader, *JwtPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, fmt.Errorf("token格式不正确")
	}
	if !CheckSignature(parts[0], parts[1], parts[2], secret) {
		return nil, nil, fmt.Errorf("签名不匹配")
	}
	var header JwtHeader
	if err := DecodeAndUnmarshal(parts[0], &header); err != nil {
		return nil, nil, fmt.Errorf("header解码失败: %v", err)
	}
	var payload JwtPayload
	if err := DecodeAndUnmarshal(parts[1], &payload); err != nil {
		return nil, nil, fmt.Errorf("payload解码失败: %v", err)
	}
	return &header, &payload, nil
}
