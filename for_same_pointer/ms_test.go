package for_same_pointer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//编写函数，实现两个协程轮流输出A 1 B 2 C 3 .... Z 26，
//其中一个goroutine 输出数字，
//另一个goroutine输出字母，并在main函数中调用，验证你的逻辑

var (
	i  = 1
	a  = 'A'
	wg = sync.WaitGroup{}
)

func TestPrint(t *testing.T) {
	// 分别使用两个缓存为1的chan，来控制两个g的打印顺序
	strChan := make(chan int, 1)
	numChan := make(chan int, 1)

	strChan <- 0 // 先往字符chan中塞入，此时strChan再塞入会堵塞

	// 负责打印字母
	go func() {
		for i := 65; i <= 90; i++ {
			<-strChan                          // strChan取出，因为之前先塞入了，所以此处不会堵塞，会直接打印字符A..Z
			fmt.Printf("%v ", string(rune(i))) // 打印字母
			numChan <- i                       // numChan 塞入，塞入后，另一个g的numChan取出操作才能进行
		}
		return
	}()

	// 负责打印数字
	go func() {
		for i := 1; i <= 26; i++ {
			<-numChan            // 一直阻塞，直到字母被打印，这样每次数字都是在字母后面被打印的
			fmt.Printf("%v ", i) /// 打印数字
			strChan <- i         // strChan塞入，此处塞入后，上面协程的strChan取出操作才能进行，才会打印字母，这样保证了打印完数字后，紧接着打印字母
		}
		return
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()
}

func number() {
	wg.Done()
	fmt.Println(i)
	i++
}

func char() {
	fmt.Println(fmt.Sprintf("%c", a))
	a++
}
