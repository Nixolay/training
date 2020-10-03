package splitstrings //nolint:testpackage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSolution(t *testing.T) {
	Convey("should test that the solution returns the correct value", t, func() {
		So(Solution("abc"), ShouldResemble, []string{"ab", "c_"})
		So(Solution("abcdef"), ShouldResemble, []string{"ab", "cd", "ef"})
	})
}
