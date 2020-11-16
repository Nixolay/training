// Package sum it is practice unsafe, array in a without bounds check.
package sum

import "unsafe"

// UnsafeSum Sum items array in a without bounds check.
// nolint:gosec
func UnsafeSum(arr []int) int {
	s := 0
	n := len(arr)
	ap := unsafe.Pointer(&arr[0])

	for i := 0; i < n; i++ {
		s += *(*int)(unsafe.Pointer(
			uintptr(ap) + uintptr(i)*unsafe.Sizeof(arr[0]),
		))
	}

	return s
}

// ByRange sum array items by range.
func ByRange(arr []int) int {
	s := 0

	for _, v := range arr {
		s += v
	}

	return s
}

// Sum array items.
func Sum(arr []int) int {
	s := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		s += arr[i]
	}

	return s
}
