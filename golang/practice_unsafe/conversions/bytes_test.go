package conversions_test

import (
	"testing"

	. "github.com/Nixolay/training/golang/practice_unsafe/conversions"
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

		Convey("Test slice uint64 conversion to byte slice by Header", func() {
			data := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
			actual := SliceUint64ToBytes(data)

			expected := UnsafeUint64sToBytes(data)
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

func BenchmarkUnsafeUint64sToBytes(b *testing.B) {
	data := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	UnsafeUint64sToBytes(data)
}

func BenchmarkSliceUint64ToBytes(b *testing.B) {
	data := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	UnsafeUint64sToBytes(data)
}
