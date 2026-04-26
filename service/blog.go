package service

import (
	"blog/database"
	"blog/model"
	"blog/util"
)

// GetBlogsByUserId 获取用户的博客列表
func GetBlogsByUserId(uid int) []*database.Blog {
	blogs := database.GetBlogByUserId(uid)
	util.LogRus.Debugf("get %d blogs of user %d", len(blogs), uid)
	return blogs
}

// GetBlogDetail 获取博客详情
func GetBlogDetail(bid int) (*database.Blog, *model.BizError) {
	blog := database.GetBlogById(bid)
	if blog == nil {
		return nil, model.ErrBlogNotFound
	}
	return blog, nil
}

// GetBlogList 分页查询博客
func GetBlogList(page model.PageQuery, uid int) *model.PageResult {
	page.Normalize()
	blogs, total := database.GetBlogList(page.Offset(), page.Size, uid, page.Keyword)
	return &model.PageResult{
		List:  blogs,
		Total: total,
		Page:  page.Page,
		Size:  page.Size,
	}
}

// CreateBlog 创建博客
func CreateBlog(uid int, title, article string) (*database.Blog, *model.BizError) {
	blog, err := database.CreateBlog(uid, title, article)
	if err != nil {
		util.LogRus.Errorf("create blog failed: %s", err)
		return nil, model.ErrInternalServer
	}
	return blog, nil
}

// UpdateBlog 更新博客（含权限校验）
func UpdateBlog(loginUid, blogId int, title, article string) *model.BizError {
	blog := database.GetBlogById(blogId)
	if blog == nil {
		return model.ErrBlogNotFound
	}
	if loginUid != blog.UserId {
		util.LogRus.Errorf("user %d attempted to modify blog %d without permission", loginUid, blogId)
		return model.ErrForbidden
	}
	if err := database.UpdateBlog(&database.Blog{Id: blogId, Title: title, Article: article}); err != nil {
		util.LogRus.Errorf("update blog %d failed: %s", blogId, err)
		return model.ErrBlogUpdate
	}
	return nil
}

// DeleteBlog 删除博客（软删除，含权限校验）
func DeleteBlog(loginUid, blogId int) *model.BizError {
	blog := database.GetBlogById(blogId)
	if blog == nil {
		return model.ErrBlogNotFound
	}
	if loginUid != blog.UserId {
		return model.ErrForbidden
	}
	if err := database.DeleteBlog(blogId); err != nil {
		util.LogRus.Errorf("delete blog %d failed: %s", blogId, err)
		return model.ErrInternalServer
	}
	return nil
}

// CheckBlogBelong 检查博客归属
func CheckBlogBelong(uid, bid int) (bool, *model.BizError) {
	blog := database.GetBlogById(bid)
	if blog == nil {
		return false, model.ErrBlogNotFound
	}
	return uid == blog.UserId, nil
}
