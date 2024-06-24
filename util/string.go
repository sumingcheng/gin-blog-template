package util

import (
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	l := len(letterRunes)

	for i := range b {
		b[i] = letterRunes[rand.Intn(l)]
	}

	return string(b)
}
