package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
