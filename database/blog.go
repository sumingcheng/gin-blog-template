package database

import (
	"blog/util"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	Id        int    `gorm:"column:id;primaryKey" json:"id"`
	UserId    int    `gorm:"column:user_id" json:"userId"`
	Title     string `gorm:"column:title" json:"title"`
	Article   string `gorm:"column:article" json:"article"`
	CreatedAt int64  `gorm:"column:created_at" json:"createdAt"`
	UpdateAt  int64  `gorm:"column:update_at" json:"updateAt"`
	DeleteAt  *int64 `gorm:"column:delete_at" json:"-"`
}

func (Blog) TableName() string {
	return "blog"
}

var allBlogField = util.GetGormFields(Blog{})

// notDeleted 软删除过滤条件
func notDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("delete_at IS NULL")
}

// GetBlogById 根据 ID 获取博客
func GetBlogById(id int) *Blog {
	db := GetBlogDBConnection()
	var blog Blog
	if err := db.Scopes(notDeleted).Select(allBlogField).Where("id = ?", id).First(&blog).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			util.LogRus.Errorf("get blog %d failed: %s", id, err)
		}
		return nil
	}
	return &blog
}

// GetBlogByUserId 根据用户 ID 获取博客列表
func GetBlogByUserId(uid int) []*Blog {
	db := GetBlogDBConnection()
	var blogs []*Blog
	if err := db.Scopes(notDeleted).Select(allBlogField).Where("user_id = ?", uid).Order("id DESC").Find(&blogs).Error; err != nil {
		util.LogRus.Errorf("get blogs of user %d failed: %s", uid, err)
		return nil
	}
	return blogs
}

// GetBlogList 分页查询博客列表
func GetBlogList(offset, limit, uid int, keyword string) ([]*Blog, int64) {
	db := GetBlogDBConnection()
	query := db.Scopes(notDeleted).Model(&Blog{})

	if uid > 0 {
		query = query.Where("user_id = ?", uid)
	}
	if keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("title ILIKE ? OR article ILIKE ?", like, like)
	}

	var total int64
	query.Count(&total)

	var blogs []*Blog
	query.Select(allBlogField).Order("id DESC").Offset(offset).Limit(limit).Find(&blogs)
	return blogs, total
}

// CreateBlog 创建博客
func CreateBlog(userId int, title, article string) (*Blog, error) {
	db := GetBlogDBConnection()
	blog := Blog{UserId: userId, Title: title, Article: article}
	if err := db.Create(&blog).Error; err != nil {
		return nil, fmt.Errorf("create blog failed: %w", err)
	}
	return &blog, nil
}

// UpdateBlog 更新博客
func UpdateBlog(blog *Blog) error {
	if blog.Id <= 0 {
		return fmt.Errorf("invalid blog id %d", blog.Id)
	}
	db := GetBlogDBConnection()
	now := time.Now().Unix()
	return db.Model(&Blog{}).Where("id = ?", blog.Id).Scopes(notDeleted).
		Updates(map[string]any{"title": blog.Title, "article": blog.Article, "update_at": now}).Error
}

// DeleteBlog 软删除博客
func DeleteBlog(id int) error {
	db := GetBlogDBConnection()
	now := time.Now().Unix()
	result := db.Model(&Blog{}).Where("id = ?", id).Scopes(notDeleted).Update("delete_at", now)
	if result.Error != nil {
		return fmt.Errorf("delete blog %d failed: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("blog %d not found", id)
	}
	return nil
}
