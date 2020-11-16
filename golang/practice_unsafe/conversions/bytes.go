// Package conversions example string conversion
package conversions

import (
	"reflect"
	"unsafe"
)

// UnsafeBytesToStringByHeader bytes conversation to string.
//nolint:gosec
func UnsafeBytesToStringByHeader(b []byte) string {
	var s string

	bp := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sp := (*reflect.StringHeader)(unsafe.Pointer(&s))

	sp.Data = bp.Data
	sp.Len = bp.Len

	return s
}

// UnsafeBytesToString bytes conversation to string.
func UnsafeBytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) //nolint:gosec
}

// UnsafeStringToByteSliceByHeader string conversation to bytes.
//nolint:gosec
func UnsafeStringToByteSliceByHeader(s string) []byte {
	var b []byte

	bp := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sp := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bp.Data = sp.Data
	bp.Len = sp.Len
	bp.Cap = sp.Len

	return b
}

// UnsafeUint64sToBytes conversation slice uint64 to bytes.
// nolint:gosec
func UnsafeUint64sToBytes(a []uint64) []byte {
	var b []byte

	ah := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	bh.Data = ah.Data
	bh.Len = ah.Len * int(unsafe.Sizeof(a[0]))
	bh.Cap = bh.Len

	return b
}

// SliceUint64ToBytes conversation slice uint64 to bytes.
func SliceUint64ToBytes(sliceUint64 []uint64) []byte {
	const (
		x = 0xff
		l = 8
	)

	r := make([]byte, 0, l*len(sliceUint64))

	for _, item := range sliceUint64 {
		for i := uint64(0); i < l; i++ {
			r = append(r, byte((item>>(i*l))&x))
		}
	}

	return r
}
