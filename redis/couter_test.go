package redis

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	var (
		r = &Redis{}
	)

	r.Conn()
	defer r.Close()

	r.Del("志杰" + LIKE)
	r.Del("泽楷" + LIKE)
	r.Del("观华" + LIKE)
	r.Del("志杰" + SUPPORT)
	r.Del("泽楷" + SUPPORT)
	r.Del("观华" + SUPPORT)
	r.Del("志杰" + COLLECTION)
	r.Del("泽楷" + COLLECTION)
	r.Del("观华" + COLLECTION)

	r.IncrementCount("志杰", LIKE)
	r.IncrementCount("志杰", LIKE)
	r.IncrementCount("志杰", LIKE)
	r.IncrementCount("志杰", LIKE)
	r.IncrementCount("泽楷", SUPPORT)
	r.IncrementCount("观华", COLLECTION)
	r.IncrementCount("观华", COLLECTION)
	r.IncrementCount("观华", COLLECTION)

	r.DecrementCount("志杰", LIKE)

	zj, _ := r.GetCount("志杰", LIKE)
	zk, _ := r.GetCount("泽楷", SUPPORT)
	gh, _ := r.GetCount("观华", COLLECTION)
	fmt.Println(zj)
	fmt.Println(zk)
	fmt.Println(gh)
}
