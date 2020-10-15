// Package bitcounting the number of one bits ("population count").
package bitcounting

import (
	"math/bits"
)

// CountBits returns the number of one bits ("population count") in n.
func CountBits(n uint) int {
	// returning bits.OnesCount
	return bits.OnesCount(n)
}
