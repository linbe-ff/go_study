package example

import (
	"fmt"
	"testing"
)

type TeatMap struct {
	Id   int
	Pid  int
	Name string
}

func TestDmmm(t *testing.T) {
	list := []*TeatMap{
		{
			Id:   1,
			Pid:  0,
			Name: "1",
		},
		{
			Id:   2,
			Pid:  1,
			Name: "2",
		},
		{
			Id:   3,
			Pid:  0,
			Name: "3",
		},
		{
			Id:   4,
			Pid:  1,
			Name: "4",
		},
	}

	m := make(map[int][]*TeatMap)
	for _, teatMap := range list {
		m[teatMap.Pid] = append(m[teatMap.Pid], teatMap)
	}

	fmt.Println(m)
}

func TestGB(t *testing.T) {
	var m map[string]string

	fmt.Println(len(m))
	fmt.Println((m != nil && m["3"] == "3"))
}
