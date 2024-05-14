package redis

import (
	"sync"
	"testing"
)

func TestSecKill(t *testing.T) {
	var (
		r       = &Redis{}
		userIds = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		ch      = make(chan int, 200)
	)
	r.Conn()
	defer func() {
		r.Close()
	}()
	
	go func() {
		for  {
			select {
			case userId <-:
				
			}
		}
	}()

	// 模拟请求打进来
	for _, userId := range userIds {
		ch <-userId
	}
	
	

	r.CreateQueue()
}
