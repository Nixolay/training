package singleton_test

import (
	"testing"

	"github.com/Nixolay/training/golang/patterns/creational/singleton"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleton(t *testing.T) {
	Convey("Test singleton", t, func() {
		single := singleton.GetInstance()
		a := 1 // I show that it is re-created for everyone

		Convey("Test instance", func() {
			So(single, ShouldNotBeNil)
			So(single.Increment(), ShouldEqual, 1)
			// I show that it is re-created for everyone
			So(a, ShouldEqual, 1)

			a++
		})
		Convey("Test increment", func() {
			So(single.Increment(), ShouldEqual, 2)
			// I show that it is re-created for everyone
			So(a, ShouldEqual, 1)

			a++
		})
		Convey("Test increment by new instance", func() {
			single := singleton.GetInstance()

			So(single.Increment(), ShouldEqual, 3)
			// I show that it is re-created for everyone
			So(a, ShouldEqual, 1)
		})
	})
}
