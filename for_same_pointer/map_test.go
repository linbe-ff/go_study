package for_same_pointer

import (
	"context"
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

func TestPrintChan(t *testing.T) {
	c1 := make(chan int, 1)
	c1 <- 1
	defer close(c1)
	fmt.Println(<-c1)
}

func TestCtx(t *testing.T) {
	ctx := context.Background()
	ctx = GetContext()
	fmt.Println(ctx)
}

func GetContext() (ctx context.Context) {
	return context.WithValue(context.Background(), "common.DB_HASH", "db_hash")
}
