package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsAnagram(t *testing.T) {
	data := []struct {
		data   [2]string
		expect bool
	}{
		{data: [2]string{"anagram", "nagaram"}, expect: true},
		{data: [2]string{"rat", "car"}, expect: false},
		{data: [2]string{"", ""}, expect: true},
		{data: [2]string{"anagram", ""}, expect: false},
	}

	for _, td := range data {
		if t.Run(fmt.Sprintf("📌: %v", td.data), func(t *testing.T) {
			result := IsAnagram(td.data[0], td.data[1])
			require.Equalf(t, td.expect, result, "🚫: e=%v r=%v", td.expect, result)
		}) {
			t.Logf("✅: %v", td.expect)
		}
	}
}

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var chars [26]int
	for idx, ch := range s {
		chars[ch-'a']++
		chars[t[idx]-'a']--
	}

	for _, count := range chars {
		if count != 0 {
			return false
		}
	}

	return true
}

func IsAnagram1(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	match := make(map[byte]int)

	for idx := range s {
		match[s[idx]]++
		match[t[idx]]--
	}

	for _, v := range match {
		if v > 0 {
			return false
		}
	}

	return true
}

func IsAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := make([]int, 26)

	for i := range s {
		chars[s[i]-'a']++
		chars[t[i]-'a']--
	}

	for i := range chars {
		if chars[i] != 0 {
			return false
		}
	}

	return true
}
