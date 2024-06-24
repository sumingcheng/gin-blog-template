package test

import (
	"blog/util"
	"fmt"
	"testing"
)

func TestRandStringRunes(t *testing.T) {
	str := util.RandStringRunes(30)
	fmt.Println(str)
}
