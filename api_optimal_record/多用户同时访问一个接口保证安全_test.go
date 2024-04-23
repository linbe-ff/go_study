package api_optimal_record

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 互斥锁 ---- begin ---- 互斥锁
var (
	mutex      sync.Mutex
	sharedData = make(map[int]int)
)

func TestManyGroupIn(t *testing.T) {

	for i := 1; i < 11; i++ {
		go f1(i)
	}
	time.Sleep(1 * time.Second)
}

func f1(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sharedData[i] = i
	fmt.Println("i = ", i)
}

// 互斥锁 ---- end ---- 互斥锁

// chan ---- begin ---- chan
var (
	dataChan    = make(chan int, 10)
	sharedData2 = make(map[int]int)
)

func TestManyGroupInChan(t *testing.T) {

	for i := 1; i < 11; i++ {
		go f2(i)
	}
	go selectChan()
	time.Sleep(1 * time.Second)
}

func f2(i int) {
	dataChan <- i
}
func selectChan() {
	for {
		select {
		case dd := <-dataChan:
			sharedData2[dd] = dd
			fmt.Println(dd)
		default:
		}
	}
}

// chan ---- end ---- chan

// 原子操作（Atomic Operations） ---- begin ---- 原子操作（Atomic Operations）
var (
	count    int32  = 0
	counter  int32  = 0
	counter2 uint32 = 0
)

func TestManyGroupInAtomic(t *testing.T) {
	//swapValue(100)

	//if incrementIfEqual(0, 1) {
	//	//	fmt.Println("Counter incremented from 0 to 1")
	//	//} else {
	//	//	fmt.Println("Counter was not 0, so it was not incremented")
	//	//}

	for i := 0; i < 1000; i++ {
		go increment()
	}
	// 最终打印的 counter 应该是 1000，即使有多个 goroutine 同时进行 increment 操作
	fmt.Println(atomic.LoadUint32(&counter2))
}

func swapValue(newVal int32) {
	oldVal := atomic.SwapInt32(&count, newVal)
	fmt.Printf("Old value: %d, New value: %d\n", oldVal, newVal)
	fmt.Println(count)
}

func incrementIfEqual(expected, new int32) bool {
	return atomic.CompareAndSwapInt32(&counter, expected, new)
}

func increment() {
	atomic.AddUint32(&counter2, 1)
}

// 原子操作（Atomic Operations） ---- end ---- 原子操作（Atomic Operations）

//队列系统：
//幂等性设计：
