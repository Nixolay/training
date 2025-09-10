package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindrome(t *testing.T) {
	data := []struct {
		data   string
		result bool
	}{
		{"amanaplanacanalpanama", true},
		{"raceacar", false},
		{"", true},
		{" ", true},
		{"bar", false},
	}

	for _, td := range data {
		result := IsPalindrome(td.data)
		require.Equalf(t, td.result, result, "🚫: %s", td.data)
		println("✅", td.data, result)
	}
}

func IsPalindrome(s string) bool {
	for l, r := 0, len(s); l < r; l++ {
		r--

		if s[l] != s[r] {
			return false
		}
	}

	return true
}
