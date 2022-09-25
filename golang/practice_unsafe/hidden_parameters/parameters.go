// Package parameters testing getting hidden attributes
package parameters

import (
	"bytes"
	"unsafe"

	"github.com/Nixolay/training/golang/practice_unsafe/hidden_parameters/hidden"
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

type bytesBuf struct {
	buf []byte
}

// UnsafeGetBuf gets buffer from bytes.Buffer.
//nolint:gosec
func UnsafeGetBuf(bb *bytes.Buffer) []byte {
	return (*bytesBuf)(unsafe.Pointer(bb)).buf
}
