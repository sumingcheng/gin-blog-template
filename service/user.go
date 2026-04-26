package service

import (
	"blog/database"
	"blog/model"
	"blog/util"
	"time"
)

var jwtSecret string

func init() {
	jwtSecret = util.CreateConfig("key").GetString("jwt")
}

type LoginResult struct {
	Uid          int    `json:"uid"`
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}

// Register 用户注册
func Register(username, password string) (*database.User, *model.BizError) {
	if existing := database.GetUserByName(username); existing != nil {
		return nil, model.ErrUserExists
	}
	hashed, err := util.HashPassword(password)
	if err != nil {
		util.LogRus.Errorf("密码哈希失败: %s", err)
		return nil, model.ErrInternalServer
	}
	user, err := database.CreateUser(username, hashed)
	if err != nil {
		util.LogRus.Errorf("创建用户失败: %s", err)
		return nil, model.ErrInternalServer
	}
	return user, nil
}

// Login 登录
func Login(username, password string) (*LoginResult, *model.BizError) {
	user := database.GetUserByName(username)
	if user == nil {
		return nil, model.ErrUserNotFound
	}
	if !util.CheckPassword(password, user.PassWd) {
		return nil, model.ErrWrongPassword
	}

	util.LogRus.Infof("user %s(%d) login", username, user.Id)

	authToken, err := util.GenAuthToken(user.Id, jwtSecret, database.TokenExpire+24*time.Hour)
	if err != nil {
		util.LogRus.Errorf("生成 auth token 失败: %s", err)
		return nil, model.ErrTokenGenFailed
	}

	refreshToken, err := util.GenRefreshToken()
	if err != nil {
		util.LogRus.Errorf("生成 refresh token 失败: %s", err)
		return nil, model.ErrTokenGenFailed
	}

	if err := database.SetToken(refreshToken, authToken); err != nil {
		util.LogRus.Errorf("写入 token 对失败: %s", err)
		return nil, model.ErrTokenStoreFailed
	}

	return &LoginResult{
		Uid:          user.Id,
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}, nil
}

// ChangePassword 修改密码
func ChangePassword(uid int, oldPass, newPass string) *model.BizError {
	user := database.GetUserById(uid)
	if user == nil {
		return model.ErrUserNotFound
	}
	if !util.CheckPassword(oldPass, user.PassWd) {
		return model.ErrWrongPassword
	}
	hashed, err := util.HashPassword(newPass)
	if err != nil {
		util.LogRus.Errorf("密码哈希失败: %s", err)
		return model.ErrInternalServer
	}
	if err := database.UpdateUserPassword(uid, hashed); err != nil {
		util.LogRus.Errorf("更新密码失败: %s", err)
		return model.ErrInternalServer
	}
	return nil
}

// GetProfile 获取用户信息
func GetProfile(uid int) (*database.User, *model.BizError) {
	user := database.GetUserById(uid)
	if user == nil {
		return nil, model.ErrUserNotFound
	}
	return user, nil
}

// RefreshAuthToken 刷新 token
func RefreshAuthToken(refreshToken string) (string, *model.BizError) {
	if refreshToken == "" {
		return "", model.ErrTokenExpired
	}
	authToken := database.GetToken(refreshToken)
	if authToken == "" {
		return "", model.ErrRefreshInvalid
	}
	return authToken, nil
}

// Logout 登出
func Logout(refreshToken string) *model.BizError {
	if refreshToken == "" {
		return model.ErrBadRequest
	}
	if err := database.RmToken(refreshToken); err != nil {
		util.LogRus.Errorf("删除 refresh token 失败: %s", err)
		return model.ErrInternalServer
	}
	return nil
}

// GetUidFromToken 从 JWT 解析 uid
func GetUidFromToken(token string) int {
	uid, err := util.VerifyAuthToken(token, jwtSecret)
	if err != nil {
		return 0
	}
	return uid
}
