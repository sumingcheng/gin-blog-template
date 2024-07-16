package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int `json:"UserID"`
	jwt.RegisteredClaims     // 内嵌标准的声明
}

var (
	DefaultHeader = JwtHeader{
		Algo: "HS256",
		Type: "JWT",
	}
)

var key = []byte("ssss")

func GetRefreshToken() (string, error) {
	// RefreshToken的 claims 不包含用户信息
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Sakura_Blog_RefreshToken", // 签发人
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 定义过期时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(key)
}

func GenJWT(
	header JwtHeader,
	payload JwtPayload,
	secret string,
) (string, error) {
	var part1, part2, signature string

	// 将header序列化为JSON，然后进行Base64编码
	if bs1, err := json.Marshal(header); err != nil {
		return "", err
	} else {
		part1 = base64.RawURLEncoding.EncodeToString(bs1)
	}

	// 将payload序列化为JSON，然后进行Base64编码
	if bs2, err := json.Marshal(payload); err != nil {
		return "", err
	} else {
		part2 = base64.RawURLEncoding.EncodeToString(bs2)
	}

	// 使用HMAC SHA256算法签名
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(part1 + "." + part2))
	signature = base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	// 拼接成完整的JWT
	return part1 + "." + part2 + "." + signature, nil
}

// VerifyJwt 验证JWT
func VerifyJwt(token, secret string) (*JwtHeader, *JwtPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, fmt.Errorf("token格式不正确")
	}

	// 验证签名
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(parts[0] + "." + parts[1]))
	expectedSignature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	if expectedSignature != parts[2] {
		return nil, nil, fmt.Errorf("签名不匹配")
	}

	// 解码Header和Payload
	var header JwtHeader
	var payload JwtPayload
	if part1, err := base64.RawURLEncoding.DecodeString(parts[0]); err != nil {
		return nil, nil, fmt.Errorf("header Base64解码失败")
	} else if err = json.Unmarshal(part1, &header); err != nil {
		return nil, nil, fmt.Errorf("header JSON解码失败")
	}

	if part2, err := base64.RawURLEncoding.DecodeString(parts[1]); err != nil {
		return nil, nil, fmt.Errorf("payload Base64解码失败")
	} else if err = json.Unmarshal(part2, &payload); err != nil {
		return nil, nil, fmt.Errorf("payload JSON解码失败")
	}

	return &header, &payload, nil
}
