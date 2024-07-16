package handler

import (
	"blog/database"
	"blog/middleware"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type LoginResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Uid   int    `json:"uid"`
	Token string `json:"auth_token"`
}

type LoginRequest struct {
	User string `json:"user" binding:"required,min=3"`
	Pass string `json:"pass" binding:"required,len=32"`
}

type TokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"auth_token"`
}

func Login(ctx *gin.Context) {
	// 未登录，登录流程
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, LoginResponse{Code: 1, Msg: util.TranslateErrors(err)})
		return
	}

	// 用户名密码校验
	user := database.GetUserByName(req.User)
	if user == nil {
		ctx.JSON(http.StatusForbidden, LoginResponse{Code: 1, Msg: "用户不存在"})
		return
	}

	if user.PassWd != req.Pass {
		ctx.JSON(http.StatusForbidden, LoginResponse{Code: 1, Msg: "密码不正确"})
		return
	}

	util.LogRus.Infof("user %s(%d) login", req.User, user.Id)

	// 用户名、密码正确，生成 token
	header := util.DefaultHeader
	payload := util.JwtPayload{
		Issue:       "blog",
		IssueAt:     time.Now().Unix(),                                               // 因为每次的IssueAt不同，所以每次生成的token也不同
		Expiration:  time.Now().Add(database.TokenExpire).Add(24 * time.Hour).Unix(), // (7+1)天后过期，需要重新登录，假设24小时内用户肯定要重启浏览器
		UserDefined: map[string]string{middleware.UidInToken: strconv.Itoa(user.Id)}, // 用户自定义字段。如果token里包含敏感信息，需结合https使用
	}
	secret := middleware.KeyConfig.GetString("jwt")

	if authToken, err := util.GenJWT(header, payload, secret); err != nil {
		util.LogRus.Errorf("Failed to generate token: %s", err)
		ctx.JSON(http.StatusOK, LoginResponse{Code: 5, Msg: "Token 生成失败"})
		return
	} else {
		refreshToken, err := util.GetRefreshToken()
		if err != nil {
			util.LogRus.Errorf("Failed to generate refresh token: %s", err)
			ctx.JSON(http.StatusOK, LoginResponse{Code: 6, Msg: "Refresh token 生成失败"})
			return
		}
		database.SetToken(refreshToken, authToken)
		ctx.SetCookie("refresh_token", refreshToken, // 受cookie本身的限制，这里的token不能超过4K
			int(database.TokenExpire.Seconds()), // maxAge，cookie的有效时间，时间单位秒。如果不设置过期时间，默认情况下关闭浏览器后cookie被删除
			"/",                                 // path，cookie存放目录
			"",                                  // cookie从属的域名,不区分协议和端口。 默认本域名
			false,                               // 是否只能通过https访问
			true,                                // 是否允许别人通过js获取自己的cookie，设为false防止XSS攻击
		)
		ctx.JSON(http.StatusOK, LoginResponse{Code: 0, Msg: "success", Uid: user.Id, Token: authToken})
		return
	}
}

func GetAuthToken(ctx *gin.Context) {
	var req TokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, LoginResponse{Code: 1, Msg: "参数错误"})
		return
	}

	authToken := database.GetToken(req.RefreshToken)

	if authToken == "" {
		ctx.JSON(http.StatusForbidden, TokenResponse{Code: 1, Msg: "Refresh token 无效"})
		util.LogRus.Errorf("refresh token %s 无效", req.RefreshToken)
	} else {
		ctx.JSON(http.StatusOK, TokenResponse{Code: 0, Msg: "success", Token: authToken})
	}
}
