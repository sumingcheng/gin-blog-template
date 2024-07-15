package database

import (
	"blog/util"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	TokenPrefix = "dual_token_"
	TokenExpire = 7 * 24 * time.Hour // 7天
)

func SetToken(refreshToken, authToken string) {
	client := GetRedisClient()
	if err := client.Set(context.Background(), TokenPrefix+refreshToken, authToken, TokenExpire).Err(); err != nil {
		util.LogRus.Errorf("write token pair(%s, %s) to redis failed: %s", refreshToken, authToken, err)
	}
}

func GetToken(refreshToken string) (authToken string) {
	client := GetRedisClient()
	var err error
	if authToken, err = client.Get(context.Background(), TokenPrefix+refreshToken).Result(); err != nil {
		if !errors.Is(redis.Nil, err) {
			util.LogRus.Errorf("get auth token of refresh token %s failed: %s", refreshToken, err)
		}
	}
	return authToken
}
