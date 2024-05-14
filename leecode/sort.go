package leecode

func 冒泡1(n []int) []int {
	var isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				var temp = n[i]
				n[i] = n[i+1]
				n[i+1] = temp
				isDone = false
			}
			i++
		}
	}
	return n
}

func 冒泡2(numberList []int) []int {
	n := len(numberList)
	for i := 0; i < n-1; i++ {
		// 提前退出冒泡循环的标志位
		swapped := false
		for j := 0; j < n-i-1; j++ {
			// 交换相邻的两个元素，如果它们的顺序错误
			if numberList[j] > numberList[j+1] {
				numberList[j], numberList[j+1] = numberList[j+1], numberList[j]
				swapped = true
			}
		}
		// 如果在某次遍历中没有发生任何交换，说明数列已经是有序的，可以提前结束
		if !swapped {
			break
		}
	}
	return numberList
}

// []int{1, 39, 2, 9, 7, 54, 11}
// 选择排序
// 做法： 下标遍历对比该下标后面所有元素, 找到最小的，放到前面
func selectSort(li []int) []int {
	for i := 0; i < len(li)-1; i++ {
		pos := i
		for j := i + 1; j < len(li); j++ {
			if li[pos] > li[j] {
				pos = j
			}
		}
		li[i], li[pos] = li[pos], li[i]
	}
	return li
}

// []int{1, 39, 2, 9, 7, 54, 11}
// 插入排序
// 做法：后一个和前面所有做比较并交换
func insertSort(li []int) []int {
	for i := 1; i < len(li); i++ {
		tmp := li[i]
		j := i - 1
		for j >= 0 && tmp < li[j] {
			li[j+1] = li[j]
			j--
		}
		li[j+1] = tmp
	}
	return li
}

// []int{1, 39, 2, 9, 7, 54, 11}
// 希尔排序
// 做法：
func shellSort(li []int) []int {
	for gap := len(li) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(li); i++ {
			tmp := li[i]
			j := i - gap
			for j >= 0 && tmp < li[j] {
				li[j+gap] = li[j]
				j -= gap
			}
			li[j+gap] = tmp
		}
	}
	return li
}
