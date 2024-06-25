package test

import (
	"blog/database"
	"sync"
	"testing"
)

func TestGetBlogDBConnection(t *testing.T) {
	const C = 100
	wg := sync.WaitGroup{}
	wg.Add(C)
	for i := 0; i < C; i++ {
		go database.GetBlogDBConnection()
		wg.Done()
	}
	wg.Wait()
}

// go test -v ".\database\test" -run ^TestGetBlogDBConnection$ -count=1
