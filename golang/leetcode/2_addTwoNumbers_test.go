package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTwoNumbers(t *testing.T) {
	data := []struct {
		data   [2]*ListNode
		expect *ListNode
	}{
		{data: [2]*ListNode{buildList([]int{2, 2, 5}), buildList([]int{2, 2, 9})}, expect: buildList([]int{4, 4, 4, 1})},
		{data: [2]*ListNode{buildList([]int{1, 2, 3}), buildList([]int{1, 2, 3})}, expect: buildList([]int{2, 4, 6})},
	}

	for _, td := range data {
		if t.Run("📌:"+td.expect.string(), func(t *testing.T) {
			result := AddTwoNumbers(td.data[0], td.data[1])
			require.Equal(t, td.expect.toSlice(), result.toSlice(), "🚫: "+result.string())
		}) {
			t.Log("✅:" + td.expect.string())
		}
	}
}

// 📝 Даны 2 числа в виде связанных списков, вернуть сумму как список.
func AddTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	head = new(ListNode)
	carr := 0

	for curr := head; l1 != nil || l2 != nil || carr > 0; curr = curr.Next {
		if l1 != nil {
			carr += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carr += l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{Val: carr % 10}
		carr /= 10
	}

	return head.Next
}
