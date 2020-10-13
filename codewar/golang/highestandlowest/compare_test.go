package compare_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/highestandlowest"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHighAndLow(t *testing.T) {
	Convey("should test that the solution returns the correct value", t, func() {
		So(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"), ShouldEqual, "42 -9")
	})
}
