package camel //nolint:testpackage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToCamelCase(t *testing.T) {
	Convey("should handle basic cases", t, func() {
		dotest("", "")
		dotest("The_Stealth_Warrior", "TheStealthWarrior")
		dotest("the-Stealth-Warrior", "theStealthWarrior")
	})
}

func dotest(str, exp string) {
	println("input:", str)

	So(ToCamelCase(str), ShouldEqual, exp)
}
