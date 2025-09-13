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
			result := IsValidParentheses(td.data)
			require.Equalf(t, td.expect, result, "🚫: %v, %s", td.expect, td.data)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func IsValidParentheses(str string) bool {
	match := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]rune, 0, len(str))

	for _, rn := range str {
		switch rn {
		case ']', '}', ')':
			if len(stack) == 0 || match[rn] != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, rn)
		}
	}

	return len(stack) == 0
}
