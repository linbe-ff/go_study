package sensitive_word

import (
	"fmt"
	"testing"
)

// 简单测试
func TestDFAFilter(t *testing.T) {
	dfa := NewDFA()
	// 添加敏感词
	dfa.AddWord("色情")
	dfa.AddWord("夜总会")
	dfa.AddWord("政治")

	// 输入文本
	text := "今夜总会想起你夜总"
	fmt.Println(dfa.Filter(text))
	err := dfa.Check(text)
	if err == nil {
		t.Error("期望得到错误，但实际上没有收到")
	} else {
		fmt.Println("捕获到错误:", err)
	}
}

func TestDFAFilterAll(t *testing.T) {
	dfa := NewDFA()
	// 添加敏感词
	dfa.AddWord("色情")
	dfa.AddWord("夜总会")
	dfa.AddWord("政治")

	// 输入文本
	text := "今夜总会想起你夜总"
	fmt.Println(dfa.Filter(text))
	err := dfa.Check(text)
	if err == nil {
		t.Error("期望得到错误，但实际上没有收到")
	} else {
		fmt.Println("捕获到错误:", err)
	}
}
