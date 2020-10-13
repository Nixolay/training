package alphanumeric_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/not_very_secure"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAlphanumeric(t *testing.T) {
	Convey("Should return the correct values for the sample test cases!", t, func() {
		So(Alphanumeric(".*?"), ShouldEqual, false)
		So(Alphanumeric("a"), ShouldEqual, true)
		So(Alphanumeric("Mazinkaiser"), ShouldEqual, true)
		So(Alphanumeric("hello world_"), ShouldEqual, false)
		So(Alphanumeric("PassW0rd"), ShouldEqual, true)
		So(Alphanumeric("     "), ShouldEqual, false)
		So(Alphanumeric(""), ShouldEqual, false)
		So(Alphanumeric("\n\t\n"), ShouldEqual, false)
		So(Alphanumeric("ciao\n$$_"), ShouldEqual, false)
		So(Alphanumeric("__ * __"), ShouldEqual, false)
		So(Alphanumeric("&)))((("), ShouldEqual, false)
		So(Alphanumeric("43534h56jmT3k"), ShouldEqual, true)
	})
}
