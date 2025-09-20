package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	data := []struct {
		data   []int
		target int
		expect []int
	}{
		{data: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, target: 10, expect: []int{1, 3}},
		{data: []int{1, 1}, target: 2, expect: []int{0, 1}},
		{[]int{1, 2, 2, 4, 9, 10}, 4, []int{1, 2}},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v", td.expect), func(t *testing.T) {
			t.Parallel()
			result := TwoSum(td.data, td.target)
			require.Equalf(t, td.expect, result, "🚫 actual:%v, target:%d expect:%v", result, td.target, td.expect)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

// 📝 Дан массив и target, найти два индекса чисел, которые дают target.
func TwoSum(nums []int, target int) []int {
	match := map[int]int{}

	for idx, num := range nums {
		if i, ok := match[target-num]; ok {
			return []int{i, idx}
		}

		match[num] = idx
	}

	return nil
}
