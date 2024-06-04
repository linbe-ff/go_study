package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// "github.com/cenkalti/backoff" 这个包也能用
func TestRetry(t *testing.T) {
	r := NewRetry(3, 1*time.Second, 5*time.Minute, func(attempt int) time.Duration {
		return 1 * time.Second
	})

	err := r.Retry(func() error {
		return errors.New("我是错误")
	})

	if err != nil {
		fmt.Printf(err.Error())
	}
}
