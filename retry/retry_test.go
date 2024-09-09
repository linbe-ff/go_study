package retry

import (
	"errors"
	"fmt"
	"github.com/cenkalti/backoff"
	"strconv"
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

func TestRetry2(t *testing.T) {
	b := backoff.NewExponentialBackOff()
	bf := backoff.WithMaxRetries(b, 2)
	retryCount := 0
	for i := 0; i < 3; i++ {
		retryCount = 0
		backoff.Retry(func() error {
			err := errors.New("我是错误" + strconv.Itoa(i))
			if err != nil {
				retryCount++
				fmt.Println(err.Error(), retryCount)
			}
			return err
		}, bf)
	}
}
