package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"sync"
	"time"
)

const (
	LIKE       = "like"       // 喜欢
	COLLECTION = "collection" // 点赞
	SUPPORT    = "support"    // 关注
)

type Redis struct {
	Rdb      *redis.Client
	Ctx      context.Context
	TmpValue interface{}
}

func (r *Redis) Conn() {
	r.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 密码，没有则留空
		DB:       2,                // 使用默认DB
	})
	r.Ctx = context.TODO()
}

func (r *Redis) Close() {
	r.Rdb.Close()
}

func (r *Redis) ZAdd(score float64, member string) error {
	if r.IsNil() {
		return errors.New("redis is nil")
	}
	add := r.Rdb.ZAdd(r.Ctx, "ranking", &redis.Z{Score: score, Member: member})
	_, err := add.Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) ZDel(playerName string) error {
	if r.IsNil() {
		return errors.New("redis is nil")
	}
	_, err := r.Rdb.ZRem(r.Ctx, "ranking", playerName).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) IsNil() bool {
	return r.Rdb == nil
}

func (r *Redis) IncrementCount(rkey, rtype string) error {
	var err error
	r.TmpValue, err = r.Rdb.IncrBy(r.Ctx, rkey+rtype, 1).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) DecrementCount(rkey, rtype string) error {
	_, err := r.Rdb.DecrBy(r.Ctx, rkey+rtype, 1).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) GetCount(rkey, rtype string) (int64, error) {
	// 使用GET命令获取指定键的值
	value, err := r.Rdb.Get(r.Ctx, rkey+rtype).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}

	// 如果键不存在或值为空字符串，返回0
	if value == "" {
		return 0, nil
	}

	// 将字符串值转换为整数
	likeCount, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}

	return likeCount, nil
}

func (r *Redis) Del(key string) error {
	if r.IsNil() {
		return errors.New("redis is nil")
	}
	_, err := r.Rdb.Del(r.Ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

// acquireLock 尝试获取锁
func (r *Redis) AcquireLock(lockKey string, total int, expireTime time.Duration) error {

	//val, err := r.rdb.Get(r.ctx, lockKey).Result()

	// 使用SETNX命令尝试设置键值对，如果键不存在
	result, err := r.Rdb.SetNX(r.Ctx, lockKey, total, expireTime).Result()
	if err != nil {
		return errors.New("Failed to acquire lock： " + err.Error())
	}
	if !result {
		return errors.New("占用中请稍等")
	}

	return nil
}

// releaseLock 释放锁
func (r *Redis) ReleaseLock(lockKey string, lockValue interface{}) error {

	luaScript := `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end;
`
	// 使用Lua脚本释放锁
	_, err := r.Rdb.Eval(context.Background(), luaScript, []string{lockKey}, lockValue).Result()
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("release lock failed ... ")
		// 处理错误
	}
	return nil
}

func (r *Redis) CreateQueue() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	var listName = "seckill"
	var total int64 = 20

	userIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	count := len(userIDs)
	ch := make(chan int, count)

	f := func(id int) {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		lLen, _ := r.Rdb.LLen(r.Ctx, listName).Result()
		if lLen < total {
			r.Rdb.RPush(r.Ctx, listName, fmt.Sprintf("%d@%v", id, time.Now()))
			fmt.Println(id, "抢购成功")
		} else {
			fmt.Println("抢购活动已结束")
		}
		ch <- id
	}

	wg.Add(count)
	for _, v := range userIDs {
		go f(v)
	}

	for i := 0; i < count; i++ {
		<-ch
	}
	close(ch)
	wg.Wait()
}

func (r *Redis) SecKill() {

}
