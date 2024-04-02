package other

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type Main struct {
	Name    string
	SonList *SonList
}

type (
	Son struct {
		Name string
	}
	SonList []*Son
)

func TestForr(t *testing.T) {
	sonList := make(SonList, 0)
	sonList = append(sonList, &Son{Name: "son1"})
	sonList = append(sonList, &Son{Name: "son2"})
	sonList = append(sonList, &Son{Name: "son3"})
	sonList = append(sonList, &Son{Name: "son4"})
	m := &Main{
		Name:    "main",
		SonList: &sonList,
	}
	for _, son := range *m.SonList {
		log.Println(son)
	}
}

// 將A~Z排列
func TestArrangeA2Z(t *testing.T) {
	var (
		prevMark = 'A' - 1
		mark     = 'A'
	)
	for i := 0; i < 100; i++ {
		if mark > 'Z' {
			prevMark++
			mark = 'A'
		}
		if prevMark == 'A'-1 {
			fmt.Println(fmt.Sprintf("%c", mark))
		} else {
			fmt.Println(fmt.Sprintf("%c%c", prevMark, mark))
		}

		mark++
	}
}

func TestBeforeTime(t *testing.T) {
	var maxTime time.Time
	now := time.Now()
	if maxTime.Before(now) {
		maxTime = now
	}
	fmt.Println(maxTime)
}
func Test1(t *testing.T) {

}
