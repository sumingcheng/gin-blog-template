package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	DefaultHeader = JwtHeader{
		Algo: "HS256",
		Type: "JWT",
	}
)

// JwtHeader 定义JWT的Header部分结构
type JwtHeader struct {
	Algo string `json:"alg"` // 使用的算法, 比如HMAC SHA256
	Type string `json:"typ"` // Token的类型，这里是JWT
}

// JwtPayload 定义JWT的Payload部分结构
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

// GenJWT 生成JWT
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

func main() {
	// 示例：创建和验证JWT
	header := JwtHeader{
		Algo: "HS256",
		Type: "JWT",
	}
	payload := JwtPayload{
		ID:          "123456",
		Issue:       "example.com",
		Audience:    "example.org",
		Subject:     "test",
		IssueAt:     1592000000,
		NotBefore:   1592000000,
		Expiration:  1623532800,
		UserDefined: map[string]string{"role": "admin"},
	}
	secret := "your-256-bit-secret"

	token, err := GenJWT(header, payload, secret)
	if err != nil {
		fmt.Println("生成JWT失败:", err)
	} else {
		fmt.Println("生成的JWT:", token)
		if _, _, err = VerifyJwt(token, secret); err != nil {
			fmt.Println("验证JWT失败:", err)
		} else {
			fmt.Println("JWT验证成功")
		}
	}
}
