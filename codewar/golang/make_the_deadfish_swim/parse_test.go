package parsefish

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T)  {
	Convey("Codewar: make the deadfish swim",t, func() {
    Convey("just o", func() {
        So(Parse("ooo"),ShouldResemble,[]int{0,0,0})
    })
    Convey("o&i", func() {
        So(Parse("ioioio"),ShouldResemble,[]int{1,2,3})
    })
    Convey("o&i&d", func() {
        So(Parse("idoiido"),ShouldResemble,[]int{0,1})
    })
    Convey("o&i&d&s", func() {
        So(Parse("isoisoiso"),ShouldResemble,[]int{1,4,25})
    })
    Convey("codewars", func() {
        So(Parse("codewars"),ShouldResemble,[]int{0})
    })
})
}