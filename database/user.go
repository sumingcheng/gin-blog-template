package database

import (
	"blog/util"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID     int `gorm:"column:id;primaryKey"`
	Name   string
	PassWd string `gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}

//var (
//	allUserField = util.GetGormFields(User{})
//)

// GetUserByName 根据用户名检索用户
func GetUserByName(name string) *User {
	db := GetBlogDBConnection()
	var user User

	if err := db.Select([]string{"id", "name", "password"}).Where("name=?", name).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) { // 如果是用户名不存在，不需要打错误日志
			util.LogRus.Errorf("get password of user %s failed: %s", name, err) // 系统性异常，才打错误日志
		}
		return nil
	}
	return &user
}
