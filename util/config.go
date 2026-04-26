package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func getProjectRootPath() string {
	// 生产环境: 配置文件与可执行文件在同一目录
	if os.Getenv("APP_ENV") == "production" {
		return "/config"
	}

	// 开发环境: 返回代码文件所在的目录上一级的 'config' 目录
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "../config")
}

// CreateConfig 用于创建并读取配置文件
func CreateConfig(file string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath(getProjectRootPath())
	config.SetConfigName(file) // 配置文件名
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件 %s.yaml 失败: %w", file, err))
	}
	return config
}
