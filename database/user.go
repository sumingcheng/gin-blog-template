package database

import (
	"blog/util"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `gorm:"column:id;primaryKey;comment:'ID，主键，自增'"`
	Name     string `gorm:"column:name;size:20;not null;unique;comment:'用户名'"`
	PassWd   string `gorm:"column:password;size:32;not null;comment:'密码MD5'"`
	UpdateAt int64  `gorm:"column:update_at;comment:'最后修改时间 UTC'"`
	DeleteAt int64  `gorm:"column:delete_at;comment:'删除时间 UTC'"`
}

func (User) TableName() string {
	return "user"
}

var (
	allUserField = util.GetGormFields(User{})
)

// GetUserByName 根据用户名检索用户
func GetUserByName(name string) *User {
	db := GetBlogDBConnection()
	var user User

	if err := db.Select(allUserField).Where("name=?", name).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			util.LogRus.Errorf("get password of user %s failed: %s", name, err)
		}
		return nil
	}
	return &user
}

func CreateUser(name, pass string) {
	db := GetBlogDBConnection()
	pass = util.Md5(pass)
	user := User{Name: name, PassWd: pass}
	// &user 传递指针，以便在创建后获取 ID
	if err := db.Create(&user).Error; err != nil {
		util.LogRus.Errorf("create user %s failed: %s", name, err)
	} else {
		util.LogRus.Infof("create user id %d", user.Id)
	}
}

func DeleteUser(name string) {
	db := GetBlogDBConnection()
	// User{} 主要用于指明操作应该影响哪个表
	if err := db.Where("name=?", name).Delete(User{}).Error; err != nil {
		util.LogRus.Errorf("delete user %s failed: %s", name, err)
	}
}
