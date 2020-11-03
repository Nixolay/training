// Package conversion example string conversion
package conversion

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
