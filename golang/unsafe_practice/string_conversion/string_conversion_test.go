package conversion_test

import (
	"testing"

	. "github.com/Nixolay/training/golang/unsafe_practice/string_conversion"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	word = "test data is my big text"
	size = 1000
)

func TestStringConversion(t *testing.T) {
	Convey("Test string conversion", t, func() {
		Convey("Test bytes conversion to string", func() {
			data := []byte(word)
			expected := UnsafeBytesToString(data)
			So(expected, ShouldEqual, word)
		})

		Convey("Test bytes conversion to string by Header", func() {
			data := []byte(word)
			expected := UnsafeBytesToStringByHeader(data)
			So(expected, ShouldEqual, word)
		})

		Convey("Test string conversion to byte slice by Header", func() {
			actual := []byte(word)
			expected := UnsafeStringToByteSliceByHeader(word)
			So(expected, ShouldResemble, actual)
		})
	})
}

func BenchmarkUnsafeBytesToString(b *testing.B) {
	data := []byte(word)
	v := make(chan string, size)

	for i := 0; i < size; i++ {
		v <- UnsafeBytesToString(data)
	}

	close(v)
}

func BenchmarkUnsafeBytesToStringByHeader(b *testing.B) {
	data := []byte(word)
	v := make(chan string, size)

	for i := 0; i < size; i++ {
		v <- UnsafeBytesToStringByHeader(data)
	}

	close(v)
}

func BenchmarkBytesToString(b *testing.B) {
	data := []byte(word)
	v := make(chan string, size)

	for i := 0; i < size; i++ {
		v <- string(data)
	}

	close(v)
}

func BenchmarkUnsafeStringToByteSliceByHeader(b *testing.B) {
	v := make(chan []byte, size)

	for i := 0; i < size; i++ {
		v <- UnsafeStringToByteSliceByHeader(word)
	}

	close(v)
}

func BenchmarkStringToByteSlice(b *testing.B) {
	v := make(chan []byte, size)

	for i := 0; i < size; i++ {
		v <- []byte(word)
	}

	close(v)
}
