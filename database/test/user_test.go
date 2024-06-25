package test

import (
	"blog/database"
	"blog/util"
	"testing"
)

func init() {
	util.InitLog("log")
}

func TestGetUserByName(t *testing.T) {
	user := database.GetUserByName("admin")
	if user == nil {
		t.Errorf("get user failed")
	}
	t.Logf("user: %v", user)
}

// go test -v ".\database\test" -run ^TestGetUserByName$ -count=1
