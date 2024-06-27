package database

import (
	"blog/util"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	Id         int       `gorm:"column:id;primaryKey"`
	UserId     int       `gorm:"column:user_id"`
	Title      string    `gorm:"column:title"`
	Article    string    `gorm:"column:article"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

// TableName 指定表名为 blog
func (Blog) TableName() string {
	return "blog"
}

var (
	allBlogField = util.GetGormFields(Blog{})
)

// GetBlogById 根据 ID 获取博客
func GetBlogById(id int) *Blog {
	db := GetBlogDBConnection() // 获取数据库连接
	var blog Blog
	if err := db.Select(allBlogField).Where("id = ?", id).First(&blog).Error; err != nil {
		// 如果记录未找到，记录错误，返回 nil
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			util.LogRus.Errorf("get content of blog %d failed: %s", id, err)
		}
		return nil
	}
	return &blog
}

// GetBlogByUserId 根据用户 ID 获取博客列表
func GetBlogByUserId(uid int) []*Blog {
	db := GetBlogDBConnection() // 获取数据库连接
	var blogs []*Blog
	if err := db.Select("id, title").Where("user_id = ?", uid).Find(&blogs).Error; err != nil {
		// 如果查询失败，记录错误，并返回 nil
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			util.LogRus.Errorf("get blogs of user %d failed: %s", uid, err)
		}
		return nil
	}
	return blogs
}

// UpdateBlog 更新博客内容
func UpdateBlog(blog *Blog) error {
	if blog.Id <= 0 {
		// 检查博客 ID 是否有效
		return fmt.Errorf("could not update blog of id %d", blog.Id)
	}
	if len(blog.Article) == 0 || len(blog.Title) == 0 {
		// 检查标题和文章是否为空
		return fmt.Errorf("could not set blog title or article to empty")
	}
	db := GetBlogDBConnection() // 获取数据库连接
	// 更新博客
	return db.Model(blog).Where("id = ?", blog.Id).Updates(map[string]any{"title": blog.Title, "article": blog.Article}).Error
}
