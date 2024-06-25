package database

import (
	"blog/util"
	"fmt"
	"github.com/go-redis/redis"
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
	dbLog         ormlog.Interface

	blog_redis      *redis.Client
	blog_redis_once sync.Once
)

func init() {
	dbLog = ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		ormlog.Config{
			SlowThreshold: 100 * time.Millisecond, // 设置 SQL 阈值
			LogLevel:      ormlog.Info,            // Log level, Silent表示不输出日志
			Colorful:      true,                   // 彩色日志打印
		},
	)
}

func createMysqlDB(
	dbname, host, user, pass string,
	port int,
) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname) // mb4表情符号
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dbLog, PrepareStmt: true}) // 启用PrepareStmt, SQL预编译，提高查询效率
	if err != nil {
		util.LogRus.Panicf("connect to mysql use dsn %s failed: %s", dsn, err) // panic() os.Exit(2)
	}
	// 设置数据库连接池参数，提高并发性能
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100) // 设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  // 设置连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超出的连接会被连接池关闭。
	util.LogRus.Infof("connect to mysql db %s", dbname)
	return db
}

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
