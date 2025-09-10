package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeIntervals(t *testing.T) {
	data := []struct {
		data   [][]int
		expect [][]int
	}{
		{data: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, expect: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{data: [][]int{{1, 4}, {4, 5}}, expect: [][]int{{1, 5}}},
		{data: [][]int{{4, 7}, {1, 4}}, expect: [][]int{{1, 7}}},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v", td.data), func(t *testing.T) {
			result := MergeIntervals(td.data)
			require.Equalf(t, td.expect, result, "🚫: e=%v r=%v", td.expect, result)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func MergeIntervals(intervals [][]int) (arr [][]int) {
	return
}
