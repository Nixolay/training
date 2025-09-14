package leetcode

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompression(t *testing.T) {
	data := []struct {
		data   string
		expect int
	}{
		{"aabbccc", 6},
		{"aaabb", 4},
		{"a", 1},
		{"abbbbbbbbbbbb", 4},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v %s", td.expect, td.data), func(t *testing.T) {
			actual := Compression([]byte(td.data))
			require.Equalf(t, td.expect, actual, "🚫: %v, %s", td.expect, actual)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func Compression(chars []byte) (pos int) {
	left := 0

	for right := 0; right <= len(chars); right++ {
		if right < len(chars) && chars[left] == chars[right] {
			continue
		}

		chars[pos] = chars[left]
		pos++

		if right-left > 1 {
			num := []byte(strconv.Itoa(right - left))
			copy(chars[pos:], num)
			pos += len(num)
		}

		left = right
	}

	return
}
