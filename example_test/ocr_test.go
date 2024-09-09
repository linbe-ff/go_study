package example

import (
	"fmt"
	"testing"
)

type ddd11 struct {
	dd string
}

func TestR(t *testing.T) {
	dd := ddd11{dd: "1"}
	dd2 := ddd11{dd: "1"}
	fmt.Println(dd == dd2)
}

func TestRa(t *testing.T) {
	m := make(map[string]*string)
	dd := "dsafsfds"
	m["a"] = &dd
	fmt.Println(*m["a"]) // dsafsfds
	dd = "dsafdf"
	fmt.Println(*m["a"]) // dsafdf
}

type MyType int

func (m MyType) ValueMethod() {
	fmt.Println("ValueMethod called")
}

func (m *MyType) PointerMethod() {
	fmt.Println("PointerMethod called")
}

var v MyType
var p = &v

func TestRb(t *testing.T) {
	v.ValueMethod() // 正常调用
	p.ValueMethod() // 正常调用，Go 自动将 p 解引用

	v.PointerMethod() // 正常调用，Go 自动取 v 的地址
	p.PointerMethod() // 正常调用
}
