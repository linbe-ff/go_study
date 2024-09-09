package example

import (
	"fmt"
	"testing"
)

type Dog struct {
	Name string
}

type DogList []*Dog

func TestNilStruct(t *testing.T) {
	var dl DogList
	for i, dog := range dl {
		fmt.Println(i, dog)
	}
}
