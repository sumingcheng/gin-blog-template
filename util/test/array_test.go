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
	newArray := util.Map(arrayTest, func(p Person, index int, people []Person) Person {
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

func TestFilter(t *testing.T) {
	filter := util.Filter(arrayTest, func(person Person, index int, people []Person) bool {
		return person.Score > 60
	})

	for i := 0; i < len(filter); i++ {
		fmt.Printf("%+v", filter[i])
	}
}

func TestReduce(t *testing.T) {
	totalScore := util.Reduce(arrayTest, 0, func(acc int, person Person) int {
		return acc + person.Score
	})

	fmt.Println("总成绩为：", totalScore)
}

func TestEvery(t *testing.T) {
	allPositive := util.Every(arrayTest, func(person Person, i int, people []Person) bool {
		return person.Score > 60
	})

	fmt.Println("allPositive:", allPositive)
}

func TestSome(t *testing.T) {
	some := util.Some(arrayTest, func(person Person, i int, people []Person) bool {
		return person.Score > 60
	})

	fmt.Println("some:", some)
}

func TestPush(t *testing.T) {
	newArr, _ := util.Push(arrayTest, Person{
		Name: "素明诚", Age: 28, Score: 90,
	})

	fmt.Printf("%+v", newArr)
}

func TestUnshift(t *testing.T) {
	newArr, _ := util.Unshift(arrayTest, Person{
		Name: "素明诚", Age: 28, Score: 90,
	})

	fmt.Printf("%+v", newArr)
}

func TestPop(t *testing.T) {
	_, arrs, _ := util.Pop(arrayTest)

	fmt.Printf("%+v", arrs)
}
