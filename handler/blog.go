package handler

import (
	"blog/database"
	"blog/middleware"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BlogListResponse struct {
	Code  int              `json:"code"`
	Msg   string           `json:"msg"`
	Blogs []*database.Blog `json:"blogs"`
}

// BlogList 获取用户的博客列表
func BlogList(ctx *gin.Context) {
	uid, err := strconv.Atoi(ctx.Param("uid"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uid")
		return
	}
	blogs := database.GetBlogByUserId(uid)
	util.LogRus.Debugf("get %d blogs of user %d", len(blogs), uid)
	ctx.JSON(http.StatusOK, BlogListResponse{
		Code:  0,
		Msg:   "success",
		Blogs: blogs,
	})
}

// BlogDetail 获取文章详细的详情
func BlogDetail(ctx *gin.Context) {
	blogId := ctx.Param("bid") //获取restful参数
	bid, err := strconv.Atoi(blogId)
	if err != nil {
		ctx.String(http.StatusBadRequest, "博客id 无效")
		return
	}

	blog := database.GetBlogById(bid)
	if blog == nil {
		ctx.String(http.StatusNotFound, "博客不存在")
		return
	}

	util.LogRus.Debug(blog.Article)
	ctx.JSON(http.StatusOK, BlogListResponse{
		Code:  0,
		Msg:   "success",
		Blogs: []*database.Blog{blog},
	})
}

type UpdateRequest struct {
	BlogId  int    `form:"bid" binding:"gt=0"`               // 索引值大于0
	Title   string `form:"title" binding:"required,min=1"`   // 字符长度大于0
	Article string `form:"article" binding:"required,min=1"` // 字符长度大于0
}

// BlogUpdate 更新博客
func BlogUpdate(ctx *gin.Context) {
	var request UpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		util.LogRus.Errorf("update blog failed: %s", err)
		return
	}

	bid := request.BlogId
	blog := database.GetBlogById(bid)
	if blog == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "blog id not exists",
		})
		util.LogRus.Errorf("blog id %d not exists", bid)
		return
	}

	loginUid, ok := ctx.Value("uid").(int)
	if !ok || loginUid != blog.UserId {
		ctx.JSON(http.StatusForbidden, gin.H{
			"code": 1,
			"msg":  "无权修改",
		})
		util.LogRus.Errorf("user %d attempted to modify blog %d without permission", loginUid, bid)
		return
	}

	err := database.UpdateBlog(&database.Blog{Id: bid, Title: request.Title, Article: request.Article})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "更新失败",
		})
		util.LogRus.Errorf("update blog %d failed: %s", bid, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

type BlogBelongRequest struct {
	Bid   int    `form:"bid" binding:"required"`
	Token string `form:"token" binding:"required"`
}

type BlogBelongResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Belong bool   `json:"belong"`
}

// BlogBelong 检查博客是否属于当前经过认证的用户
func BlogBelong(ctx *gin.Context) {
	var req BlogBelongRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, BlogBelongResponse{
			Code:   1,
			Msg:    err.Error(),
			Belong: false,
		})
		return
	}

	blog := database.GetBlogById(req.Bid)
	if blog == nil {
		ctx.JSON(http.StatusOK, BlogBelongResponse{
			Code:   1,
			Msg:    "blog id not exists",
			Belong: false,
		})
		return
	}

	loginUid := middleware.GetUidFromJwt(req.Token)
	belong := loginUid == blog.UserId

	ctx.JSON(http.StatusOK, BlogBelongResponse{
		Code:   0,
		Msg:    "success",
		Belong: belong,
	})
}
