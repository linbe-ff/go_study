package excel

import "fmt"

// 將A~Z排列
func ArrangeA2Z(rangeCount int) []string {
	var (
		rangeChar = make([]string, 0)
		prevMark  = 'A' - 1
		mark      = 'A'
	)
	for i := 0; i < rangeCount; i++ {
		if mark > 'Z' {
			prevMark++
			mark = 'A'
		}
		if prevMark == 'A'-1 {
			//fmt.Println(fmt.Sprintf("%c", mark))
			rangeChar = append(rangeChar, fmt.Sprintf("%c", mark))
		} else {
			//fmt.Println(fmt.Sprintf("%c%c", prevMark, mark))
			rangeChar = append(rangeChar, fmt.Sprintf("%c%c", prevMark, mark))
		}

		mark++
	}
	return rangeChar
}

func AddOne(a *int) *int {
	*a += 1
	return a
}
