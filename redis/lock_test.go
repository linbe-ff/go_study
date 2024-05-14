package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"go_study/retry"
	"testing"
	"time"
)

var (
	lockKey     = "diaomao"
	distributed = "distributed"
	lockVal     = 121
	timeout     = 25 * time.Hour
)

func TestRedis(t *testing.T) {
	var (
		r = &Redis{}
	)

	r.Conn()
	defer func() {
		r.Close()
	}()

	err := r.AcquireLock(lockKey, lockVal, time.Second*60)
	fmt.Println("AcquireLock ", err)
	err = r.ReleaseLock(lockKey, lockVal)
	fmt.Println("ReleaseLock ", err)
}

func TestLockRetry(t *testing.T) {
	var (
		r = &Redis{}
	)

	r.Conn()
	defer func() {
		r.ReleaseLock(lockKey, lockVal)
		r.Close()
	}()

	retry := retry.NewRetry(5, 5*time.Second, 30*time.Second, func(attempt int) time.Duration {
		return 1 * time.Second
	})

	err := r.AcquireLock(lockKey, 121, time.Second*20)
	if err != nil {
		err = retry.Retry(func() error {
			return r.AcquireLock(lockKey, 121, time.Second*20)
		})
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}

// 设置常用键值 + 分布式锁
func TestAddOneSetNX(t *testing.T) {
	var (
		r      = &Redis{}
		incVal int64
		today  = "20240515"
	)
	lockKey = lockKey + today

	r.Conn()
	defer func() {
		r.ReleaseLock(distributed+lockKey, incVal)
		r.Close()
	}()

	_, err := r.rdb.Get(r.ctx, lockKey).Result()
	if err == redis.Nil {
		_, err = r.rdb.Set(r.ctx, lockKey, lockVal, timeout).Result()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	incVal, err = r.rdb.Incr(r.ctx, lockKey).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = r.AcquireLock(distributed+lockKey, int(incVal), timeout)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
