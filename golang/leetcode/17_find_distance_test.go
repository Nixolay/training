package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindDistance(t *testing.T) {
	data := []struct {
		data   [2][]int
		d      int
		expect int
	}{
		{data: [2][]int{{4, 5, 8}, {10, 9, 1, 8}}, d: 2, expect: 2},
		{data: [2][]int{{1, 4, 2, 3}, {-4, -3, 6, 10, 20, 30}}, d: 3, expect: 2},
		{data: [2][]int{{2, 1, 100, 3}, {-5, -2, 10, -3, 7}}, d: 6, expect: 1},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v", td.data), func(t *testing.T) {
			result := FindDistance(td.data[0], td.data[1], td.d)
			require.Equalf(t, td.expect, result, "🚫: e=%v r=%v", td.expect, result)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func FindDistance(arr1 []int, arr2 []int, d int) (count int) {
	for _, num := range arr1 {
		valid := true

		for _, num2 := range arr2 {
			if abs(num-num2) <= d {
				valid = !valid
				break
			}
		}

		if valid {
			count++
		}
	}

	return
}
