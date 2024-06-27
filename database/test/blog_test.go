package test

import (
	"blog/database"
	"fmt"
	"testing"
)

func TestUpdateBlog(t *testing.T) {
	blog := database.Blog{Id: 1, Title: "文章一", Article: "文章一来自海洋的声音，体验大海的奥秘。电波像海浪一样，心愿如星空般辽阔。"}

	if err := database.UpdateBlog(&blog); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
