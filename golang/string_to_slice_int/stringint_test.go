package stringint_test

import (
	"testing"

	. "github.com/Nixolay/training/golang/string_to_slice_int"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStringToSliceInt(t *testing.T) {
	Convey("Test converting string to slice integers", t, func() {
		Convey("Correct", func() {
			actual, err := StringToSliceInt("9119")
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []int{9, 1, 1, 9})
		})
		Convey("Incorrect", func() {
			_, err := StringToSliceInt("91a19")
			So(err, ShouldNotBeNil)
		})
	})
}
