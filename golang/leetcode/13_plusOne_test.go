package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlusOne(t *testing.T) {
	data := []struct {
		data   []int
		expect []int
	}{
		{data: []int{9}, expect: []int{1, 0}},
		{data: []int{1, 2, 3}, expect: []int{1, 2, 4}},
		{data: []int{4, 3, 2, 1}, expect: []int{4, 3, 2, 2}},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v", td.expect), func(t *testing.T) {
			result := PlusOne(td.data)
			require.Equalf(t, td.expect, result, "🚫: e=%v r=%v", td.expect, result)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func PlusOne(nums []int) []int {
	return nums
}
