package sum_test

import (
	"testing"

	. "github.com/Nixolay/training/golang/practice_unsafe/sum"
	. "github.com/smartystreets/goconvey/convey"
)

const numberOfChecks = 100

func TestSum(t *testing.T) {
	Convey("Test sum int range", t, func() {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		actual := 45
		Convey("Test function sum", func() {
			expected := Sum(data)
			So(expected, ShouldEqual, actual)
		})
		Convey("Test function sum by range", func() {
			expected := ByRange(data)
			So(expected, ShouldEqual, actual)
		})
		Convey("Test function sum by unsafe", func() {
			expected := UnsafeSum(data)
			So(expected, ShouldEqual, actual)
		})
	})
}

func BenchmarkSum(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < numberOfChecks; i++ {
		Sum(data)
	}
}

func BenchmarkSumByRange(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < numberOfChecks; i++ {
		ByRange(data)
	}
}

func BenchmarkUnsafeSum(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < numberOfChecks; i++ {
		UnsafeSum(data)
	}
}
