// Package number returning string phone number.
package number

import "fmt"

// CreatePhoneNumber returning string phone number.
func CreatePhoneNumber(numbers [10]uint) string {
	args := make([]interface{}, 0, 10)

	for i := range numbers {
		args = append(args, numbers[i])
	}

	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", args...)
}
