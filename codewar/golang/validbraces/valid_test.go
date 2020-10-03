package validbraces //nolint:testpackage

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidBraces(t *testing.T) {
	Convey("Valid Braces", t, func() {
		singleTest("(){}[]", true)
		singleTest("([{}])", true)
		singleTest("(}", false)
		singleTest("[(])", false)
		singleTest("[({)](]", false)
	})
}

func singleTest(str string, res bool) {
	fmt.Printf("should return %v for \"%v\"\n", res, str)

	So(ValidBraces(str), ShouldEqual, res)
}
