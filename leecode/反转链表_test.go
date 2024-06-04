package leecode

import (
	"fmt"
	"github.com/mohae/deepcopy"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func (l *ListNode) AddSingleLink(val int) {
	newNode := &ListNode{Val: val}
	current := l
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// if over index, add to last
func (l *ListNode) AddSingleLinkAtIndex(index, val int) {
	var (
		count   = 1
		current = l
		newNode = &ListNode{Val: val}
	)

	if index == 0 {
		tmp := deepcopy.Copy(l).(*ListNode)
		l.Val = val
		l.Next = tmp
		return
	}

	for current.Next != nil {
		if count == index {
			tmp := current.Next
			newNode.Next = tmp
			current.Next = newNode
		} else {
			current = current.Next
		}
		count++
	}
	if index > count {
		current.Next = newNode
	}
}

func (l *ListNode) RemoveAtIndex(index int) {
	if index == 0 {
		// 如果要删除的是头节点，直接将头指向下下一个节点
		l.Val = l.Next.Val
		l.Next = l.Next.Next
		return
	}

	prev := l
	current := l.Next
	count := 1 // 已经经过头节点

	// 遍历链表直到找到待删除节点的前一个节点
	for current != nil && count < index {
		prev = current
		current = current.Next
		count++
	}

	// 检查是否找到了有效的节点进行删除
	if current != nil {
		// 删除节点
		prev.Next = current.Next
	} else {
		// 下标超出了链表范围
		fmt.Println("Index out of bounds")
	}
}

func TestAddSingleLinkByIndex(t *testing.T) {
	head := &ListNode{Val: 1}
	head.AddSingleLink(2)
	head.AddSingleLink(3)
	head.AddSingleLink(4)
	head.AddSingleLinkAtIndex(0, 5)
	printList(head)
	head.RemoveAtIndex(2)
	printList(head)
}

func TestAddSingleLink(t *testing.T) {
	head := &ListNode{Val: 1}
	head.AddSingleLink(2)
	head.AddSingleLink(3)
	head.AddSingleLink(4)
	printList(head)
}

func TestReverserLink(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println("before:")
	printList(head)

	reversedList := reverseList(head)
	fmt.Println("after:")
	printList(reversedList)

	// before:
	//	1 2 3 4 5
	// after:
	//	5 4 3 2 1
	//	练习题LeetCode对应编号：206，141，21，19，876
}
