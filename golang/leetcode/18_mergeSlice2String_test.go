package leetcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeSlice2String(t *testing.T) {
	data := []struct {
		data   []int
		result string
	}{
		{[]int{1, 4, 5, 2, 3, 9, 8, 11, 0, 13}, "0-5,8-9,11,13"},
	}
	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v", td.data), func(t *testing.T) {
			result := MergeSlice2String(td.data)
			require.Equalf(t, td.result, result, "🚫: %v", td.data)
		}) {
			t.Logf("✅: %v", td.result)
		}
	}
}

func MergeSlice2String(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	sort.Ints(nums)

	var builder strings.Builder
	builder.WriteString(strconv.Itoa(nums[0]))

	start := 0
	for i, num := range nums[1:] {
		if num-1 == nums[i] {
			continue
		}

		if i-start > 0 {
			builder.WriteString("-")
			builder.WriteString(strconv.Itoa(nums[i]))
		}

		builder.WriteString(",")
		builder.WriteString(strconv.Itoa(num))

		start = i + 1
	}

	return builder.String()
}
