package for_same_pointer

import (
	"fmt"
	"testing"
	"time"
)

func TestCS(t *testing.T) {
	var (
		c0 = make(chan int, 10)
	)
	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("0生產了", i)
			c0 <- i
		}
	}()
	go func() {
		for v := range c0 {
			fmt.Println("0消費了", v)
		}
	}()
	time.Sleep(time.Second)
}
