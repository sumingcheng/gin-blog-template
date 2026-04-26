package model

import "net/http"

// BizError 业务错误
type BizError struct {
	HTTPStatus int
	Code       int
	Msg        string
}

func (e *BizError) Error() string {
	return e.Msg
}

// 通用错误
var (
	ErrBadRequest     = &BizError{http.StatusBadRequest, 40000, "请求参数错误"}
	ErrUnauthorized   = &BizError{http.StatusUnauthorized, 40100, "未登录或登录已过期"}
	ErrForbidden      = &BizError{http.StatusForbidden, 40300, "无权操作"}
	ErrNotFound       = &BizError{http.StatusNotFound, 40400, "资源不存在"}
	ErrInternalServer = &BizError{http.StatusInternalServerError, 50000, "服务器内部错误"}
)

// 用户模块 1xxxx
var (
	ErrUserNotFound    = &BizError{http.StatusOK, 10001, "用户不存在"}
	ErrWrongPassword   = &BizError{http.StatusOK, 10002, "密码不正确"}
	ErrUserExists      = &BizError{http.StatusOK, 10003, "用户名已存在"}
	ErrTokenGenFailed  = &BizError{http.StatusInternalServerError, 10004, "Token 生成失败"}
	ErrTokenStoreFailed = &BizError{http.StatusInternalServerError, 10005, "Token 存储失败"}
	ErrTokenExpired    = &BizError{http.StatusOK, 10006, "登录已过期，请重新登录"}
	ErrRefreshInvalid  = &BizError{http.StatusOK, 10007, "Refresh token 已失效"}
)

// 博客模块 2xxxx
var (
	ErrBlogNotFound = &BizError{http.StatusNotFound, 20001, "博客不存在"}
	ErrBlogUpdate   = &BizError{http.StatusInternalServerError, 20002, "更新博客失败"}
)
