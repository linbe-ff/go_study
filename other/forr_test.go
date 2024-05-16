package other

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	uuid2 "github.com/google/uuid"
	"github.com/modern-go/reflect2"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

type Main struct {
	Name    string
	SonList *SonList
}

type (
	Son struct {
		Name string
	}
	SonList []*Son
)

func TestForr(t *testing.T) {
	sonList := make(SonList, 0)
	sonList = append(sonList, &Son{Name: "son1"})
	sonList = append(sonList, &Son{Name: "son2"})
	sonList = append(sonList, &Son{Name: "son3"})
	sonList = append(sonList, &Son{Name: "son4"})
	m := &Main{
		Name:    "main",
		SonList: &sonList,
	}
	for _, son := range *m.SonList {
		log.Println(son)
	}
}

// 將A~Z排列
func TestArrangeA2Z(t *testing.T) {
	var (
		prevMark = 'A' - 1
		mark     = 'A'
	)
	for i := 0; i < 100; i++ {
		if mark > 'Z' {
			prevMark++
			mark = 'A'
		}
		if prevMark == 'A'-1 {
			fmt.Println(fmt.Sprintf("%c", mark))
		} else {
			fmt.Println(fmt.Sprintf("%c%c", prevMark, mark))
		}

		mark++
	}
}

func TestBeforeTime(t *testing.T) {
	var maxTime time.Time
	now := time.Now()
	if maxTime.Before(now) {
		maxTime = now
	}
	fmt.Println(maxTime)
}
func Test1(t *testing.T) {
	count := 2
	if count != 2 && count != 3 {
		fmt.Println("count is not 2 or 3")
	}
}

var key bool = false

const (
	//url = "http://192.168.1.20:3001/procurementManaging/blanketManagement"
	url           = "http://192.168.1.127:50002/hcscm/admin/v1/info_basic_data/infoBasicDefect/list"
	authorization = "8152a0029a3c7857e7324b97ab1a2ef075671b9bb5fe4de99ac1874db4b45ef6310a10bf08b3e580ab4e6ffcc6e4e60e5398ead2659cd230d4142871ea5d0cff"
	platform      = "1"
	thread        = 12
	times         = 100
)

func fastGet(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// 设置 Authorization 头部
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Platform", platform)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func worker(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	var sum int
	for i := 0; i < times; i++ {
		if key {
			break
		}
		resp, err := fastGet(url) // 忽略错误处理，实际应用中应添加
		if err != nil {
			return
		}
		resp.Body.Close() // 不要忘记关闭响应体
		sum++
	}
	key = true
	ch <- sum
}

func TestPerFast(t *testing.T) {
	c := make(chan int)
	var wg sync.WaitGroup

	start := time.Now().UnixMilli()

	for i := 0; i < thread; i++ {
		wg.Add(1)
		go worker(&wg, c)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	total := 0
	for num := range c {
		total += num
	}

	end := time.Now().UnixMilli()
	diff := end - start
	log.Printf("总耗时: %f seconds", float64(diff)/1000.0)

	log.Printf("请求总数: %d", total)
	log.Printf("QPS: %f", float64(total)/float64(diff)*1000.0)
}

func worker2(c context.Context, wg *sync.WaitGroup, id int) {

	defer wg.Done()

	select {
	case <-time.After(time.Second):
		fmt.Sprintln("执行完成 %", id)
	case <-c.Done():
		fmt.Sprintln("请求超时 %", id)
	}
}

func Test1000(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go worker2(ctx, &wg, i)
	}

	wg.Wait()
}

func TestGroupMp(t *testing.T) {
	var (
		syncMap sync.Map
		count   int
		m       = make(map[int]int)
	)
	m[1] = 1
	fmt.Println(m)
	for i := 1; i < 1000; i++ {
		go func() {
			syncMap.Store(i, i)
			count++
		}()
	}
	time.Sleep(time.Second * 2)
	//syncMap.Range(func(key, value interface{}) bool {
	//	fmt.Println(key, value)
	//	return true
	//})
	fmt.Println(count)

}

func addMapValue(m sync.Map, k int) {
	m.Store(k, k)
}

func TestMapPointer(t *testing.T) {
	var (
		m = make(map[int]int, 10)
	)
	m[1] = 1
	nowPointer := reflect2.PtrOf(m[1])
	nowPointer2 := reflect2.PtrOf(m[1])
	fmt.Println(nowPointer == nowPointer2)
	fmt.Println(nowPointer, nowPointer2)
	for i := 2; i < 1000000; i++ {
		m[i] = i
		if nowPointer != reflect2.PtrOf(m[1]) {
			fmt.Println(nowPointer)
			fmt.Println(reflect2.PtrOf(m[1]))
			fmt.Println("nowPointer != &m  cap = ", i)
			return
		}
	}

}

func TestSlicePointer(t *testing.T) {
	var (
		m          = make([]int, 10)
		nowPointer = reflect2.PtrOf(m)
	)
	//fmt.Println(nowPointer)
	for i := 1; i < 1000000; i++ {
		m = append(m, i)
		if nowPointer != reflect2.PtrOf(m) {
			fmt.Println(nowPointer)
			fmt.Println(reflect2.PtrOf(m))
			fmt.Println("nowPointer != &m  cap = ", i)
			return
		}
	}

}

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

func TestChan(t *testing.T) {
	var (
		ch = make(chan int, 2)
	)
	go func() {
		for {
			select {
			case n := <-ch:
				time.Sleep(time.Second)
				fmt.Println(n)
			}
		}
	}()
	for i := 1; i < 11; i++ {
		select {
		case ch <- i:
		default:
			fmt.Println("丢弃请求")
			time.Sleep(time.Second)
		}
	}
	time.Sleep(10 * time.Second)
}

func TestEEE(t *testing.T) {
	old := []string{"a", "b", "c"}
	new1 := []string{"a", "d", "e", "c"}
	result := make([]string, 0)

	oldMap := make(map[string]bool)
	for _, v := range old {
		oldMap[v] = true
	}

	// 去掉old中 new没有的元素
	for _, v := range new1 {
		if oldMap[v] {
			result = append(result, v)
		}
	}

	// 创建一个新的切片，仅包含old中也存在于new1中的元素

	fmt.Println(result)

}
