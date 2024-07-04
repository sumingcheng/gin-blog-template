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
	uid, err := strconv.Atoi(ctx.Param("uid")) //获取restful参数
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
		ctx.String(http.StatusBadRequest, "invalid blog id")
		return
	}

	blog := database.GetBlogById(bid)
	if blog == nil {
		ctx.String(http.StatusNotFound, "博客不存在")
		return
	}

	util.LogRus.Debug(blog.Article)
	ctx.HTML(http.StatusOK, "blog.html", gin.H{"title": blog.Title, "article": blog.Article, "bid": blogId, "update_time": blog.UpdateTime.Format("2006-01-02 15:04:05")})
}

type UpdateRequest struct {
	BlogId  int    `form:"bid" binding:"gt=0"`     // 索引值大于0
	Title   string `form:"title" binding:"gt=0"`   // 字符长度大于0
	Article string `form:"article" binding:"gt=0"` // 字符长度大于0
}

// BlogUpdate 更新博客
func BlogUpdate(ctx *gin.Context) {
	//blogId := ctx.PostForm("bid") //获取post form参数
	//title := ctx.PostForm("title")
	//article := ctx.PostForm("article")
	//bid, err := strconv.Atoi(blogId)
	//if err != nil {
	//	ctx.String(http.StatusBadRequest, "invalid blog id")
	//	return
	//}

	var request UpdateRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.String(http.StatusBadRequest, "invalid request")
		return
	}
	bid := request.BlogId
	title := request.Title
	article := request.Article

	blog := database.GetBlogById(bid)
	if blog == nil {
		ctx.String(http.StatusBadRequest, "blog id not exists")
		return
	}

	loginUid := ctx.Value("uid") //从ctx中取得当前登录用户的uid
	if loginUid == nil || loginUid.(int) != blog.UserId {
		ctx.String(http.StatusForbidden, "无权修改")
		return
	}

	err := database.UpdateBlog(&database.Blog{Id: bid, Title: title, Article: article})
	if err != nil {
		util.LogRus.Errorf("update blog %d failed: %s", bid, err)
		ctx.String(http.StatusInternalServerError, "更新失败") //不要把原始的err返回给前端，不利用隐藏错误信息的安全性mysql未
		return
	}
	ctx.String(http.StatusOK, "success")
}

// BlogBelong 检查博客是否属于当前经过认证的用户
func BlogBelong(ctx *gin.Context) {
	bids := ctx.Query("bid")
	token := ctx.Query("token")
	bid, err := strconv.Atoi(bids)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid blog id")
		return
	}
	blog := database.GetBlogById(bid)
	if blog == nil {
		ctx.String(http.StatusBadRequest, "blog id not exists")
		return
	}
	loginUid := middleware.GetUidFromJwt(token)
	if loginUid == blog.UserId {
		ctx.String(http.StatusOK, "true")
	} else {
		ctx.String(http.StatusOK, "false")
	}
}
