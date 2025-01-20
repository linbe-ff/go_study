package sensitive_word

import (
	"errors"
)

// TrieNode 定义 DFA 的节点
type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

// NewTrieNode 创建一个新的 Trie 节点
func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		IsEnd:    false,
	}
}

// DFA 结构体包含 DFA 的根节点
type DFA struct {
	Root *TrieNode
}

// NewDFA 初始化一个新的 DFA
func NewDFA() *DFA {
	return &DFA{
		Root: NewTrieNode(),
	}
}

// AddWord 将敏感词添加到 DFA 中
func (d *DFA) AddWord(word string) {
	node := d.Root
	for _, char := range word {
		if _, exists := node.Children[char]; !exists {
			node.Children[char] = NewTrieNode()
		}
		node = node.Children[char]
	}
	node.IsEnd = true
}

// Filter 过滤输入的文本并替换敏感词
func (d *DFA) Filter(text string) string {
	result := []rune(text)
	for i := 0; i < len(result); i++ {
		node := d.Root
		j := i
		for j < len(result) {
			if nextNode, exists := node.Children[result[j]]; exists {
				node = nextNode
				if node.IsEnd {
					for k := i; k <= j; k++ {
						result[k] = '*'
					}
				}
				j++
			} else {
				break
			}
		}
	}
	return string(result)
}

// Check 检查输入的文本是否包含敏感词
func (d *DFA) Check(text string) error {
	result := []rune(text)
	for i := 0; i < len(result); {
		node := d.Root
		start := i
		matched := false
		for j := i; j < len(result); j++ {
			char := result[j]
			if nextNode, exists := node.Children[char]; exists {
				node = nextNode
				if node.IsEnd {
					matched = true
					return errors.New("包含敏感词: " + string(result[start:j+1]))
				}
			} else {
				break
			}
		}
		if !matched {
			i++
		}
	}
	return nil
}
