package for_same_pointer

import (
	"fmt"
	"testing"
)

func TestExpCap(t *testing.T) {
	// 创建一个map，初始容量为5
	m := make(map[int]string, 5)

	// 填充map
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	m[4] = "four"
	m[5] = "five"

	// 打印map的容量和长度
	fmt.Printf("Capacity: , Length: %d\n", len(m))

	// 继续添加元素，触发扩容
	m[6] = "six"

	// 扩容后打印容量和长度
	fmt.Printf("Capacity: , Length: %d\n", len(m))
}
