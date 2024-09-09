package example

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	var h = 65
	var c, d, e = make(chan int), make(chan int), make(chan int)
	defer func() {
		close(c)
		close(d)
		close(e)
	}()

	go func() {
		for i := 0; i < 26; i++ {
			<-c
			fmt.Print(i + 1)
			d <- 1
		}
	}()
	go func() {
		for i := h; i < h+26; i++ {
			<-d
			fmt.Printf("%c", i)
			if i < h+26-1 {
				c <- 1
			}
		}
		e <- 1
	}()
	c <- 1
	<-e
}
