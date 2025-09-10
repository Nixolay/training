package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPalindrome(t *testing.T) {
	data := []struct {
		data, expect string
	}{
		{data: "babad", expect: "bab"},
		{data: "cbbd", expect: "bb"},
		{data: "b", expect: "b"},
		{data: "", expect: ""},
	}

	for _, td := range data {
		if t.Run("📌:"+td.expect, func(t *testing.T) {
			result := LongestPalindromeMain2(td.data)
			require.Equal(t, td.expect, result, "🚫: "+result)
		}) {
			t.Log("✅:" + td.expect)
		}
	}
}

// 📝 Найти самую длинную палиндромную подстроку.
// 🔑 Идея: для каждой позиции расширять палиндром по центру.
func LongestPalindromeMain2(str string) string {
	left, maxLen := 0, 0

	expand := func(l, r int) {
		for l >= 0 && r < len(str) && str[l] == str[r] {
			if r-l+1 > maxLen {
				maxLen = r - l + 1
				left = l
			}

			r++
			l--
		}
	}

	for idx := range str {
		expand(idx, idx)
		expand(idx, idx+1)
	}

	return str[left : left+maxLen]
}
