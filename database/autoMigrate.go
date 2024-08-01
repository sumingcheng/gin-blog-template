package database

import (
	"gorm.io/gorm"
	"log"
)

func AutoMigrate() {
	db := GetBlogDBConnection()
	err := db.AutoMigrate(&Blog{}, &User{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := createInitialUsers(db); err != nil {
		log.Fatalf("Failed to create initial users: %v", err)
	}

	if err := createInitialBlogs(db); err != nil {
		log.Fatalf("Failed to create initial blogs: %v", err)
	}
}

func createInitialUsers(db *gorm.DB) error {
	var count int64
	db.Model(&User{}).Count(&count)
	if count == 0 {
		users := []User{
			{Name: "admin", PassWd: "e10adc3949ba59abbe56e057f20f883e"},
		}
		return db.Create(&users).Error
	}
	return nil
}

func createInitialBlogs(db *gorm.DB) error {
	var count int64
	db.Model(&Blog{}).Count(&count)
	if count == 0 {
		blogs := []Blog{
			{UserId: 1, Title: "博客标题1", Article: "博客内容1"},
			{UserId: 1, Title: "博客标题2", Article: "博客内容2"},
			{UserId: 1, Title: "博客标题3", Article: "博客内容3"},
			{UserId: 1, Title: "博客标题4", Article: "博客内容4"},
			{UserId: 1, Title: "博客标题5", Article: "博客内容5"},
		}
		return db.Create(&blogs).Error
	}
	return nil
}
