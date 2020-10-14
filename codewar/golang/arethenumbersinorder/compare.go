// Package inascorder checking sorted numbers.
package inascorder

import "sort"

// InAscOrder checking sorted numbers.
func InAscOrder(numbers []int) bool {
	return sort.IntsAreSorted(numbers)
}
