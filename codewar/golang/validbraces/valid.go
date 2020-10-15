// Package validbraces ...
package validbraces

// ValidBraces valid braces.
func ValidBraces(str string) bool {
	m := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}

	arr := make([]string, 0)

	for _, r := range str {
		r := string(r)
		if len(arr) > 0 && m[arr[len(arr)-1]] == r {
			arr = arr[:len(arr)-1]
		} else {
			arr = append(arr, r)
		}
	}

	return len(arr) == 0
}
