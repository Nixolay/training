// Package stringint example converting rune to number
package stringint

import (
	"errors"
	"unicode"
)

// StringToSliceInt converted string to int slice.
//nolint:goerr113
func StringToSliceInt(str string) ([]int, error) {
	sliceInt := make([]int, len(str))

	for i, r := range str {
		if !unicode.IsDigit(r) {
			return nil, errors.New("contains a character other than a number")
		}

		sliceInt[i] = int(r - '0')
	}

	return sliceInt, nil
}
