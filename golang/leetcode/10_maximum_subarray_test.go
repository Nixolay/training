package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumSubArrayMain2(t *testing.T) {
	data := []struct {
		data   []int
		expect int
	}{
		{data: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expect: 6},
		{data: []int{1}, expect: 1},
		{data: []int{5, 4, -1, 7, 8}, expect: 23},
	}

	for _, td := range data {
		if t.Run("📌: "+strconv.Itoa(td.expect), func(t *testing.T) {
			result := MaximumSubArrayMain2(td.data)
			require.Equal(t, td.expect, result, "🚫: "+strconv.Itoa(result))
		}) {
			t.Log("✅: " + strconv.Itoa(td.expect))
		}
	}
}

// 📝 Найти подмассив с максимальной суммой.
func MaximumSubArrayMain2(nums []int) (sum int) {
	sum, cur := nums[0], nums[0]

	for _, num := range nums[1:] {
		cur = max(num, num+cur)
		sum = max(sum, cur)
	}

	return
}
