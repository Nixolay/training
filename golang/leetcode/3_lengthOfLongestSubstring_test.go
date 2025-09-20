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
	lft := 0

	for rgt, rn := range str {
		if idx, ok := match[rn]; ok {
			lft = max(lft, idx+1)
		}

		match[rn] = rgt
		maxLen = max(maxLen, rgt-lft+1)
	}

	return
}
