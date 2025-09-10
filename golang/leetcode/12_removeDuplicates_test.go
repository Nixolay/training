package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	data := []struct {
		data   []int
		expect int
	}{
		{data: []int{1, 1, 2}, expect: 2},                      // [1,2,_]
		{data: []int{1, 2, 3}, expect: 3},                      // [1,2,3]
		{data: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expect: 5}, // 0,1,2,3,4,_,_,_,_,_
		{data: []int{1, 1, 1, 1, 1, 1, 1, 2, 3}, expect: 3},
		{data: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expect: 10},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v", td.expect), func(t *testing.T) {
			result := RemoveDuplicates(td.data)
			require.Equalf(t, td.expect, result, "🚫: e=%v r=%v", td.expect, result)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func RemoveDuplicates(nums []int) (i int) {
	return i + 1
}
