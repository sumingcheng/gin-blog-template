package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

type testResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func parseResponse(t *testing.T, w *httptest.ResponseRecorder) testResponse {
	t.Helper()
	var resp testResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("parse response failed: %v, body: %s", err, w.Body.String())
	}
	return resp
}

// ---- Login 参数校验 ----

func TestLogin_EmptyBody(t *testing.T) {
	r := gin.New()
	r.POST("/api/login", Login)

	req := httptest.NewRequest(http.MethodPost, "/api/login", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

func TestLogin_ShortPassword(t *testing.T) {
	r := gin.New()
	r.POST("/api/login", Login)

	body := `{"user":"admin","pass":"123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000 (bad request)", resp.Code)
	}
}

func TestLogin_ShortUsername(t *testing.T) {
	r := gin.New()
	r.POST("/api/login", Login)

	body := `{"user":"ab","pass":"123456"}`
	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

// ---- Register 参数校验 ----

func TestRegister_EmptyBody(t *testing.T) {
	r := gin.New()
	r.POST("/api/register", Register)

	req := httptest.NewRequest(http.MethodPost, "/api/register", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

func TestRegister_UsernameTooLong(t *testing.T) {
	r := gin.New()
	r.POST("/api/register", Register)

	body := `{"user":"abcdefghijklmnopqrstuvwxyz","pass":"123456"}`
	req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

// ---- Blog 参数校验 ----

func TestBlogList_InvalidUID(t *testing.T) {
	r := gin.New()
	r.GET("/api/blog/list", BlogList)

	// page/size 非法不会报错（omitempty），但 uid=abc 也不报错因为是 query 参数
	req := httptest.NewRequest(http.MethodGet, "/api/blog/list?page=-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

func TestBlogDetail_InvalidBid(t *testing.T) {
	r := gin.New()
	r.GET("/api/blog/:bid", BlogDetail)

	req := httptest.NewRequest(http.MethodGet, "/api/blog/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

func TestBlogCreate_EmptyBody(t *testing.T) {
	r := gin.New()
	r.POST("/api/blog/create", func(c *gin.Context) {
		c.Set("uid", 1)
		BlogCreate(c)
	})

	req := httptest.NewRequest(http.MethodPost, "/api/blog/create", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

func TestBlogUpdate_InvalidBlogId(t *testing.T) {
	r := gin.New()
	r.POST("/api/blog/update", func(c *gin.Context) {
		c.Set("uid", 1)
		BlogUpdate(c)
	})

	body := `{"blogId":0,"title":"t","article":"a"}`
	req := httptest.NewRequest(http.MethodPost, "/api/blog/update", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

func TestBlogBelong_EmptyBody(t *testing.T) {
	r := gin.New()
	r.POST("/api/blog/belong", BlogBelong)

	req := httptest.NewRequest(http.MethodPost, "/api/blog/belong", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

// ---- ChangePassword 参数校验 ----

func TestChangePassword_EmptyBody(t *testing.T) {
	r := gin.New()
	r.POST("/api/user/password", func(c *gin.Context) {
		c.Set("uid", 1)
		ChangePassword(c)
	})

	req := httptest.NewRequest(http.MethodPost, "/api/user/password", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

func TestChangePassword_ShortNewPass(t *testing.T) {
	r := gin.New()
	r.POST("/api/user/password", func(c *gin.Context) {
		c.Set("uid", 1)
		ChangePassword(c)
	})

	body := `{"oldPass":"123456","newPass":"12"}`
	req := httptest.NewRequest(http.MethodPost, "/api/user/password", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := parseResponse(t, w)
	if resp.Code != 40000 {
		t.Errorf("code = %d, want 40000", resp.Code)
	}
}

// ---- 统一响应格式验证 ----

func TestResponse_Format(t *testing.T) {
	r := gin.New()
	r.POST("/api/login", Login)

	req := httptest.NewRequest(http.MethodPost, "/api/login", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var raw map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &raw); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	if _, ok := raw["code"]; !ok {
		t.Error("response missing 'code' field")
	}
	if _, ok := raw["msg"]; !ok {
		t.Error("response missing 'msg' field")
	}
}
