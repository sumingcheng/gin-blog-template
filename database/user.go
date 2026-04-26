package database

import (
	"blog/util"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Id        int    `gorm:"column:id;primaryKey" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	PassWd    string `gorm:"column:password" json:"-"`
	CreatedAt int64  `gorm:"column:created_at" json:"createdAt"`
	UpdateAt  int64  `gorm:"column:update_at" json:"updateAt"`
	DeleteAt  *int64 `gorm:"column:delete_at" json:"-"`
}

func (User) TableName() string {
	return "user"
}

var allUserField = util.GetGormFields(User{})

// GetUserByName 根据用户名检索用户（排除已删除）
func GetUserByName(name string) *User {
	db := GetBlogDBConnection()
	var user User
	if err := db.Select(allUserField).Where("name = ? AND delete_at IS NULL", name).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			util.LogRus.Errorf("get user %s failed: %s", name, err)
		}
		return nil
	}
	return &user
}

// GetUserById 根据 ID 检索用户
func GetUserById(id int) *User {
	db := GetBlogDBConnection()
	var user User
	if err := db.Select(allUserField).Where("id = ? AND delete_at IS NULL", id).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			util.LogRus.Errorf("get user id %d failed: %s", id, err)
		}
		return nil
	}
	return &user
}

// CreateUser 创建用户
func CreateUser(name, hashedPass string) (*User, error) {
	db := GetBlogDBConnection()
	user := User{Name: name, PassWd: hashedPass}
	if err := db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("create user %s failed: %w", name, err)
	}
	util.LogRus.Infof("create user id %d", user.Id)
	return &user, nil
}

// UpdateUserPassword 更新密码
func UpdateUserPassword(id int, hashedPass string) error {
	db := GetBlogDBConnection()
	result := db.Model(&User{}).Where("id = ? AND delete_at IS NULL", id).Update("password", hashedPass)
	if result.Error != nil {
		return fmt.Errorf("update user %d password failed: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user %d not found", id)
	}
	return nil
}

// DeleteUser 软删除
func DeleteUser(name string) error {
	db := GetBlogDBConnection()
	if err := db.Where("name = ?", name).Delete(User{}).Error; err != nil {
		return fmt.Errorf("delete user %s failed: %w", name, err)
	}
	return nil
}
