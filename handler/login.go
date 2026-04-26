package handler

import (
	"blog/database"
	"blog/model"
	"blog/service"
	"blog/util"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	User string `json:"user" binding:"required,min=3"`
	Pass string `json:"pass" binding:"required,min=6"`
}

type RegisterRequest struct {
	User string `json:"user" binding:"required,min=3,max=20"`
	Pass string `json:"pass" binding:"required,min=6"`
}

type ChangePasswordRequest struct {
	OldPass string `json:"oldPass" binding:"required"`
	NewPass string `json:"newPass" binding:"required,min=6"`
}

func Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		model.FailMsg(ctx, model.ErrBadRequest, util.TranslateErrors(err))
		return
	}
	user, bizErr := service.Register(req.User, req.Pass)
	if bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	model.OK(ctx, gin.H{"uid": user.Id, "name": user.Name})
}

func Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		model.FailMsg(ctx, model.ErrBadRequest, util.TranslateErrors(err))
		return
	}
	result, bizErr := service.Login(req.User, req.Pass)
	if bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	ctx.SetCookie("refresh_token", result.RefreshToken,
		int(database.TokenExpire.Seconds()), "/", "", false, true,
	)
	model.OK(ctx, gin.H{"uid": result.Uid, "auth_token": result.AuthToken})
}

func GetAuthToken(ctx *gin.Context) {
	refreshToken, _ := ctx.Cookie("refresh_token")
	authToken, bizErr := service.RefreshAuthToken(refreshToken)
	if bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	model.OK(ctx, gin.H{"auth_token": authToken})
}

func Logout(ctx *gin.Context) {
	refreshToken, _ := ctx.Cookie("refresh_token")
	if bizErr := service.Logout(refreshToken); bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	ctx.SetCookie("refresh_token", "", -1, "/", "", false, true)
	model.OK(ctx, nil)
}

func ChangePassword(ctx *gin.Context) {
	var req ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		model.FailMsg(ctx, model.ErrBadRequest, util.TranslateErrors(err))
		return
	}
	uid, _ := ctx.Get("uid")
	if bizErr := service.ChangePassword(uid.(int), req.OldPass, req.NewPass); bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	model.OK(ctx, nil)
}

func GetProfile(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	user, bizErr := service.GetProfile(uid.(int))
	if bizErr != nil {
		model.Fail(ctx, bizErr)
		return
	}
	model.OK(ctx, user)
}
