package test

import (
	"blog/util"
	"fmt"
	"testing"
)

type Person struct {
	Name  string
	Age   int
	Score int
}

var arrayTest = []Person{
	{Name: "Alice", Age: 28, Score: 90},
	{Name: "Bob", Age: 24, Score: 80},
	{Name: "Cindy", Age: 22, Score: 70},
	{Name: "David", Age: 26, Score: 60},
	{Name: "Eva", Age: 30, Score: 50},
}

func TestForEach(t *testing.T) {
	var countScore = 0
	util.ForEach(arrayTest, func(p Person, index int, arr []Person) {
		countScore += p.Score
		fmt.Println(p)
	})
}

func TestMap(t *testing.T) {
	newArray := util.Map(arrayTest, func(p Person, i int, people []Person) Person {
		return Person{
			Name:  p.Name,
			Age:   p.Age,
			Score: p.Score * 10, // 将分数乘以2
		}
	})

	// 打印新数组的内容，查看分数是否已更新
	for _, p := range newArray {
		fmt.Printf("%s: %d\n", p.Name, p.Score)
	}
}
