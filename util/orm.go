package util

import (
	"reflect"
	"strings"
)

func Camel2Snake(name string) string {
	// 创建一个字符串构建器以提高性能。
	var result strings.Builder
	for i, r := range name {
		if r >= 'A' && r <= 'Z' && i > 0 {
			result.WriteByte('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

// GetGormFields 返回基于结构体为 GORM 定义的标签的数据库列名切片。
func GetGormFields(stc any) []string {
	typ := reflect.TypeOf(stc)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil // 提前返回，减少嵌套。
	}

	columns := make([]string, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !field.IsExported() {
			continue // 忽略未导出的字段。
		}

		tag := field.Tag.Get("gorm")
		if tag == "-" {
			continue // 忽略标记为“-”的字段。
		}

		name := Camel2Snake(field.Name)
		if len(tag) > 0 {
			params := strings.Split(tag, ";")
			for _, param := range params {
				if strings.HasPrefix(param, "column:") {
					name = strings.TrimPrefix(param, "column:")
					break
				}
			}
		}
		columns = append(columns, name)
	}
	return columns
}
