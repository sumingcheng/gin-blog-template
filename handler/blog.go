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

type UpdateRequest struct {
	BlogId  int    `form:"bid" binding:"gt=0"`
	Title   string `form:"title" binding:"required,min=1"`
	Article string `form:"article" binding:"required,min=1"`
}

type BlogBelongRequest struct {
	Bid int `form:"bid" binding:"required"`
}

type BlogBelongResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Belong bool   `json:"belong"`
}

func BlogList(ctx *gin.Context) {
	uid, err := strconv.Atoi(ctx.Param("uid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "无效的用户id",
		})
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

func BlogDetail(ctx *gin.Context) {
	blogId := ctx.Param("bid")
	bid, err := strconv.Atoi(blogId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "无效的博客id",
		})
		return
	}

	blog := database.GetBlogById(bid)
	if blog == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 1,
			"msg":  "博客不存在",
		})
		return
	}

	util.LogRus.Debug("get blog detail: ", blog.Article)
	ctx.JSON(http.StatusOK, BlogListResponse{
		Code:  0,
		Msg:   "success",
		Blogs: []*database.Blog{blog},
	})
}

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

// BlogBelong 是否有权修改这篇博客
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

	token := ctx.Request.Header.Get("auth_token")
	loginUid := middleware.GetUidFromJwt(token)
	belong := loginUid == blog.UserId

	ctx.JSON(http.StatusOK, BlogBelongResponse{
		Code:   0,
		Msg:    "success",
		Belong: belong,
	})
}
