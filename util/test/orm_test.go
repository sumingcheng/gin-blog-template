package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type MyStruct struct {
	Id         int    `gorm:"column:id;primaryKey"`          // Tag
	Passwd     string `json:"passwd" gorm:"column:password"` // 密码敏感的字段应该是加密后的，否则从DB里查询出来不能给客户端使用
	Name       string
	FamilyName string    `gorm:"-"` // family_name
	CreateTime time.Time `gorm:"create_time" binding:"required,before_today" time_format:"2006-01-02" time_utc:"8"`
}

// 打印结构体的成员属性信息
func PrintFieldInfo(object any) {
	tp := reflect.TypeOf(object) // 使用reflect.Type获取对象的类型
	fieldNum := tp.NumField()    // 反射获取的字段个数，包括未导出成员
	for i := 0; i < fieldNum; i++ {
		field := tp.Field(i)
		fmt.Printf("%d %s offset %d anonymous %t type %s exported %t gorm tag=%s json tag=%s\n", i,
			field.Name,            // 字段名称
			field.Offset,          // 字段在结构体中地址的偏移量，string类型会占用16个字节
			field.Anonymous,       // 是否为匿名字段
			field.Type,            // 字段类型，reflect.Type类型
			field.IsExported(),    // 字段是否可以导入（即是否以大写字母开头）
			field.Tag.Get("gorm"), // 获取字段标签：标记该字段的gorm
			field.Tag.Get("json"), // 获取字段标签：标记该字段的tag
		)
	}
	fmt.Println()
}

func TestPrintFieldInfoInfo(t *testing.T) {
	PrintFieldInfo(MyStruct{})
}
