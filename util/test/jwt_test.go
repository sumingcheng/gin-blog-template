package test

import (
	"blog/util"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"time"
)

// TestBase64 测试Base64编码和解码
func TestBase64(t *testing.T) {
	text := "你好，世界"                                           // 原始字符串
	cipher := base64.StdEncoding.EncodeToString([]byte(text)) // Base64编码
	fmt.Println(cipher)

	bs, _ := base64.StdEncoding.DecodeString(cipher) // Base64解码
	if string(bs) != text {
		t.Fail() // 如果解码后的字符串与原始字符串不一致，则测试失败
	}
}

// TestJWT 测试JWT的生成和验证
func TestJWT(t *testing.T) {
	secret := "123456" // 密钥
	header := util.DefaultHeader
	payload := util.JwtPayload{
		ID:          "rj4t49tu49",
		Issue:       "颁发者",
		Audience:    "目标受众",
		Subject:     "主题",
		IssueAt:     time.Now().Unix(),
		Expiration:  time.Now().Add(2 * time.Hour).Unix(),
		UserDefined: map[string]string{"name": strings.Repeat("素明诚", 100)}, // 自定义字段
	}

	if token, err := util.GenJWT(header, payload, secret); err != nil {
		fmt.Printf("生成json views token失败: %v", err)
	} else {
		fmt.Println(token)
		if _, p, err := util.VerifyJwt(token, secret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("JWT验证通过。欢迎 %s !\n", p.UserDefined["name"])
		}
	}
}
