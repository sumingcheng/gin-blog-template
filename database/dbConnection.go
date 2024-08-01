package database

import (
	"blog/util"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var (
	blogMysql     *gorm.DB
	blogMysqlOnce sync.Once

	blogRedis     *redis.Client
	blogRedisOnce sync.Once
)

func GetBlogDBConnection() *gorm.DB {
	// 并发安全的单例模式
	blogMysqlOnce.Do(func() {
		if blogMysql == nil {
			dbName := "blog"
			viper := util.CreateConfig("mysql")
			host := viper.GetString(dbName + ".host")
			port := viper.GetInt(dbName + ".port")
			user := viper.GetString(dbName + ".user")
			pass := viper.GetString(dbName + ".pass")
			blogMysql = createMysqlDB(dbName, host, user, pass, port)
		}
	})
	return blogMysql
}

func createMysqlDB(
	dbname, host, user, pass string,
	port int,
) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(
		mysql.Open(dsn), &gorm.Config{
			Logger: ormlog.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				ormlog.Config{
					SlowThreshold: 100 * time.Millisecond, // 设置 SQL 阈值
					LogLevel:      ormlog.Silent,          // Silent表示不输出日志
					Colorful:      true,                   // 彩色日志打印
				},
			),
			PrepareStmt: true, // 启用PrepareStmt, SQL预编译，提高查询效率
		},
	)

	if err != nil {
		util.LogRus.Panicf("connect to mysql use dsn %v failed: %v", dsn, err)
	}

	// 设置数据库连接池参数，提高并发性能
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100) // 设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  // 设置连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超出的连接会被连接池关闭。
	util.LogRus.Infof("connect to mysql db %v", dbname)

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
