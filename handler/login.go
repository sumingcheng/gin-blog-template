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
	Token string `json:"token"`
}

// Login godoc
// @Summary User login
// @Description Perform user login with username and password.
// @Tags authentication
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   user formData string true "User name"
// @Param   pass formData string true "Password"
// @Success 200 {object} LoginResponse "Successful login, token returned"
// @Failure 400 {object} LoginResponse "Invalid request, user name or password format error"
// @Failure 403 {object} LoginResponse "Unauthorized, wrong user name or password"
// @Failure 500 {object} LoginResponse "Internal server error, token generation failed"
// @Router /login [get]
func Login(ctx *gin.Context) {
	name := ctx.PostForm("user") //从post form中获取参数
	pass := ctx.PostForm("pass")
	if len(name) == 0 {
		ctx.JSON(http.StatusBadRequest, LoginResponse{1, "must indicate user name", 0, ""})
		return
	}
	if len(pass) != 32 {
		ctx.JSON(http.StatusBadRequest, LoginResponse{2, "invalid password", 0, ""})
		return
	}
	user := database.GetUserByName(name)
	if user == nil {
		ctx.JSON(http.StatusForbidden, LoginResponse{3, "用户名不存在", 0, ""})
		return
	}
	if user.PassWd != pass {
		ctx.JSON(http.StatusForbidden, LoginResponse{4, "密码错误", 0, ""})
		return
	}

	util.LogRus.Infof("user %s(%d) login", name, user.Id)
	// 用户名、密码正确，向客户端返回一个token
	header := util.DefaultHeader
	payload := util.JwtPayload{ //payload以明文形式编码在token中，server用自己的密钥可以校验该信息是否被篡改过
		Issue:       "blog",
		IssueAt:     time.Now().Unix(),                                               //因为每次的IssueAt不同，所以每次生成的token也不同
		Expiration:  time.Now().Add(database.TokenExpire).Add(24 * time.Hour).Unix(), //(7+1)天后过期，需要重新登录，假设24小时内用户肯定要重启浏览器
		UserDefined: map[string]string{middleware.UidInToken: strconv.Itoa(user.Id)}, //用户自定义字段。如果token里包含敏感信息，请结合https使用
	}
	// 生成token
	// accessToken, err := util.GetAccessToken(user.Id)
	if accessToken, err := util.GenJWT(header, payload, middleware.KeyConfig.GetString("jwt")); err != nil {
		util.LogRus.Errorf("生成token失败: %s", err)
		ctx.JSON(http.StatusInternalServerError, LoginResponse{5, "token生成失败", 0, ""})
		return
	} else {
		refreshToken, _ := util.GetRefreshToken()    //生成RefreshToken
		database.SetToken(refreshToken, accessToken) //把<refreshToken, authToken>写入redis
		//response header里会有一条 Set-Cookie: auth_token=xxx; other_key=other_value，浏览器后续请求会自动把同域名下的cookie再放到request header里来，即request header里会有一条Cookie: auth_token=xxx; other_key=other_value
		ctx.SetCookie("refresh_token", refreshToken, //注意：受cookie本身的限制，这里的token不能超过4K
			int(database.TokenExpire.Seconds()), //maxAge，cookie的有效时间，时间单位秒。如果不设置过期时间，默认情况下关闭浏览器后cookie被删除
			"/",                                 //path，cookie存放目录
			"localhost",                         //cookie从属的域名,不区分协议和端口。如果不指定domain则默认为本host(如b.a.com)，如果指定的domain是一级域名(如a.com)，则二级域名(b.a.com)下也可以访问
			false,                               //是否只能通过https访问
			true,                                //是否允许别人通过js获取自己的cookie，设为false防止XSS攻击
		)
		ctx.JSON(http.StatusOK, LoginResponse{0, "success", user.Id, accessToken})
		return
	}
}
func GetAuthToken(ctx *gin.Context) {
	refreshToken := ctx.Param("refresh_token")
	authToken := database.GetToken(refreshToken)
	ctx.String(http.StatusOK, authToken)
}
