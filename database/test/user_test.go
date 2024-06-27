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

func TestCreateUser(t *testing.T) {
	database.CreateUser("素明诚", "123456")
}

func TestDeleteUser(t *testing.T) {
	database.DeleteUser("素明诚")
}

// go test -v ".\database\test" -run ^TestGetUserByName$ -count=1
