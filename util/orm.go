package util

import (
	"reflect"
	"strings"
)

// Camel2Snake 将驼峰式命名转换为下划线命名（蛇形命名）。
func Camel2Snake(name string) string {
	return strings.ToLower(name)
}

// GetGormFields 返回基于结构体为 GORM 定义的标签的数据库列名切片。
func GetGormFields(stc any) []string {
	typ := reflect.TypeOf(stc)
	// 检查提供的接口是否为指针，如果是，获取它指向的类型。
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// 仅当提供的类型是结构体时处理。
	if typ.Kind() == reflect.Struct {
		columns := make([]string, 0, typ.NumField()) // 初始化一个切片来存储列名。
		for i := 0; i < typ.NumField(); i++ {
			fieldType := typ.Field(i) // 通过索引获取结构体字段。

			// 检查字段是否为导出（公共）字段；未导出的字段将被忽略。
			if fieldType.IsExported() {
				// 检查字段在 GORM 中是否应被忽略（使用“-”标签）。
				if fieldType.Tag.Get("gorm") == "-" {
					continue
				}
				// 将字段名称从驼峰式命名转换为蛇形命名作为默认列名。
				name := Camel2Snake(fieldType.Name)
				// 检查是否存在 GORM 标签且不为空。
				if len(fieldType.Tag.Get("gorm")) > 0 {
					content := fieldType.Tag.Get("gorm")
					// 检查标签是否明确指定了列名。
					if strings.HasPrefix(content, "column:") {
						content = content[7:]              // 去掉前缀“column:”以获取实际的列名。
						pos := strings.Index(content, ";") // 查找分号的位置（如果有）。
						if pos > 0 {
							name = content[:pos] // 使用指定到分号为止的列名。
						} else if pos < 0 {
							name = content // 如果没有分号，使用整个内容作为列名。
						}
					}
				}
				columns = append(columns, name) // 将确定的列名添加到切片中。
			}
		}
		return columns // 返回包含列名的切片。
	} else {
		return nil // 如果提供的接口不是结构体，返回 nil。
	}
}
