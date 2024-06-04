package for_same_pointer

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

type student struct {
	Name string
	Age  int
}

// 22已经无此问题
func TestForSamePointer(t *testing.T) {
	//定义map
	m := make(map[string]*student)

	//定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//将数组依次添加到map中
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	//打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}

func TestAddOne(t *testing.T) {
	var a = 0
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		fmt.Println(*addOne(&a))
	}
}
func addOne(a *int) *int {
	*a += 1
	return a
}

func TestSlic(t *testing.T) {
	//str := "2024-04-24T16:00:00.000Z"
	fmt.Println(GenerateSalt())
	fmt.Println(GenerateSalt())
	fmt.Println(GenerateSalt())
	fmt.Println(GenerateSalt())
}

func GenerateSalt() string {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}
