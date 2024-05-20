package consistency

import (
	"encoding/json"
	"fmt"
	"go_study/mysql"
	loc_redis "go_study/redis"
	"strconv"
	"sync"
	"testing"
	"time"
)

var tableName = "users"

type UsersInfo struct {
	Id         int
	Account    string // 用户名
	Password   string // 密码
	Email      string // 邮箱
	EmployeeId int32  // 员工id
	Type       int32  // 1:pc 2小程序
	Phone      string
	Key        string
}

func TestCh(t *testing.T) {
	r := loc_redis.Redis{}
	wg := &sync.WaitGroup{}
	r.Conn()
	defer r.Close()

	connect, _ := mysql.Connect()
	data := []UsersInfo{}
	connect.Table(tableName).Find(&data)

	ch := make(chan UsersInfo, len(data))
	wg.Add(len(data))
	now := time.Now()
	for i := 0; i < 50; i++ {
		go func() {
			for info := range ch {
				wg.Done()
				marshal, _ := json.Marshal(info)
				//r.Rdb.Set(r.Ctx, tableName+":"+strconv.Itoa(info.Id), marshal, 0)
				r.Rdb.HSet(r.Ctx, tableName+":"+strconv.Itoa(info.Id), marshal, 0)
				r.ZAdd(float64(info.Id), string(marshal))
			}
		}()
	}

	for _, dataMap := range data {
		ch <- dataMap
	}

	//wg.Add(len(data))
	//for _, dataMap := range data {
	//	go func() {
	//		defer wg.Done()
	//		marshal, _ := json.Marshal(dataMap)
	//		r.Rdb.Set(r.Ctx, tableName+":"+strconv.Itoa(dataMap.Id), marshal, 3*time.Minute)
	//	}()
	//}
	wg.Wait()
	now2 := time.Now()

	fmt.Println(now2.Sub(now))
}

func TestMysqlLimit(t *testing.T) {
	begin := time.Now()
	//r := loc_redis.Redis{}
	//r.Conn()
	//defer r.Close()
	//
	//matches, err := r.Rdb.ZRange(r.Ctx, "ranking", 900000, 900010-1).Result()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, match := range matches {
	//	fmt.Println(match)
	//}
	db, _ := mysql.Connect()
	data := []UsersInfo{}
	db.Table("users").Limit(10).Offset(900000).Find(&data)
	fmt.Println(data)
	fmt.Println(time.Now().Sub(begin))
}

func TestRedisLimit(t *testing.T) {
	begin := time.Now()
	r := loc_redis.Redis{}
	r.Conn()
	defer r.Close()

	matches, err := r.Rdb.ZRange(r.Ctx, "ranking", 900000, 900010-1).Result()
	if err != nil {
		panic(err)
	}

	for _, match := range matches {
		fmt.Println(match)
	}
	fmt.Println(time.Now().Sub(begin))
	//db, _ := mysql.Connect()
	//data := []UsersInfo{}
	//db.Table("users").Limit(900000).Offset(10).Find(&data)
	//fmt.Println(data)
}
