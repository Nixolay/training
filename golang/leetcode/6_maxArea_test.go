package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxAreaMain2(t *testing.T) {
	data := []struct {
		data   []int
		expect int
	}{
		{data: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expect: 49},
		{data: []int{1, 1}, expect: 1},
	}

	for _, td := range data {
		if t.Run("📌: "+strconv.Itoa(td.expect), func(t *testing.T) {
			result := MaxArea(td.data)
			require.Equal(t, td.expect, result, "🚫: "+strconv.Itoa(result))
		}) {
			t.Log("✅: " + strconv.Itoa(td.expect))
		}
	}
}

// MaxAreaMain2
// 📝 Найти 2 линии, образующие контейнер с максимальной водой.
// 🔑 Идея: двигать меньшую высоту.
func MaxArea(height []int) (res int) {
	l, r := 0, len(height)-1

	for l < r {
		h := min(height[l], height[r])
		res = max(res, h*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return
}
