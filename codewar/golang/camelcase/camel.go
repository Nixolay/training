package camel

import (
	"bytes"
	"unicode"
)

const (
	change rune = 32
	upper  rune = 123
	down   rune = 97
)

func ToCamelCase(phrase string) string {
	buf, up := bytes.Buffer{}, false

	for _, letter := range phrase {
		if !unicode.IsLetter(letter) {
			up = true

			continue
		}

		if up && letter >= down && letter <= upper {
			letter -= change
		}

		buf.WriteRune(letter)

		up = false
	}

	return buf.String()
}
