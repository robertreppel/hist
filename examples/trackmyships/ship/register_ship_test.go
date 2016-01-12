package ship

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewShip(t *testing.T) {
	Convey("When I try to register a ship without a name", t, func() {
		Convey("I should not be able to do it", func() {
			ship, result, err := Register("", AtSea)
			if err != nil {
				panic(err)
			}
			So(result, ShouldEqual, "Ship name cannot be blank.")
			So(ship, ShouldBeNil)
		})
	})

	Convey("When I try to register a ship without a location (either at port or at sea)", t, func() {
		Convey("I should not be able to do it", func() {
			ship, result, err := Register("Polarstern", "")
			if err != nil {
				panic(err)
			}
			So(result, ShouldEqual, "Location cannot be blank.")
			So(ship, ShouldBeNil)
		})
	})

	Convey("When I register a ship", t, func() {
		Convey("then a ship registration event occurs.", func() {
			ship, result, err := Register("Polarstern", AtSea)
			if err != nil {
				panic(err)
			}
			noOfEvents := len(ship.Changes)
			So(result, ShouldEqual, "Success")
			So(noOfEvents, ShouldEqual, 1)
			isShipRegisteredEvent := reflect.TypeOf(ship.Changes[0]) != reflect.TypeOf(Arrived{})
			So(isShipRegisteredEvent, ShouldBeTrue)
		})
	})

}
