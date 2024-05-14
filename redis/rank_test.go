package redis

import (
	"fmt"
	"testing"
)

// 排行榜
func TestSort(t *testing.T) {
	var (
		r = &Redis{}
	)
	r.Conn()
	r.ZAdd(60, "php")
	r.ZAdd(95, "python")
	r.ZAdd(90, "java")
	r.ZAdd(100, "go")
	r.ZAdd(98, "rust")

	r.ZDel("java")

	//tuples, err := r.rdb.ZRangeWithScores(r.ctx, "ranking", 0, -1).Result() // 正序
	tuples, err := r.rdb.ZRevRangeWithScores(r.ctx, "ranking", 0, -1).Result() // 倒序

	if err != nil {
		panic(err)
	}
	rankingInfos := make([]*RankingInfo, len(tuples))
	for i, tuple := range tuples {
		rankingInfo := &RankingInfo{
			PlayerName: tuple.Member.(string),
			Score:      int(tuple.Score),
		}
		rankingInfos[i] = rankingInfo

		fmt.Printf("playerName: %s, score: %v\n", tuple.Member, tuple.Score)
	}

}

// result

//playerName: go, score: 100
//playerName: rust, score: 98
//playerName: python, score: 95
//playerName: php, score: 60
