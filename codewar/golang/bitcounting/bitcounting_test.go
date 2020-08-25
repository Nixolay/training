package bitcounting

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBitCounting(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	Convey("Bit Counting", t, func() {
		So(CountBits(uint(data[0])), ShouldEqual, data[0])
		So(CountBits(uint(data[4])), ShouldEqual, data[1])
		So(CountBits(uint(data[7])), ShouldEqual, data[3])
		So(CountBits(uint(data[9])), ShouldEqual, data[2])
		So(CountBits(uint(data[10])), ShouldEqual, data[2])
	})
}
