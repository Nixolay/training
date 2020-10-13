package stopgninnips_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/stopgninnips"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpinWords(t *testing.T) {
	Convey("should test that the solution returns the correct value for single word inputs", t, func() {
		So(SpinWords("Welcome"), ShouldEqual, "emocleW")
		So(SpinWords("to"), ShouldEqual, "to")
		So(SpinWords("CodeWars"), ShouldEqual, "sraWedoC")
	})
	Convey("should test that the solution returns the correct value for multiple word outputs", t, func() {
		So(SpinWords("Hey fellow warriors"), ShouldEqual, "Hey wollef sroirraw")
		So(SpinWords("Burgers are my favorite fruit"), ShouldEqual, "sregruB are my etirovaf tiurf")
		So(SpinWords("Pizza is the best vegetable"), ShouldEqual, "azziP is the best elbategev")
	})
}
