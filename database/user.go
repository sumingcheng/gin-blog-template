package database

import (
	"blog/util"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	Id     int `gorm:"column:id;primaryKey"`
	Name   string
	PassWd string `gorm:"column:password"`
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
		if !errors.Is(err, gorm.ErrRecordNotFound) { // 如果是用户名不存在，不需要打错误日志
			util.LogRus.Errorf("get password of user %s failed: %s", name, err) // 系统性异常，才打错误日志
		}
		return nil
	}
	return &user
}

// CreateUser 创建一个用户
func CreateUser(name, pass string) {
	db := GetBlogDBConnection()                    // 获取数据库连接
	pass = util.Md5(pass)                          // 使用MD5加密密码
	user := User{Name: name, PassWd: pass}         // 创建用户实例
	if err := db.Create(&user).Error; err != nil { // 创建用户
		// 注意 Create(&user) 这里传递结构体指针是因为 GORM 背后要修改这个user的主键ID
		// 如果调用成功了, 你需要通过 Create 函数去修改 user 的 Id 字段 所以需要传递指针
		util.LogRus.Errorf("create user %s failed: %s", name, err) // 创建失败，记录错误
	} else {
		util.LogRus.Infof("create user id %d", user.Id) // 创建成功，记录用户ID
	}
}

// DeleteUser 删除一个用户
func DeleteUser(name string) {
	db := GetBlogDBConnection()                                           // 获取数据库连接
	if err := db.Where("name=?", name).Delete(User{}).Error; err != nil { // 删除指定用户名的用户
		// 注意 Delete(User{}) 这里传递结构体的目的是为了找到表明
		util.LogRus.Errorf("delete user %s failed: %s", name, err) // 删除失败，记录错误
	}
}
