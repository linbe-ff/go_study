package scope

import (
	"fmt"
	"testing"
)

var block = "package"

func TestScope1(t *testing.T) {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}

// output
// The block is inner.
// The block is function.
