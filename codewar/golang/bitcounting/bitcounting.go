package bitcounting

import (
	"math/bits"
)

func CountBits(n uint) int {
	return bits.OnesCount(n)
}
