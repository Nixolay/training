// Package romandecoder ...
package romandecoder

//nolint:gomnd
// Decode ...
func Decode(roman string) int {
	println(roman)

	var (
		out       int
		past      int
		romanRune = map[rune]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}
	)

	for i, r := range roman {
		if i == 0 {
			past = romanRune[r]
		}

		if past < romanRune[r] {
			out -= past + past
		}

		out += romanRune[r]
		past = romanRune[r]
	}

	return out
}
