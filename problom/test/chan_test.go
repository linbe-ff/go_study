package test

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	var (
		chani = make(chan int, 1)
		chana = make(chan int, 1)
	)
	defer func() {
		close(chani)
		close(chana)
	}()
	chana <- 0
	go func() {
		for i := 0; i < 26; i++ {
			<-chana
			fmt.Println(i + 1)
			chani <- i
		}
	}()
	go func() {
		a := 'A'
		for i := 0; i < 26; i++ {
			<-chani
			fmt.Println(fmt.Sprintf("%c", a))
			a++
			chana <- i
		}
	}()
	time.Sleep(time.Second)
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
