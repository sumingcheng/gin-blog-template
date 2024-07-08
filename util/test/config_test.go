package test

import (
	"blog/util"
	"fmt"
	"testing"
)

// func TestGetOnCurrentPath(t *testing.T) {
// 	str := util.GetOnCurrentPath()
// 	fmt.Println(str)
// }

func TestCreateConfig(t *testing.T) {
	dbViper := util.CreateConfig("key")
	dbViper.WatchConfig() // 监控配置文件变化
	if !dbViper.IsSet("apiVersion") {
		t.Fail()
	}
	fmt.Println(dbViper.GetString("apiVersion"))
}
