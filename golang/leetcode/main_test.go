package leetcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// ---------------------- MAIN TESTS ----------------------
func TestMain2(t *testing.T) {
	fmt.Println("✅ 1.  Two Sum:", twoSum([]int{2, 7, 11, 15}, 9))                                                        // [0,1]
	fmt.Println("✅ 2.  Add Two Numbers:", addTwoNumbers(buildList([]int{2, 4, 3}), buildList([]int{5, 6, 4})).toSlice()) // [7,0,8]
	fmt.Println("✅ 3.  Longest Substring:", lengthOfLongestSubstring("abcabcbb"))                                        // 3
	fmt.Println("✅ 4.  Valid Palindrome:", isPalindrome("ascsa"))                                                        // true
	fmt.Println("✅ 5.  Longest Palindrome:", longestPalindrome("babad"))                                                 // "bab" или "aba"
	fmt.Println("✅ 6.  Container With Most Water:", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))                           // 49
	fmt.Println("✅ 7.  Revers Linked List:",
		nodeStringify(reverseLinkedList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}))) // 321
	fmt.Println("✅ 8.  Valid Parentheses:", isValidParentheses("()[]{}"))                                                         // true
	fmt.Println("✅ 9.  String Compression:", StringCompression([]byte("aaabb")))                                                  // a3b2
	fmt.Println("✅ 10. Maximum Subarray:", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))                                     // 6
	fmt.Println("✅ 11. Valid Mountain Array (Горный массив):", validMountainArray([]int{0, 3, 2, 1}))                             // true
	fmt.Println("✅ 12. Remove Duplicates from Sorted Array:", removeDuplicates([]int{1, 1, 2}))                                   // 2
	fmt.Println("✅ 13. Plus One (Плюс один):", plusOne([]int{9, 9, 9}))                                                           // 1 0 0 0
	fmt.Println("✅ 14. Valid Anagram (Анаграмма):", isAnagram("anagram", "nagaram"))                                              // true
	fmt.Println("✅ 15. Word Pattern (Шаблон слов):", wordPattern("abba", "dog cat cat dog"))                                      // true
	fmt.Println("✅ 16. Merge Intervals (Слияние интервалов):", mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))        // [[1 6] [8 10] [15 18]]
	fmt.Println("✅ 17. Find the Distance Value Between Two Arrays:", findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2)) // 2
	fmt.Println("✅ 18. Cвертка списка в диапазоны:", mergeSlice2String([]int{1, 4, 5, 2, 3, 9, 8, 11, 0, 13}))                    // "0-5,8-9,11,13"
}

// ---------------------- 1. Two Sum (Две суммы) ----------------------
// 📌 https://leetcode.com/problems/two-sum/
// 📝 Дан массив и target, найти два индекса чисел, которые дают target.
// ⏰ O(n) — один проход с hash-map.
// 💾 O(n) — храним все числа в hash-map.
// 🔑 Идея: искать "комплемент" target - nums[i] в map.
func twoSum(nums []int, target int) []int {
	match := make(map[int]int)

	for idx, num := range nums {
		if i, ok := match[target-num]; ok {
			return []int{i, idx}
		}

		match[num] = idx
	}

	return nil
}

// ---------------------- 2. Add Two Numbers (Сложение чисел в виде списка) ----------------------
// 📌 https://leetcode.com/problems/add-two-numbers/
// 📝 Даны 2 числа в виде связанных списков, вернуть сумму как список.
// ⏰ O(max(m,n)) — проходим оба списка.
// 💾 O(1) — создаём новый список на результат.
// 🔑 Идея: складываем поразрядно с переносом.
// carry - нести переносить понести вынести
// dummy - манекен
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{} // фиктивная голова
	carry := 0          // перенос при сложении

	for curr := head; l1 != nil || l2 != nil || carry > 0; curr = curr.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{Val: carry % 10}
		carry /= 10
	}

	return head.Next
}

// ---------------------- 3. Longest Substring Without Repeating Characters (Длинная подстрока без повторов) ----------------------
// 📌 https://leetcode.com/problems/longest-substring-without-repeating-characters/
// 📝 Найти длину подстроки без повторяющихся символов.
// ⏰ O(n) — скользящее окно.
// 💾 O(k) — k уникальных символов.
// 🔑 Идея: расширяем окно и двигаем левую границу при повторе.
func lengthOfLongestSubstring(str string) (maxLen int) {
	match := make(map[rune]int)
	left := 0

	for right, ch := range str {
		if idx, found := match[ch]; found {
			left = max(left, idx+1)
		}

		match[ch] = right
		maxLen = max(maxLen, right-left+1)
	}

	return
}

// ---------------------- 4. Valid Palindrome (Проверка палиндрома) ----------------------
// 📌 https://leetcode.com/problems/valid-palindrome/
// 📝 Проверить, является ли строка палиндромом (игнорировать пробелы, регистр и небуквенные символы).
// ⏰ O(n) — два указателя.
// 💾 O(1).
// 🔑 Идея: идём с двух концов, сравниваем только буквы и цифры.
func isPalindrome2(s string) bool {
	for i, j := 0, len(s)-1; i < j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}

	return true
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	isAlphaNum := func(ch byte) bool {
		return (ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9')
	}

	for l < r {
		for l < r && !isAlphaNum(s[l]) {
			l++
		}
		for l < r && !isAlphaNum(s[r]) {
			r--
		}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

// ---------------------- 5. Longest Palindromic Substring (Длиннейшая палиндромная подстрока) ----------------------
// 📌 https://leetcode.com/problems/longest-palindromic-substring/
// 📝 Найти самую длинную палиндромную подстроку.
// ⏰ O(n^2) — расширяем от центра.
// 💾 O(1) — только индексы.
// 🔑 Идея: для каждой позиции расширять палиндром по центру.
func longestPalindrome(str string) string {
	start, maxLen := 0, 0
	expand := func(l, r int) {
		for l >= 0 && r < len(str) && str[l] == str[r] {
			if r-l+1 > maxLen {
				start, maxLen = l, r-l+1
			}

			l--
			r++
		}
	}

	for idx := range str {
		expand(idx, idx)   // нечетный палидром
		expand(idx, idx+1) // четный палидром
	}

	return str[start : start+maxLen]
}

// ---------------------- 6. Container With Most Water (Контейнер с наибольшей площадью) ----------------------
// 📌 https://leetcode.com/problems/container-with-most-water/
// 📝 Найти 2 линии, образующие контейнер с максимальной водой.
// ⏰ O(n) — два указателя.
// 💾 O(1) — только переменные.
// 🔑 Идея: двигать меньшую высоту.
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
func maxArea(height []int) (res int) {
	l, r := 0, len(height)-1

	for l < r {
		h := min(height[l], height[r])
		res = max(res, h*(r-l))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return
}

// ---------------------- 7. Reverse Linked List (Разворот списка) ----------------------
// 📌 https://leetcode.com/problems/reverse-linked-list/
// 📝 Развернуть односвязный список.
// ⏰ O(n).
// 💾 O(1).
// 🔑 Идея: итеративно переставляем ссылки.
func reverseLinkedList(head *ListNode) (out *ListNode) {
	for head != nil {
		tmp := head.Next
		head.Next = out
		out = head
		head = tmp
	}
	return
}

// ---------------------- 8. Valid Parentheses (Правильные скобки) ----------------------
// 📌 https://leetcode.com/problems/valid-parentheses/
// 📝 Проверить, правильно ли расставлены скобки.
// ⏰ O(n) — один проход.
// 💾 O(n) — стек.
// 🔑 Идея: использовать стек для парных скобок.
// "()[]{}","(]","([])","([)]"
func isValidParentheses(str string) bool {
	cls := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0, len(str))

	for _, ch := range str {
		switch ch {
		case '{', '(', '[':
			stack = append(stack, ch)
			continue
		}

		if cls[ch] != stack[len(stack)-1] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

// ---------------------- 9. String Compression (Сжатие строки) ----------------------
// 📌 https://leetcode.com/problems/string-compression/
// 📝 Сжать строку по правилу: "aaabb" -> "a3b2". Вернуть новую длину массива символов.
// ⏰ O(n).
// 💾 O(1).
// 🔑 Идея: два указателя — один читает, другой пишет.
func StringCompression(chars []byte) int {
	write, read := 0, 0
	for read < len(chars) {
		ch := chars[read]
		count := 0
		for read < len(chars) && chars[read] == ch {
			read++
			count++
		}
		chars[write] = ch
		write++
		if count > 1 {
			for _, c := range []byte(strconv.Itoa(count)) {
				chars[write] = c
				write++
			}
		}
	}
	return write
}

// ---------------------- 10. Maximum Subarray (Максимальная сумма подмассива) ----------------------
// 📌 https://leetcode.com/problems/maximum-subarray/
// 📝 Найти подмассив с максимальной суммой.
// ⏰ O(n) — алгоритм Кадане.
// 💾 O(1) — только переменные.
// 🔑 Идея: dp[i] = max(nums[i], nums[i]+dp[i-1]).
func maxSubArray(nums []int) (sum int) {
	sum, cur := nums[0], nums[0]

	for _, num := range nums[1:] {
		cur = max(num, cur+num)
		sum = max(sum, cur)
	}

	return
}

// ---------------------- 11. Valid Mountain Array (Горный массив) ----------------------
// 📌 https://leetcode.com/problems/valid-mountain-array/
// 📝 Проверить, массив: строго вверх до пика, затем строго вниз; пик не первый/последний.
// ⏰ O(n).
// 💾 O(1).
// 🔑 Идея: пройти вверх, затем вниз и убедиться, что прошли обе фазы.
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	l, r := 0, len(arr)-1

	for l < r {
		if arr[l] < arr[l+1] {
			l++
		} else if arr[r-1] > arr[r] {
			r--
		} else {
			break
		}
	}

	return l > 0 && r < len(arr)-1 && l == r
}

// ---------------------- 12. Remove Duplicates from Sorted Array ------------------------
// 📌 https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// 📝 Удалить дубликаты на месте, вернуть новую длину уникальной части.
// ⏰ O(n).
// 💾 O(1).
// 🔑 Идея: два указателя (write читает уникальные).
func removeDuplicates(nums []int) (i int) {
	for _, num := range nums {
		if num != nums[i] {
			i++
			nums[i] = num
		}
	}

	return i + 1
}

// ---------------------- 13. Plus One (Плюс один) --------------------------------------
// 📌 https://leetcode.com/problems/plus-one/
// 📝 Прибавить 1 к числу, представленному массивом цифр.
// ⏰ O(n).
// 💾 O(1) (без учёта возвращаемого массива).
// 🔑 Идея: идти с конца, переноси́м перенос.
func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		carry += digits[i]
		digits[i] = carry % 10
		carry /= 10
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}

// ---------------------- 14. Valid Anagram (Анаграмма) ---------------------------------
// 📌 https://leetcode.com/problems/valid-anagram/
// 📝 Проверить, является ли t анаграммой s.
// ⏰ O(n).
// 💾 O(1)/O(k) — k размер алфавита.
// 🔑 Идея: счётчик символов. 26 символов
// Input: s = "anagram", t = "nagaram" Output: true
// Input: s = "rat",     t = "car"     Output: false
func isAnagram(s, t string) bool {
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

// ---------------------- 15. Word Pattern (Шаблон слов) --------------------------------
// 📌 https://leetcode.com/problems/word-pattern/
// 📝 Символ шаблона должен соответствовать слову из строки
// ⏰ O(n).
// 💾 O(k) — количество различных символов/слов.
// 🔑 Идея: две мапы: char->word и word->char.
func wordPattern(pattern string, s string) bool {
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

// ---------------------- 16. Merge Intervals (Слияние интервалов) ----------------------
// 📌 https://leetcode.com/problems/merge-intervals/
// 📝 Слить пересекающиеся интервалы в один, не пересекающие оставить
// ⏰ O(n log n) — сортировка.
// 💾 O(n). — зависит от реализации.
// 🔑 Идея: отсортировать по началу, затем расширять текущий интервал.
func mergeIntervals(intervals [][]int) (arr [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	arr = append(arr, intervals[0])
	for _, item := range intervals[1:] {
		if item[0] > arr[len(arr)-1][1] {
			arr = append(arr, item)
			continue
		}

		if item[1] > arr[len(arr)-1][1] {
			arr[len(arr)-1][1] = item[1]
		}
	}

	return
}

// ---------------------- 17. Find the Distance Value Between Two Arrays -----------------
// 📌 https://leetcode.com/problems/find-the-distance-value-between-two-arrays/
// 📝 Количество элементов arr1, у которых |x - y| > d для всех y из arr2.
// ⏰ O(n log m) — сортируем arr2 и бинпоиск.
// 💾 O(1) доп.памяти (не считая сортировки).
// 🔑 Идея: для каждого x ищем ближайших соседей в arr2 и проверяем дистанцию.
func findTheDistanceValue(arr1 []int, arr2 []int, d int) (count int) {
	// sort.Ints(arr2)

	// Проверяем каждый элемент из первого массива
	for _, num1 := range arr1 {
		valid := true
		// Для каждого элемента из arr1 проверяем все элементы из arr2
		for _, num2 := range arr2 {
			// Если расстояние <= d, то элемент не удовлетворяет условию
			if abs(num1-num2) <= d {
				valid = false
				break // Прекращаем проверку для текущего num1
			}
		}
		// Если ни один элемент из arr2 не находится ближе чем d, увеличиваем счетчик
		if valid {
			count++
		}
	}

	return
}

// ---------------------- 18. Cвертка списка в диапазоны
// 📌
// 📝Дан массив целых чисел, повторяющихся элементов в массиве нет.
// 📝Нужно преобразовать в строку, сворачивая соседние по числовому ряду числа в диапазоны.
// ⏰ O(n log n) — сортировка.
// 💾 O(n).
// 🔑 Идея: отсортировать по началу, затем расширять строку
// Input:  [1,4,5,2,3,9,8,11,0,13]
// Output: "0-5,8-9,11,13"
func mergeSlice2String(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	sort.Ints(nums)

	var builder strings.Builder
	builder.WriteString(strconv.Itoa(nums[0]))

	start := 0
	for i, num := range nums[1:] {
		if num-1 == nums[i] {
			continue
		}

		if i-start > 0 {
			builder.WriteString("-")
			builder.WriteString(strconv.Itoa(nums[i]))
		}

		builder.WriteString(",")
		builder.WriteString(strconv.Itoa(num))

		start = i + 1
	}

	return builder.String()
}
