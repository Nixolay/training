package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	data := []struct {
		data   string
		expect bool
	}{
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([]))", false},
		{"([)]", false},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v %s", td.expect, td.data), func(t *testing.T) {
			result := IsValidParenthesesMain2(td.data)
			require.Equalf(t, td.expect, result, "🚫: %v, %s", td.expect, td.data)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func IsValidParenthesesMain2(str string) bool {
	return true
}
