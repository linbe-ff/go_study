package example

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	count = 0
)

func worker(c context.Context, wg *sync.WaitGroup, id int) {

	defer wg.Done()
	count++
	//fmt.Println("开始执行 %", count)
	select {
	case <-time.After(time.Second):
		fmt.Println("执行完成 %", id)
	case <-c.Done():
		fmt.Println("请求超时 %", id)
	}
}

func TestTimeout(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {

		//time.Sleep(time.Second)
		go func() {
			defer wg.Done()
			count++
			//fmt.Println("开始执行 %", count)
			select {
			case <-time.After(time.Second):
				//fmt.Println("执行完成 %", i)
			case <-ctx.Done():
				fmt.Println("请求超时 %", count)
				//cancel()
			}
		}()

		//go worker(ctx, &wg, i)
	}

	go func() {
		wg.Wait()
		cancel()
	}()
}
