package leecode

import (
	"fmt"
	"testing"
)

func Test冒泡(t *testing.T) {
	var numberList = []int{1, 39, 2, 9, 7, 54, 11}
	var numberList2 = []int{1, 39, 2, 9, 7, 54, 11}

	numberList = 冒泡1(numberList)
	fmt.Println(numberList)

	numberList = 冒泡2(numberList2)
	fmt.Println(numberList2)
}

func TestSelectSort(t *testing.T) {
	var numberList = []int{1, 39, 2, 9, 7, 54, 11}

	numberList = selectSort(numberList)

	fmt.Println(numberList)
}

func TestInsertSort(t *testing.T) {
	var numberList = []int{1, 39, 2, 9, 7, 54, 11}

	numberList = insertSort(numberList)

	fmt.Println(numberList)
}

func TestShellSort(t *testing.T) {
	var numberList = []int{1, 39, 2, 9, 7, 54, 11}

	numberList = shellSort(numberList)

	fmt.Println(numberList)
}
