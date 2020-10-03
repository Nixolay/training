package alphanumeric //nolint:testpackage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAlphanumeric(t *testing.T) {
	Convey("Should return the correct values for the sample test cases!", t, func() {
		So(alphanumeric(".*?"), ShouldEqual, false)
		So(alphanumeric("a"), ShouldEqual, true)
		So(alphanumeric("Mazinkaiser"), ShouldEqual, true)
		So(alphanumeric("hello world_"), ShouldEqual, false)
		So(alphanumeric("PassW0rd"), ShouldEqual, true)
		So(alphanumeric("     "), ShouldEqual, false)
		So(alphanumeric(""), ShouldEqual, false)
		So(alphanumeric("\n\t\n"), ShouldEqual, false)
		So(alphanumeric("ciao\n$$_"), ShouldEqual, false)
		So(alphanumeric("__ * __"), ShouldEqual, false)
		So(alphanumeric("&)))((("), ShouldEqual, false)
		So(alphanumeric("43534h56jmT3k"), ShouldEqual, true)
	})
}
