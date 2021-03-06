/*
Package alphanumeric В этом примере вы должны проверить, является ли строка
ввода пользователя буквенно-цифровой. Данная строка не
nil/null/NULL/None, поэтому вам не нужно проверять это.

Строка имеет следующие условия, чтобы быть буквенно-цифровой:

     Хотя бы один символ ("" недействителен)
     Допустимые символы - заглавные / строчные латинские буквы и цифры от 0 до 9
		 Нет пробелов / подчеркивания
*/
package alphanumeric

import (
	"regexp"
)

// Alphanumeric checked.
func Alphanumeric(str string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(str)
}
