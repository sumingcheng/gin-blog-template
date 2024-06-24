package util

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs" // 导入日志文件轮转库
	"github.com/sirupsen/logrus"                        // 导入Logrus日志库
	"strings"
	"time"
)

var (
	LogRus *logrus.Logger // 定义全局的Logger实例
)

func InitLog(configFile string) {
	viper := CreateConfig(configFile) // 创建配置管理器
	LogRus = logrus.New()             // 实例化一个新的Logger

	// 根据配置文件设置日志级别
	switch strings.ToLower(viper.GetString("level")) {
	case "debug":
		LogRus.SetLevel(logrus.DebugLevel)
	case "info":
		LogRus.SetLevel(logrus.InfoLevel)
	case "warn":
		LogRus.SetLevel(logrus.WarnLevel)
	case "error":
		LogRus.SetLevel(logrus.ErrorLevel)
	case "panic":
		LogRus.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("invalid log level %s", viper.GetString("level"))) // 配置文件中日志级别无效时抛出异常
	}

	// 设置日志格式
	LogRus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 时间戳格式
	})

	// 设置日志文件路径和轮转策略
	logFile := ProjectRootPath + viper.GetString("file") // 日志文件路径
	fout, err := rotatelogs.New(
		logFile+".%Y%m%d%H",                      // 日志文件名称格式
		rotatelogs.WithLinkName(logFile),         // 设置最新日志的链接
		rotatelogs.WithRotationTime(1*time.Hour), // 日志轮转时间为1小时
		rotatelogs.WithMaxAge(7*24*time.Hour),    // 日志文件最大保存时间为7天
	)

	if err != nil {
		panic(err) // 创建日志文件失败时抛出异常
	}

	LogRus.SetOutput(fout)        // 设置日志输出到文件
	LogRus.SetReportCaller(false) // 日志中包含调用者信息
}
