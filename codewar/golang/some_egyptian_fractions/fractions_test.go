package fractions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecompose(t *testing.T) {
	Convey("Tests Decompose", t, func() {
		So(Decompose("0"), ShouldResemble, []string{})
		So(Decompose("1"), ShouldResemble, []string{"1"})
		So(Decompose("12/4"), ShouldResemble, []string{"3"})
		So(Decompose("21/23"), ShouldResemble, []string{"1/2", "1/3", "1/13", "1/359", "1/644046"})
		So(Decompose("0.66"), ShouldResemble, []string{"1/2", "1/7", "1/59", "1/5163", "1/53307975"})
	})
}
