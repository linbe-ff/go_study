package for_same_pointer

import (
	"crypto/rand"
	"fmt"
	"io"
	"sync"
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

func TestPC(t *testing.T) {
	var (
		ch = make(chan int)
		wg = &sync.WaitGroup{}
	)
	defer close(ch)

	wg.Add(4)
	go func() {
		for i := 1; i < 5; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			select {
			case i, ok := <-ch:
				if ok {
					wg.Done()
					fmt.Println(i)
				}
			}
		}
	}()
	wg.Wait()
}

func TestCh(t *testing.T) {
	var (
		ch = make(chan int, 1)
	)
	defer func() {
		fmt.Println(<-ch)
	}()

	//close(ch)

	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}

// 定义一个模板方法接口
type TemplateMethodInterface interface {
	PrimitiveOperation1() string
	PrimitiveOperation2() string
	TemplateMethod() string
}

// 实现具体的模板方法
type ConcreteClass struct {
}

func (c *ConcreteClass) PrimitiveOperation1() string {
	return "具体操作1实现"
}

func (c *ConcreteClass) PrimitiveOperation2() string {
	return "具体操作2实现"
}

// 模板方法的具体实现
func (c *ConcreteClass) TemplateMethod() string {
	return c.PrimitiveOperation1() + "\n" + c.PrimitiveOperation2()
}

type ConcreteClass2 struct {
	ConcreteClass
}

type ConcreteClass2List []ConcreteClass2

func TestCXMBFF(t *testing.T) {
	sm := sync.Map{}
	sm.Store("one", 1)
	sm.Store("two", 2)
	sm.Store("two", "2")
	sm.Delete("one")
	value, ok := sm.Load("one")
	value2, ok2 := sm.Load("two")
	fmt.Println(value, ok)
	fmt.Println(value2, ok2)

}

func dddd(i TemplateMethodInterface) {
	i.PrimitiveOperation1()
}
