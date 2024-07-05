package util

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

const (
	KeyPrefix = "auth_cookie_"
)

var (
	blogRedis     *redis.Client
	blogRedisOnce sync.Once
	ctx           = context.Background()
)

func createRedisClient(
	address, passwd string,
	db int,
) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: passwd,
		DB:       db,
	})
	if err := cli.Ping(ctx).Err(); err != nil {
		panic(fmt.Errorf("connect to redis %d failed %v", db, err))
	} else {
		fmt.Printf("connect to redis %d\n", db)
	}
	return cli
}

func GetRedisClient() *redis.Client {
	blogRedisOnce.Do(func() {
		if blogRedis == nil {
			blogRedis = createRedisClient("127.0.0.1:6379", "", 0)
		}
	})
	return blogRedis
}

func SetCookieAuth(cookieValue, uid string) {
	client := GetRedisClient()
	if err := client.Set(ctx, KeyPrefix+cookieValue, uid, time.Hour*24*30).Err(); err != nil { // 30天之后过期
		fmt.Printf("write pair(%s, %s) to redis failed: %s\n", cookieValue, uid, err)
	}
}

func GetCookieAuth(cookieValue string) (uid string) {
	client := GetRedisClient()
	var err error
	if uid, err = client.Get(ctx, KeyPrefix+cookieValue).Result(); err != nil {
		fmt.Printf("get auth info %s failed: %s\n", cookieValue, err)
	}
	return
}
