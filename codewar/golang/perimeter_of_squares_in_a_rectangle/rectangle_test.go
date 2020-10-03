//https://www.codewars.com/kata/559a28007caad2ac4e000083/train/go
package rectangle //nolint:testpackage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//nolint:gomnd
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
