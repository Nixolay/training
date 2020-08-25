package decimal_to_factorial

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func dotest1(nb int, exp string) {
	var ans = Dec2FactString(nb)
	So(ans, ShouldEqual, exp)
}
func dotest2(str string, exp int) {
	var ans = FactString2Dec(str)
	So(ans, ShouldEqual, exp)
}

func TestDescribe(t *testing.T) {
	Convey("Decimal to Factorial and Back", t, func() {

		Convey("should handle basic cases Dec2FactString", func() {
			dotest1(36288000, "A0000000000")
			dotest1(2982, "4041000")
		})

		Convey("should handle basic cases FactString2Dec", func() {
			dotest2("341010", 463)
			dotest2("4042100", 2990)

		})

	})
}
