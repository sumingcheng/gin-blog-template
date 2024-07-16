package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

func EncodeBase64(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func DecodeBase64(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}

// GenerateSignature 生成 HMAC SHA256 签名。
func GenerateSignature(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return EncodeBase64(h.Sum(nil))
}

// CheckSignature 检查提供的签名是否有效。
func CheckSignature(part1, part2, receivedSig, secret string) bool {
	data := part1 + "." + part2
	expectedSig := GenerateSignature(data, secret)
	return expectedSig == receivedSig
}

func MarshalAndEncode(obj interface{}) (string, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return EncodeBase64(jsonBytes), nil
}

func DecodeAndUnmarshal(encodedStr string, obj interface{}) error {
	decodedBytes, err := DecodeBase64(encodedStr)
	if err != nil {
		return err
	}
	return json.Unmarshal(decodedBytes, obj)
}
