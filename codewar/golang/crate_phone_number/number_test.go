package number_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/crate_phone_number"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNumber(t *testing.T) {
	Convey("Test Create Phone Number", t, func() {
		So(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), ShouldEqual, "(123) 456-7890")
	})
}
