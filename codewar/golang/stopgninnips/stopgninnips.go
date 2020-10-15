/*Package stopgninnips Напишите функцию, которая принимает строку из одного или нескольких
слов и возвращает ту же строку, но со всеми пятью или более буквенными
словами в обратном порядке (точно так же, как имя этого ката).
Переданные строки будут состоять только из букв и пробелов.
Пробелы будут включены только при наличии более одного слова.

Examples:
spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"
spinWords( "This is a test") => returns "This is a test"
spinWords( "This is another test" )=> returns "This is rehtona test"
*/
package stopgninnips

import "strings"

// SpinWords reverse word.
func SpinWords(str string) string {
	words := strings.Split(str, " ")

	const maxLen = 4

	for i, word := range words {
		if len([]rune(word)) > maxLen {
			words[i] = reverse(word)
		}
	}

	return strings.Join(words, " ")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
