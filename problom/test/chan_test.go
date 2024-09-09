package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestChan(t *testing.T) {
	var (
		chani = make(chan int, 1)
		chana = make(chan int, 1)
		wg    = &sync.WaitGroup{}
	)
	defer func() {
		close(chani)
		close(chana)
	}()
	chana <- 0
	wg.Add(2)
	go func() {
		for i := 0; i < 26; i++ {
			<-chana
			fmt.Println(i + 1)
			chani <- i
		}
		defer wg.Done()
	}()
	go func() {
		a := 'A'
		for i := 0; i < 26; i++ {
			<-chani
			fmt.Println(fmt.Sprintf("%c", a))
			a++
			chana <- i
		}
		defer wg.Done()
	}()
	wg.Wait()
	//time.Sleep(time.Second)
}

func TestChan2(t *testing.T) {
	var (
		chani = make(chan int, 1)
	)
	chani <- 0
	close(chani)
	val, ok := <-chani
	val2, ok2 := <-chani
	fmt.Println(val, ok)
	fmt.Println(val2, ok2)
}

func f(sli []int) {
	sli[0] = 1
	sli = append(sli, 4, 5)
	sli[1] = 2
}
func TestSli(t *testing.T) {
	s1 := make([]int, 7)
	s2 := s1[3:6]
	f(s2)
	fmt.Println(s1)
	fmt.Println(s2)
	//[0 0 0 1 0 0 0]
	//[1 0 0]
}

func f11() int {
	var r int
	defer func() {
		// 这个操作实际上修改的是 f11 函数栈帧上的 r 副本
		// 因为 return 语句先于 defer 执行
		r++
	}()
	return r
}

func f12(i int) int {
	defer func() {
		i++
	}()
	return i
}

func f13(i int) int {
	defer func(i int) {
		// 都是建的一个副本和外面没关系
		i = 5
	}(i)
	return i
}

func TestZYY(t *testing.T) {
	fmt.Println(f11())
	fmt.Println(f12(0))
	fmt.Println(f13(0))
}
