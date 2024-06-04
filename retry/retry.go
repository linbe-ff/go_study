package retry

import (
	"fmt"
	"time"
)

type Retry struct {
	retryCount        int
	retryDelay        time.Duration
	retryTimeout      time.Duration
	retryIntervalFunc func(int) time.Duration
}

func NewRetry(
	retryCount int,
	retryDelay time.Duration,
	retryTimeout time.Duration,
	retryIntervalFunc func(attempt int) time.Duration,
) *Retry {
	return &Retry{
		retryCount:        retryCount,
		retryDelay:        retryDelay,
		retryTimeout:      retryTimeout,
		retryIntervalFunc: retryIntervalFunc,
	}
}

// 重试方法
func (r *Retry) Retry(fn func() error) error {
	retryStart := time.Now()
	var err error
	retryCount := 0

	for {
		fmt.Println("retryCount:  ", retryCount+1)

		time.Sleep(r.retryDelay)

		err = fn()

		if err == nil {
			return nil
		}

		if err != nil {
			if retryCount < r.retryCount-1 {
				retryCount++
				continue
			}
			break
		}

		if retryStart.Add(r.retryTimeout).Before(time.Now()) {
			break
		}
	}
	return err
}
