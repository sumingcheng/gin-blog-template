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
	User string `json:"user" binding:"required,min=1"`  // 用户名必须非空
	Pass string `json:"pass" binding:"required,len=32"` // 密码必须非空且长度为32
}

func Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, LoginResponse{Code: 1, Msg: "Invalid request"})
		return
	}
	// 是否已登录
	CheckRefreshToken(ctx)

	// 用户名密码校验
	user := database.GetUserByName(req.User)
	if user == nil {
		ctx.JSON(http.StatusForbidden, LoginResponse{Code: 3, Msg: "User does not exist"})
		return
	}

	if user.PassWd != req.Pass {
		ctx.JSON(http.StatusForbidden, LoginResponse{Code: 4, Msg: "Password incorrect"})
		return
	}

	util.LogRus.Infof("user %s(%d) login", req.User, user.Id)

	// 用户名、密码正确，向客户端返回一个 token
	header := util.DefaultHeader // 默认JWT头部信息，定义在 util 包中
	payload := util.JwtPayload{  // payload以明文形式编码在token中，server用自己的密钥可以校验该信息是否被篡改过
		Issue:       "blog",
		IssueAt:     time.Now().Unix(),                                               //因为每次的IssueAt不同，所以每次生成的token也不同
		Expiration:  time.Now().Add(database.TokenExpire).Add(24 * time.Hour).Unix(), //(7+1)天后过期，需要重新登录，假设24小时内用户肯定要重启浏览器
		UserDefined: map[string]string{middleware.UidInToken: strconv.Itoa(user.Id)}, //用户自定义字段。如果token里包含敏感信息，请结合https使用
	}

	if accessToken, err := util.GenJWT(header, payload, middleware.KeyConfig.GetString("jwt")); err != nil {
		util.LogRus.Errorf("Failed to generate token: %s", err)
		ctx.JSON(http.StatusInternalServerError, LoginResponse{Code: 5, Msg: "Token generation failed"})
		return
	} else {
		refreshToken, _ := util.GetRefreshToken()    //生成RefreshToken
		database.SetToken(refreshToken, accessToken) //把<refreshToken, authToken>写入redis
		ctx.SetCookie("refresh_token", refreshToken, //注意：受cookie本身的限制，这里的token不能超过4K
			int(database.TokenExpire.Seconds()), //maxAge，cookie的有效时间，时间单位秒。如果不设置过期时间，默认情况下关闭浏览器后cookie被删除
			"/",                                 //path，cookie存放目录
			"",                                  //cookie从属的域名,不区分协议和端口。如果不指定domain则默认为本host(如b.a.com)，如果指定的domain是一级域名(如a.com)，则二级域名(b.a.com)下也可以访问
			false,                               //是否只能通过https访问
			true,                                //是否允许别人通过js获取自己的cookie，设为false防止XSS攻击
		)
		ctx.JSON(http.StatusOK, LoginResponse{Code: 0, Msg: "success", Uid: user.Id, Token: accessToken})
		return
	}
}

type TokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"auth_token"`
}

func GetAuthToken(ctx *gin.Context) {
	var req TokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, LoginResponse{Code: 1, Msg: "参数错误"})
		return
	}
	authToken := database.GetToken(req.RefreshToken)
	if authToken == "" {
		ctx.JSON(http.StatusForbidden, TokenResponse{Code: 1, Msg: "Refresh token is invalid"})
		util.LogRus.Errorf("refresh token %s is invalid", req.RefreshToken)
	} else {
		ctx.JSON(http.StatusOK, TokenResponse{Code: 0, Msg: "success", Token: authToken})
	}
}

// CheckRefreshToken 检查 refresh_token 是否有效
func CheckRefreshToken(ctx *gin.Context) {

}
