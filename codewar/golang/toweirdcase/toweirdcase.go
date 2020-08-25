package toweirdcase

import "fmt"

func toWeirdCase(str string) string {
	byteSlice := make([]byte, len(str))
	upper := true

	for i := range byteSlice {
		char := str[i]

		if char == ' ' {
			upper = true
		} else {
			if upper {
				fmt.Printf("char: %d = %b, 95 = %b", char, char, 95)
				char &= 95
				fmt.Printf(", result: %d = %b\n\n", char, char)
			} else {
				char |= 32
			}
			upper = !upper
		}

		byteSlice[i] = char
	}

	return string(byteSlice)
}
