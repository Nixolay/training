package factory_test

import (
	"testing"

	"github.com/Nixolay/training/golang/patterns/creational/factory"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreatePaymentMethod(t *testing.T) {
	Convey("", t, func() {
		Convey("Cash", func() {
			payment, err := factory.GetPaymentMethod(factory.Cash)
			So(err, ShouldBeNil)

			msg := payment.Pay(10.30)
			So(msg, ShouldContainSubstring, "paid using cash")
		})

		Convey("DebirCard", func() {
			payment, err := factory.GetPaymentMethod(factory.DebitCard)
			So(err, ShouldBeNil)

			msg := payment.Pay(22.30)
			So(msg, ShouldContainSubstring, "paid using debit card")
		})

		Convey("Wrong", func() {
			_, err := factory.GetPaymentMethod(0)
			So(err, ShouldNotBeNil)
		})
	})
}
