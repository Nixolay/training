// Package rectangle perimeter calculate.
package rectangle

// Perimeter calculate.
func Perimeter(n int) int {
	a, b, sum := 1, 1, 2
	for n--; n > 0; n-- {
		a, b = a+b, a
		sum += a
	}

	four := 4

	return sum * four
}
