package builder_test

import (
	"testing"

	"github.com/Nixolay/training/golang/patterns/creational/builder"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBuilder(t *testing.T) {
	Convey("Test builder", t, func() {
		var director builder.ManufacturingDirector
		Convey("Test creation car", func() {
			var car builder.CarBuilder
			vehicle := builder.VehicleProduct{Seats: 5, Wheels: 4, Structure: "Car"}

			director.SetBuilder(&car)
			director.Construct()

			So(car.GetVehicle(), ShouldResemble, vehicle)
		})
		Convey("Test creation bike", func() {
			var bike builder.BikeBuilder
			vehicle := builder.VehicleProduct{Seats: 2, Wheels: 2, Structure: "Motorbike"}

			director.SetBuilder(&bike)
			director.Construct()

			So(bike.GetVehicle(), ShouldResemble, vehicle)
		})
	})
}
