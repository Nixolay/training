package inascorder_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/arethenumbersinorder"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInAscOrder(t *testing.T) {
	Convey("should test that the solution returns the correct value", t, func() {
		So(InAscOrder([]int{1, 2, 4, 7, 19}), ShouldEqual, true)
		So(InAscOrder([]int{1, 2, 3, 4, 5}), ShouldEqual, true)
		So(InAscOrder([]int{1, 6, 10, 18, 2, 4, 20}), ShouldEqual, false)
		So(InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}), ShouldEqual, false)
	})
}
