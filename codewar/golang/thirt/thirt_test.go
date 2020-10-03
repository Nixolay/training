package thirt //nolint:testpackage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//nolint:gomnd
func TestThirt(t *testing.T) {
	Convey("should handle basic cases", t, func() {
		testequal(1234567, 87)
		testequal(8529, 79)
		testequal(85299258, 31)
		testequal(5634, 57)
	})
}

func testequal(n int, exp int) {
	ans := Thirt(n)

	So(ans, ShouldEqual, exp)
}
