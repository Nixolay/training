package multiples_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/multiples_of_3_or_5"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMultiple3And5(t *testing.T) {
	Convey("Multiples of 3 and 5", t, func() {
		m, equal := 10, 23

		So(Multiple3And5(m), ShouldEqual, equal)
	})
}
