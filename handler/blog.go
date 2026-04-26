package handler

import (
	"blog/model"
	"blog/service"
	"blog/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateBlogRequest struct {
	Title   string `json:"title" binding:"required,min=1"`
	Article string `json:"article" binding:"required,min=1"`
}

type UpdateRequest struct {
	BlogId  int    `json:"blogId" binding:"gt=0"`
	Title   string `json:"title" binding:"required,min=1"`
	Article string `json:"article" binding:"required,min=1"`
}

type BlogBelongRequest struct {
	Bid int `json:"bid" binding:"required"`
}

// BlogList 分页列表（支持 ?page=&size=&uid=&keyword=）
func BlogList(ctx *gin.Context) {
	var page model.PageQuery
	if err := ctx.ShouldBindQuery(&page); err != nil {
		model.FailMsg(ctx, model.ErrBadRequest, util.TranslateErrors(err))
		return
	}
	uid, _ := strconv.Atoi(ctx.Query("uid"))
	result := service.GetBlogList(page, uid)
	model.OK(ctx, result)
}

func BlogDetail(ctx *gin.Context) {
	bid, err := strconv.Atoi(ctx.Param("bid"))
	if err != nil {
		model.FailMsg(ctx, model.ErrBadRequest, "无效的博客 ID")
		return
	}
	blog, bizErr := service.GetBlogDetail(bid)
	if bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	model.OK(ctx, blog)
}

func BlogCreate(ctx *gin.Context) {
	var req CreateBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		model.FailMsg(ctx, model.ErrBadRequest, util.TranslateErrors(err))
		return
	}
	uid, _ := ctx.Get("uid")
	blog, bizErr := service.CreateBlog(uid.(int), req.Title, req.Article)
	if bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	model.OK(ctx, blog)
}

func BlogUpdate(ctx *gin.Context) {
	var req UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		model.FailMsg(ctx, model.ErrBadRequest, util.TranslateErrors(err))
		return
	}
	loginUid, _ := ctx.Get("uid")
	bizErr := service.UpdateBlog(loginUid.(int), req.BlogId, req.Title, req.Article)
	if bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	model.OK(ctx, nil)
}

func BlogDelete(ctx *gin.Context) {
	bid, err := strconv.Atoi(ctx.Param("bid"))
	if err != nil {
		model.FailMsg(ctx, model.ErrBadRequest, "无效的博客 ID")
		return
	}
	loginUid, _ := ctx.Get("uid")
	bizErr := service.DeleteBlog(loginUid.(int), bid)
	if bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	model.OK(ctx, nil)
}

func BlogBelong(ctx *gin.Context) {
	var req BlogBelongRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		model.FailMsg(ctx, model.ErrBadRequest, util.TranslateErrors(err))
		return
	}
	token := ctx.Request.Header.Get("auth_token")
	uid := service.GetUidFromToken(token)
	belong, bizErr := service.CheckBlogBelong(uid, req.Bid)
	if bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	model.OK(ctx, gin.H{"belong": belong})
}
