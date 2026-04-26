package database

import (
	"blog/util"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	TokenPrefix = "dual_token_"
	TokenExpire = 7 * 24 * time.Hour // 7天
)

func SetToken(refreshToken, authToken string) error {
	client := GetRedisClient()
	if err := client.Set(context.Background(), TokenPrefix+refreshToken, authToken, TokenExpire).Err(); err != nil {
		return fmt.Errorf("write token pair to redis failed: %w", err)
	}
	return nil
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

func RmToken(refreshToken string) error {
	client := GetRedisClient()
	_, err := client.Del(context.Background(), TokenPrefix+refreshToken).Result()
	if err != nil {
		util.LogRus.Errorf("Failed to delete refresh token %s: %s", refreshToken, err)
		return err
	}
	return nil
}

func VerifyRefreshToken(refreshToken string) (authToken string, valid bool) {
	client := GetRedisClient()
	authToken, err := client.Get(context.Background(), TokenPrefix+refreshToken).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			util.LogRus.Warnf("refresh token 已过期或不存在: %s", refreshToken)
		} else {
			util.LogRus.Errorf("查询 refresh token %s 失败: %s", refreshToken, err)
		}
		return "", false
	}

	return authToken, true
}
