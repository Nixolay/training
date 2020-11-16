package parameters_test

import (
	"bytes"
	"testing"

	. "github.com/Nixolay/training/golang/practice_unsafe/hidden_parameters"
	"github.com/Nixolay/training/golang/practice_unsafe/hidden_parameters/hidden"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetHidden(t *testing.T) {
	Convey("Test get hidden parameters", t, func() {
		Convey("Func GetHidden", func() {
			data := hidden.CreateHidden()

			expected := GetHidden(data)
			So(expected, ShouldEqual, hidden.IsHidden)
		})

		Convey("Func UnsafeGetBuf", func() {
			data := "test data"
			bb := bytes.NewBufferString(data)

			buf := UnsafeGetBuf(bb)
			So(string(buf), ShouldEqual, data)
		})
	})
}
