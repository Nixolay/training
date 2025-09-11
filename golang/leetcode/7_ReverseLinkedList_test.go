package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseLinkedList(t *testing.T) {
	data := []struct {
		data   *ListNode
		expect *ListNode
	}{
		{data: buildList([]int{1, 2, 3, 4}), expect: buildList([]int{4, 3, 2, 1})},
	}

	for _, td := range data {
		if t.Run("📌: "+td.expect.string(), func(t *testing.T) {
			result := ReverseLinkedList(td.data)
			require.Equal(t, td.expect.string(), result.string(), "🚫: "+result.string())
		}) {
			t.Log("✅: " + td.expect.string())
		}
	}
}

func ReverseLinkedList(head *ListNode) (out *ListNode) {
	for head != nil {
		tmp := head.Next
		head.Next = out
		out = head
		head = tmp
	}
	return
}
