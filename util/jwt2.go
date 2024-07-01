package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var key = []byte("ssss")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int `json:"UserID"`
	jwt.RegisteredClaims     // 内嵌标准的声明
}

// GetAccessToken  生成AccessToken,一般生命周期短,30分钟
func GetAccessToken(UserId int) (string, error) {
	claims := CustomClaims{
		UserID: UserId, // 自定义字段
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Sakura_Blog_AccessToken", // 签发人
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)), // 定义过期时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(key)
}

// GetRefreshToken 生成RefreshToken,生命周期长24小时
// RefreshToken可以不包含用户信息
func GetRefreshToken() (string, error) {
	// RefreshToken的 claims 可以不包含用户信息
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

func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token,通过结构体解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
