package scope

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

}

func TestSlice1(t *testing.T) {
	// 示例 1。
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
	// The length of s1: 5
	// The capacity of s1: 5
	// The value of s1: [0 0 0 0 0]
	// The length of s2: 5
	// The capacity of s2: 8
	// The value of s2: [0 0 0 0 0]
}

func TestSlice2(t *testing.T) {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	//s4的容量就是其底层数组的长度8, 减去上述切片表达式中的那个起始索引3，即5。
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	//The length of s4: 3
	//The capacity of s4: 5
	//The value of s4: [4 5 6]
}

func TestSlice3(t *testing.T) {

}
