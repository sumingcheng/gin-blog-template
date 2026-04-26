package database

import (
	"blog/util"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
)

var (
	blogDB     *gorm.DB
	blogDBOnce sync.Once

	blogRedis     *redis.Client
	blogRedisOnce sync.Once
)

func GetBlogDBConnection() *gorm.DB {
	blogDBOnce.Do(func() {
		if blogDB == nil {
			cfg := util.CreateConfig("postgres")
			host := cfg.GetString("host")
			port := cfg.GetInt("port")
			user := cfg.GetString("user")
			pass := cfg.GetString("pass")
			dbname := cfg.GetString("dbname")
			blogDB = createPostgresDB(dbname, host, user, pass, port)
		}
	})
	return blogDB
}

func createPostgresDB(dbname, host, user, pass string, port int) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, port, user, pass, dbname)
	db, err := gorm.Open(
		postgres.Open(dsn), &gorm.Config{
			Logger: ormlog.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				ormlog.Config{
					SlowThreshold: 100 * time.Millisecond,
					LogLevel:      ormlog.Warn,
					Colorful:      true,
				},
			),
			PrepareStmt: true,
		},
	)
	if err != nil {
		util.LogRus.Panicf("连接 postgres 失败: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		util.LogRus.Panicf("获取底层 sql.DB 失败: %v", err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	util.LogRus.Infof("connected to postgres db %s", dbname)

	return db
}

func createRedisClient(
	address, passwd string,
	db int,
) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Username: "default",
		Addr:     address,
		Password: passwd,
		DB:       db,
	})
	// PING redis
	if err := cli.Ping(context.Background()).Err(); err != nil {
		util.LogRus.Panicf("connect to redis %d failed %v", db, err)
	} else {
		util.LogRus.Infof("connect to redis %d", db)
	}
	return cli
}

func GetRedisClient() *redis.Client {
	blogRedisOnce.Do(func() {
		if blogRedis == nil {
			viper := util.CreateConfig("redis")
			addr := viper.GetString("addr")
			pass := viper.GetString("pass")
			db := viper.GetInt("db")
			blogRedis = createRedisClient(addr, pass, db)
		}
	})

	return blogRedis
}
