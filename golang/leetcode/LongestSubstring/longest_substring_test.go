package longestsubstring

/*
Для заданной строки s найдите длину самой длинной подстрока без повторяющихся символов.

Пример 1:
Ввод: s = "abcabcbb"
Вывод:3
Объяснение: Ответ "abc" длиной 3.
Пример 2:

Ввод: s = "bbbbb"
Вывод:1
Объяснение: Ответ "b" длиной 1.
Пример 3:

Ввод: s = "pwwkew"
Вывод:3
Пояснение: Ответ "wke" длиной 3.
Обратите внимание, что ответ должен быть подстрокой, "pwke" - это подпоследовательность, а не подстрока.
*/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestSubstring(t *testing.T) {
	t.Parallel()

	type testData struct {
		s        string
		expected int
	}

	data := []testData{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"dvdf", 3},
	}

	for _, v := range data {
		require.Equal(t, v.expected, LongestSubstring(v.s), v.s)
	}
}

func LongestSubstring(str string) int {
	lastSeen := make(map[rune]int)
	start, max := 0, 0

	for end, symbol := range str {
		// Если символ уже встречался и его последний индекс >= начала текущего окна
		if index, found := lastSeen[symbol]; found && index >= start {
			// Обновляем начало окна, чтобы избежать повторяющегося символа
			start = index + 1
		}

		// Обновляем или добавляем текущий символ в словарь с его индексом
		lastSeen[symbol] = end

		// Вычисляем текущую длину уникальной подстроки и обновляем max
		if (end - start + 1) > max {
			max = end - start + 1
		}
	}

	return max
}
