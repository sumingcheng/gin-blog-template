package util

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword bcrypt 哈希，cost=10
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword 比对明文与 bcrypt 哈希
func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
