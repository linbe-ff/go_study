package time20240909

import (
	"fmt"
	"sync"
	"testing"
)

func TestMS(t *testing.T) {
	var (
		c1      = make(chan int, 10)
		wg      = &sync.WaitGroup{}
		product = 100
	)

	defer close(c1)
	wg.Add(2)
	// c
	go func() {
		for i := 0; i < 100; i++ {
			c1 <- i
		}
		defer wg.Done()
	}()

	// p
	go func() {
		sum := 0
		for val := range c1 {
			sum++
			if sum == product {
				break
			}
			fmt.Println(val)
		}
		defer wg.Done()
	}()
	wg.Wait()
}

func TestMP(t *testing.T) {
	var (
		s1 = []int{1, 4, 2, 3, 7, 8, 6, 5}
	)
	count := len(s1)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if s1[i] < s1[j] {
				s1[i], s1[j] = s1[j], s1[i]
			}
		}
	}
	fmt.Println(s1)
}
