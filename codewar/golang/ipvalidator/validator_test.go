package ipvalidator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIPValidator(t *testing.T) {
	Convey("should test correct", t, valid(IsValidIP))
	Convey("should test correct", t, valid(Is_valid_ip))
}

func valid(f func(string) bool) func() {
	return func() {
		So(f("12.255.56.1"), ShouldEqual, true)
		So(f(""), ShouldEqual, false)
		So(f("abc.def.ghi.jkl"), ShouldEqual, false)
		So(f("123.456.789.0"), ShouldEqual, false)
		So(f("12.34.56"), ShouldEqual, false)
		So(f("12.34.56 .1"), ShouldEqual, false)
		So(f("12.34.56.-1"), ShouldEqual, false)
		So(f("123.045.067.089"), ShouldEqual, false)
		So(f("127.1.1.0"), ShouldEqual, true)
		So(f("0.0.0.0"), ShouldEqual, true)
		So(f("0.34.82.53"), ShouldEqual, true)
		So(f("192.168.1.300"), ShouldEqual, false)
	}
}
