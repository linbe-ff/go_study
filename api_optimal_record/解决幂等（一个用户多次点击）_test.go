package api_optimal_record

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	uuid2 "github.com/google/uuid"
	"testing"
	"time"
)

// redis来验证幂等
// 场景（订单、支付等）
func TestRedis(t *testing.T) {
	var (
		err     error
		newStr  string
		newUUID uuid2.UUID
		del     = &redis.IntCmd{}
		ctx     = context.Background()
	)
	newUUID, err = uuid2.NewUUID()
	if err != nil {
		return
	}
	newStr = newUUID.String()

	// 创建Redis客户端连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 密码，没有则留空
		DB:       0,                // 使用默认DB
	})
	// 关闭连接
	defer rdb.Close()

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pong) // Output: PONG
	rdb.Set(ctx, newStr, "", 10*time.Second)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 10)
		del = rdb.Del(ctx, newStr)
		result, err := del.Result()
		if err != nil {
			return
		}
		if result == 1 {
			fmt.Println("del success")
			fmt.Println("向下执行")
			// 向下执行...
		} else {
			fmt.Println("del fail")
			fmt.Println("执行失败（幂等）")
			// 执行失败（幂等）
			return
		}
	}
}
