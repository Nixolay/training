package fractions

import (
	"fmt"
	"strconv"
	"strings"
)

func Decompose(s string) []string {
	out := make([]string, 0)

	if s == "0" {
		return out
	}

	if !strings.Contains(s, "/") && !strings.Contains(s, ".") {
		return []string{s}
	}

	if strings.Contains(s, "/") {
		data := strings.Split(s, "/")

		a, err := strconv.Atoi(data[0])
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}

		if !strings.Contains(fmt.Sprint(a/b), ".") {
			return []string{fmt.Sprint(a / b)}
		}
	}

	return out
}
