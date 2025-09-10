package leetcode

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordPattern(t *testing.T) {
	data := []struct {
		data   [2]string
		expect bool
	}{
		{data: [2]string{"abba", "dog cat cat dog"}, expect: true},
		{data: [2]string{"abba", "dog cat cat fish"}, expect: false},
		{data: [2]string{"abba", "dog dog dog dog"}, expect: false},
		{data: [2]string{"", ""}, expect: true},
		{data: [2]string{"aaaa", "dog cat cat dog"}, expect: false},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v", td.data), func(t *testing.T) {
			result := WordPattern(td.data[0], td.data[1])
			require.Equalf(t, td.expect, result, "🚫: e=%v r=%v", td.expect, result)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func WordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}

	p2w := map[rune]string{}
	w2p := map[string]rune{}

	for idx, ch := range pattern {
		if word, ok := p2w[ch]; ok && (word != words[idx] || w2p[words[idx]] != ch) {
			return false
		}

		w2p[words[idx]] = ch
		p2w[ch] = words[idx]
	}

	return true
}
