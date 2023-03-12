package dec2fac

import (
	"strconv"
	"strings"
)

func FactString2Dec(str string) int {
	response := 0
	step := len(str)
	iter := 0

	for step > 0 {
		rest, _ := strconv.ParseInt(string(str[iter]), 16, 10)
		response = response*step + int(rest)
		step--
		iter++
	}

	return response
}

func Dec2FactString(nb int) string {
	var response string

	step := 1
	remainder := 0
	for nb > 0 {
		nb, remainder = nb/step, nb%step
		response = strconv.FormatInt(int64(remainder), 16) + response
		step++
	}

	return strings.ToUpper(response)
}
