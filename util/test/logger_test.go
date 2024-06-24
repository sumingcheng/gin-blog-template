package test

import (
	"blog/util"
	"testing"
)

func TestLogger(t *testing.T) {
	util.InitLog("log")

	util.LogRus.Info("This is a test log")
	util.LogRus.Warn("This is a test log")
	util.LogRus.Error("This is a test log")
	util.LogRus.Debug("This is a test log")
	//util.LogRus.Panic("This is a test log")
}
