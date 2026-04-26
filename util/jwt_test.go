package util

import (
	"testing"
	"time"
)

func TestGenAndVerifyAuthToken(t *testing.T) {
	secret := "test-secret-key-for-unit-test"

	tests := []struct {
		name      string
		uid       int
		expire    time.Duration
		wantErr   bool
	}{
		{"正常 token", 42, time.Hour, false},
		{"uid=0", 0, time.Hour, false},
		{"短过期", 1, time.Second, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GenAuthToken(tt.uid, secret, tt.expire)
			if err != nil {
				t.Fatalf("GenAuthToken failed: %v", err)
			}
			if token == "" {
				t.Fatal("token should not be empty")
			}

			uid, err := VerifyAuthToken(token, secret)
			if (err != nil) != tt.wantErr {
				t.Fatalf("VerifyAuthToken error = %v, wantErr %v", err, tt.wantErr)
			}
			if uid != tt.uid {
				t.Errorf("uid = %d, want %d", uid, tt.uid)
			}
		})
	}
}

func TestVerifyAuthToken_Expired(t *testing.T) {
	secret := "test-secret"
	token, _ := GenAuthToken(1, secret, -time.Hour)
	_, err := VerifyAuthToken(token, secret)
	if err == nil {
		t.Fatal("should reject expired token")
	}
}

func TestVerifyAuthToken_WrongSecret(t *testing.T) {
	token, _ := GenAuthToken(1, "secret-a", time.Hour)
	_, err := VerifyAuthToken(token, "secret-b")
	if err == nil {
		t.Fatal("should reject token with wrong secret")
	}
}

func TestGenRefreshToken(t *testing.T) {
	token1, err := GenRefreshToken()
	if err != nil {
		t.Fatalf("GenRefreshToken failed: %v", err)
	}
	token2, _ := GenRefreshToken()
	if token1 == token2 {
		t.Fatal("two refresh tokens should be different")
	}
	if len(token1) != 64 {
		t.Errorf("refresh token length = %d, want 64", len(token1))
	}
}
