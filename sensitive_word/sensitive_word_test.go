package sensitive_word

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

// 简单测试
//func TestDFAFilter(t *testing.T) {
//	dfa := NewDFA()
//	// 添加敏感词
//	dfa.AddWord("色情")
//	dfa.AddWord("夜总会")
//	dfa.AddWord("政治")
//
//	// 输入文本
//	text := "今夜总会想起你夜总"
//	fmt.Println(dfa.Filter(text))
//	err := dfa.Check(text)
//	if err == nil {
//		t.Error("期望得到错误，但实际上没有收到")
//	} else {
//		fmt.Println("捕获到错误:", err)
//	}
//}

// 定义结构体来映射 JSON 数据
type Data struct {
	IllegalKeywords []string `json:"illegalKeywords"`
}

var (
	dfa       *DFA
	benchData Data
	text      = strings.Repeat("今夜总会想起你夜总", 10) + "最淫官员"
)

func init() {
	dfa = NewDFA()
	// 读取文件
	file, err := os.ReadFile("illegalWords.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 将 JSON 数据解析到结构体中
	err = json.Unmarshal(file, &benchData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 添加敏感词
	//for _, keyword := range benchData.IllegalKeywords {
	//	dfa.AddWord(keyword)
	//}
}
func init() {
	dfa = NewDFA()
	// 读取文件
	file, err := os.ReadFile("illegalWords.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 将 JSON 数据解析到结构体中
	err = json.Unmarshal(file, &benchData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 添加敏感词
	for _, keyword := range benchData.IllegalKeywords {
		dfa.AddWord(keyword)
	}
}

func BenchmarkDFAFilterAll(b *testing.B) {
	//dfa := NewDFA()
	//// 读取文件
	//file, err := os.ReadFile("illegalWords.json")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//var data Data
	//// 将 JSON 数据解析到结构体中
	//err = json.Unmarshal(file, &data)
	//if err != nil {
	//	fmt.Println("Error parsing JSON:", err)
	//	return
	//}
	//
	//// 添加敏感词
	//for _, keyword := range data.IllegalKeywords {
	//	dfa.AddWord(keyword)
	//}

	// 输入文本
	//text := strings.Repeat("今夜总会想起你夜总最淫官员", 10)
	err := dfa.Check(text)
	if err == nil {
		b.Error("期望得到错误，但实际上没有收到")
	} else {
		//fmt.Println("捕获到错误:", err)
	}

}

func BenchmarkDFAFilterForr(b *testing.B) {
	// 添加敏感词
	for _, keyword := range benchData.IllegalKeywords {
		if strings.Contains(text, keyword) {
			//fmt.Println("存在敏感词", keyword)
			return
		}
	}
}
