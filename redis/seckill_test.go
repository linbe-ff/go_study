package redis

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var (
	userIds = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		2, 3, 4, 5, 6, 7, 8, 9, 102, 13, 14, 15, 16, 17, 18, 19, 110}
	listName    = "seckill"
	exprireTime = 200 * time.Millisecond
)

// chan 模式实现
func TestSecKill(t *testing.T) {
	var (
		r                = &Redis{}
		ch               = make(chan int, 200)
		lLen       int64 = 0
		stockCount int64 = 5 // 模拟库存已经写在redis的库存数
	)
	r.Conn()
	defer func() {
		r.Close()
	}()

	go func() {
		for {
			select {
			case userId := <-ch:
				// 假设`uniqueUsers`是存储已抢购用户ID的集合键
				userAlreadyPurchased, _ := r.Rdb.SIsMember(r.Ctx, "uniqueUsers", userId).Result()

				if !userAlreadyPurchased {
					lLen, _ = r.Rdb.LLen(r.Ctx, listName).Result()
					if lLen < stockCount {
						tmpStr := fmt.Sprintf("%d@%v", userId, time.Now())
						r.Rdb.RPush(r.Ctx, listName, tmpStr)
						fmt.Println("抢购成功", tmpStr)
						// 添加用户ID到集合
						r.Rdb.SAdd(r.Ctx, "uniqueUsers", userId)
					} else {
						fmt.Println("抢购结束")
					}
				} else {
					fmt.Println("用户已经参与")
				}
			}
		}
	}()

	// 模拟请求打进来
	for _, userId := range userIds {
		exists, _ := r.Rdb.Exists(r.Ctx, listName+":"+strconv.Itoa(userId)).Result()
		if exists > 0 {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("请求频繁")
			continue
		}
		r.Rdb.Set(r.Ctx, listName+strconv.Itoa(userId), "", exprireTime)
		ch <- userId
	}
	time.Sleep(30 * time.Second)
}

func TestSecKill2(t *testing.T) {
	var (
		r  = &Redis{}
		ch = make(chan int, 1000)
	)
	r.Conn()
	defer func() {
		r.Close()
	}()

	r.Rdb.Set(r.Ctx, listName+":"+"stock", 1000, 0)

	for i := 0; i < 100000; i++ {
		userIds = append(userIds, i+10)
	}

	go func() {
		for {
			select {
			case userId := <-ch:
				// 假设`uniqueUsers`是存储已抢购用户ID的集合键
				userAlreadyPurchased, _ := r.Rdb.SIsMember(r.Ctx, "uniqueUsers", userId).Result()
				if !userAlreadyPurchased {
					//lLen, _ = r.rdb.LLen(r.ctx, listName).Result()
					result, _ := r.Rdb.Get(r.Ctx, listName+":"+"stock").Result()

					atoi, _ := strconv.Atoi(result)
					if atoi == 0 {
						fmt.Println("抢购结束")
						return
					}
					if atoi > 0 {
						r.Rdb.DecrBy(r.Ctx, listName+":"+"stock", 1).Result()
						tmpStr := fmt.Sprintf("%d@%v", userId, time.Now())
						r.Rdb.RPush(r.Ctx, listName, tmpStr)
						//fmt.Println("抢购成功", tmpStr)
						// 添加用户ID到集合
						r.Rdb.SAdd(r.Ctx, "uniqueUsers", userId)
					} else {
						//fmt.Println("抢购结束")
					}
				} else {
					//fmt.Println("用户已参与")
				}
			}
		}
	}()

	// 模拟请求打进来
	go func() {
		for _, userId := range userIds {
			exists, _ := r.Rdb.Exists(r.Ctx, listName+":"+strconv.Itoa(userId)).Result()
			if exists > 0 {
				//time.Sleep(200 * time.Millisecond)
				fmt.Println("请求频繁")
				continue
			}
			select {
			case ch <- userId:
				r.Rdb.Set(r.Ctx, listName+strconv.Itoa(userId), "", exprireTime)
			default:
				//fmt.Println("拥挤中")
				continue
			}
		}
	}()
	go func() {
		for _, userId := range userIds {
			exists, _ := r.Rdb.Exists(r.Ctx, listName+":"+strconv.Itoa(userId)).Result()
			if exists > 0 {
				//time.Sleep(200 * time.Millisecond)
				fmt.Println("请求频繁")
				continue
			}
			select {
			case ch <- userId:
				r.Rdb.Set(r.Ctx, listName+strconv.Itoa(userId), "", exprireTime)
			default:
				//fmt.Println("拥挤中")
				continue
			}
		}
	}()
	go func() {
		for _, userId := range userIds {
			exists, _ := r.Rdb.Exists(r.Ctx, listName+":"+strconv.Itoa(userId)).Result()
			if exists > 0 {
				//time.Sleep(200 * time.Millisecond)
				fmt.Println("请求频繁")
				continue
			}
			select {
			case ch <- userId:
				r.Rdb.Set(r.Ctx, listName+strconv.Itoa(userId), "", exprireTime)
			default:
				//fmt.Println("拥挤中")
				continue
			}
		}
	}()
	time.Sleep(3 * time.Second)

}

func TestDeleteKey(t *testing.T) {
	r := Redis{}
	r.Conn()
	defer r.Close()
	r.Rdb.Del(r.Ctx, "uniqueUsers")
	r.Rdb.Del(r.Ctx, listName)
	r.Rdb.Del(r.Ctx, "users")
	r.Rdb.Del(r.Ctx, listName+":"+"stock")
}
