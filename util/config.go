package util

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var (
	ProjectRootPath = GetOnCurrentPath() + "/../"
)

func GetOnCurrentPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
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
