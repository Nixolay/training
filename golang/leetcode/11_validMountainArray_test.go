package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidMountainArray(t *testing.T) {
	data := []struct {
		data   []int
		expect bool
	}{
		{data: []int{2, 1}, expect: false},
		{data: []int{3, 5, 5}, expect: false},
		{data: []int{0, 3, 2, 1}, expect: true},
		{data: []int{1, 7, 9, 5, 4, 1, 2}, expect: false},
		{data: []int{1, 1, 1, 1, 1, 1, 1, 2, 1}, expect: false},
		{data: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expect: false},
		{data: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, expect: false},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v", td.data), func(t *testing.T) {
			result := ValidMountainArray(td.data)
			require.Equalf(t, td.expect, result, "🚫: e=%v r=%v", td.expect, result)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func ValidMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	peak := 0
	for ; peak < len(arr)-1 && arr[peak] < arr[peak+1]; peak++ {
		if arr[peak] == arr[peak+1] {
			return false
		}
	}

	if peak == 0 || peak+1 == len(arr) {
		return false
	}

	for ; peak < len(arr)-1; peak++ {
		if arr[peak] <= arr[peak+1] {
			return false
		}
	}

	return true
}
