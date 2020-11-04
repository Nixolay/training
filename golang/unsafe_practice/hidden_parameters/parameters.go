// Package parameters testing getting hidden attributes
package parameters

import (
	"bytes"
	"unsafe"

	"github.com/Nixolay/training/golang/unsafe_practice/hidden_parameters/hidden"
)

// GetHidden allows you to get a hidden field, but the order of
// the attributes from the target structure must be respected.
//nolint:gosec
func GetHidden(h *hidden.Hidden) string {
	type hiddenStructure struct {
		DataNoHidden int
		h            string
	}

	hh := (*hiddenStructure)(unsafe.Pointer(h))

	return hh.h
}

// UnsafeGetBuf gets buffer from bytes.Buffer.
//nolint:gosec
func UnsafeGetBuf(bb *bytes.Buffer) []byte {
	type bytesBuf struct {
		buf []byte
	}

	b := (*bytesBuf)(unsafe.Pointer(bb))

	return b.buf
}
