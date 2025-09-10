package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodeStringify(node *ListNode) (str string) {
	for n := node; n != nil; n = n.Next {
		str += fmt.Sprint(n.Val)
	}
	return
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func (l *ListNode) toSlice() []int {
	res := []int{}
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func (l *ListNode) string() (str string) {
	for n := l; n != nil; n = n.Next {
		str += fmt.Sprint(n.Val)
	}
	return
}

// Вспомогательная функция для вычисления абсолютного значения
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
