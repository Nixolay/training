package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	data := []struct {
		data   string
		result int
	}{
		{"HelloWorld", 5},
		{"abcabcbb", 3},
		{" ", 1},
		{"aa", 1},
		{"aab", 2},
	}

	for _, td := range data {
		result := LengthOfLongestSubstringMain2(td.data)
		require.Equalf(t, td.result, result, "🚫: %s", td.data)
		println("✅", td.data, result)
	}
}

// 📝 Найти длину подстроки без повторяющихся символов.
func LengthOfLongestSubstringMain2(str string) (maxLen int) {
	match := map[rune]int{}
	left := 0

	for right, ch := range str {
		if idx, ok := match[ch]; ok {
			left = max(left, idx+1)
		}

		match[ch] = right
		maxLen = max(maxLen, right-left+1)
	}

	return
}
