// https://www.codewars.com/kata/559a28007caad2ac4e000083/train/go
package rectangle_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/perimeter_of_squares_in_a_rectangle"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPerimeter(t *testing.T) {
	Convey("Perimeter of squares in a rectangle", t, func() {
		dotest(5, 80)
		dotest(7, 216)
		dotest(20, 114624)
		dotest(30, 14098308)
	})
}

func dotest(n, exp int) {
	ans := Perimeter(n)

	So(ans, ShouldEqual, exp)
}
