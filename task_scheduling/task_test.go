package task_scheduling

import (
	"context"
	"fmt"
	"go.uber.org/goleak"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("start")
	goleak.VerifyTestMain(m)
}

func TestPoller(t *testing.T) {
	producer := NewPoller(5)
	producer.Poll(context.Background())
}
