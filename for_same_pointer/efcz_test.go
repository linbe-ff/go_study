package for_same_pointer

import (
	"fmt"
	"testing"
)

func TestEF(t *testing.T) {
	var (
		//sli = []int{1, 2, 3, 4, 5, 6}
		sli = []int{1, 5, 4, 3, 2, 6}
		//targ = 3
	)
	//fmt.Println(fs(sli, targ))

	//fmt.Println(findMax(sli))
	fmt.Println(findMaxSec(sli))

}

func fs(sli []int, targ int) int {
	left := 0
	right := len(sli) - 1

	for left <= right {
		mid := left + (right-left)/2
		if sli[mid] == targ {
			return mid
		} else if sli[mid] < targ {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 都没有
	return 0
}

func findMax(sli []int) int {
	var maxNum int = 0
	var maxIndex int = 0
	//var secMax int = 0
	for i, v := range sli {
		if v > maxNum {
			maxIndex = i
		}
	}
	sli = append(sli[:maxIndex], sli[maxIndex+1:]...)
	for _, v := range sli {
		if v > maxNum {
			maxNum = v
			maxIndex = i
		}
	}
	return maxNum
}

func findMaxSec(sli []int) int {
	var maxNum int = 0
	var maxSecNum int = 0
	//var maxIndex int = 0
	//var secMax int = 0
	for i, v := range sli {
		if maxSecNum == 0 {
			maxSecNum = v
		}
		if v > maxNum {
			maxNum = i
		}
		if v > maxSecNum && v != maxNum {
			maxSecNum = v
		}
	}
	//sli = append(sli[:maxIndex], sli[maxIndex+1:]...)
	//for _, v := range sli {
	//	if v > maxNum {
	//		maxNum = v
	//		maxIndex = i
	//	}
	//}
	return maxSecNum
}
