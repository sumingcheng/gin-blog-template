package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRequestID_Generate(t *testing.T) {
	r := gin.New()
	r.Use(RequestID())
	r.GET("/test", func(c *gin.Context) {
		rid, _ := c.Get(RequestIDKey)
		c.String(http.StatusOK, rid.(string))
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 响应头应该有 X-Request-ID
	rid := w.Header().Get(RequestIDKey)
	if rid == "" {
		t.Fatal("response should have X-Request-ID header")
	}
	if len(rid) != 32 {
		t.Errorf("request ID length = %d, want 32 (16 bytes hex)", len(rid))
	}
	// body 应该也是同一个 ID
	if w.Body.String() != rid {
		t.Errorf("body = %q, header = %q, should match", w.Body.String(), rid)
	}
}

func TestRequestID_Passthrough(t *testing.T) {
	r := gin.New()
	r.Use(RequestID())
	r.GET("/test", func(c *gin.Context) {
		rid, _ := c.Get(RequestIDKey)
		c.String(http.StatusOK, rid.(string))
	})

	customID := "my-custom-request-id-12345"
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set(RequestIDKey, customID)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 应该透传客户端提供的 ID
	if w.Header().Get(RequestIDKey) != customID {
		t.Errorf("should passthrough custom ID, got %q", w.Header().Get(RequestIDKey))
	}
	if w.Body.String() != customID {
		t.Errorf("body = %q, want %q", w.Body.String(), customID)
	}
}

func TestRequestID_Unique(t *testing.T) {
	r := gin.New()
	r.Use(RequestID())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, httptest.NewRequest(http.MethodGet, "/test", nil))
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, httptest.NewRequest(http.MethodGet, "/test", nil))

	id1 := w1.Header().Get(RequestIDKey)
	id2 := w2.Header().Get(RequestIDKey)
	if id1 == id2 {
		t.Error("two requests should have different IDs")
	}
}
