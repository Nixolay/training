package aretheythesame

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func dotest(array1 []int, array2 []int, exp bool) {
	ans := Comp(array1, array2)

	So(ans, ShouldEqual, exp)
}

func TestComp(t *testing.T) {
	Convey(`Are they the "same"?`, t, func() {
		var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
		var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
		dotest(a1, a2, true)
		a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
		a2 = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
		dotest(a1, a2, false)
		a1 = nil
		a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
		dotest(a1, a2, false)
		a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
		a2 = []int{121, 14641, 20736, 36100, 25921, 361, 20736, 361}
		dotest(a1, a2, false)
	})
}
