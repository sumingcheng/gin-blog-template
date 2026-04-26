package util

import "testing"

func TestHashPassword(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}
	if len(hash) == 0 {
		t.Fatal("hash should not be empty")
	}
	if hash == password {
		t.Fatal("hash should not equal plaintext")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "testpassword123"
	hash, _ := HashPassword(password)

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"正确密码", "testpassword123", true},
		{"错误密码", "wrongpassword", false},
		{"空密码", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckPassword(tt.input, hash)
			if result != tt.expected {
				t.Errorf("CheckPassword(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
