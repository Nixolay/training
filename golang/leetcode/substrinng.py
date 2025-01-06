# Для заданной строки s найдите длину самой длинной подстрока без повторяющихся символов.

# Пример 1:
# Ввод: s = "abcabcbb"
# Вывод:3
# Объяснение: Ответ "abc" длиной 3.
# Пример 2:

# Ввод: s = "bbbbb"
# Вывод:1
# Объяснение: Ответ "b" длиной 1.
# Пример 3:

# Ввод: s = "pwwkew"
# Вывод:3
# Пояснение: Ответ "wke" длиной 3.
# Обратите внимание, что ответ должен быть подстрокой, "pwke" - это подпоследовательность, а не подстрока.
# O(n^2)

def longest_substring(s: str) -> int:
    if not s:
        raise ValueError("s is empty")

    m = 0
    start = 0
    seen = {}

    for end in range(len(s)):
        c = s[end]
        if c in seen and start <= seen[c]:
            start = seen[c] + 1
        elif m < end - start + 1:
            m =  end - start + 1

        seen[c] = end

    return m


print(longest_substring('abcabcbb'))
print(longest_substring('bbbbb'))
print(longest_substring('pwwkew'))
print(longest_substring('dvdf'))