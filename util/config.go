package util

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	ProjectRootPath = getProjectRootPath()
)

func getProjectRootPath() string {
	// 从环境变量获取配置路径，如果未设置，则使用默认路径
	if rootPath := os.Getenv("CONFIG_PATH"); rootPath != "" {
		return rootPath
	}
	return "./"
}

// CreateConfig 用于创建并读取配置文件
func CreateConfig(file string) *viper.Viper {
	config := viper.New()
	configPath := ProjectRootPath + "config/"
	config.AddConfigPath(configPath) // 设置配置文件路径
	config.SetConfigName(file)       // 设置文件名
	config.SetConfigType("yaml")     // 设置文件类型

	configFile := configPath + file + ".yaml"

	// 尝试读取配置文件
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("找不到配置文件：%s", configFile))
		} else {
			panic(fmt.Errorf("解析配置文件出错：%s; 错误：%s", configFile, err))
		}
	}
	return config
}
