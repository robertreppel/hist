package ship

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDepartureFromPort(t *testing.T) {
	Convey("Given a ship is at sea", t, func() {
		ship, _, _ := Register("FlyingDutchman", AtSea)
		noOfEventsBeforeTryingToDepart := len(ship.Changes)
		Convey("when trying to record a departure from port", func() {
			result, err := ship.Depart()
			if err != nil {
				panic(err)
			}
			Convey("then departure cannot happen because the ship is at sea.", func() {
				noOfEventsAfterTryingToDepart := len(ship.Changes)
				So(noOfEventsAfterTryingToDepart, ShouldEqual, noOfEventsBeforeTryingToDepart)
				So(result, ShouldEqual, "Cannot depart from port: Ship is at sea.")
			})
		})
	})

	Convey("Given a ship is in port", t, func() {
		ship, _, _ := Register("SeaEagle", "Vancouver")
		noOfEventsBeforeTryingToDepart := len(ship.Changes)
		Convey("when trying to record a departure from port", func() {
			result, err := ship.Depart()
			if err != nil {
				panic(err)
			}
			Convey("then a departure is recorded.", func() {
				noOfEventsAfterTryingToDepart := len(ship.Changes)
				So(noOfEventsAfterTryingToDepart, ShouldEqual, noOfEventsBeforeTryingToDepart+1)
				secondEvent := ship.Changes[1]
				So(reflect.TypeOf(secondEvent) == reflect.TypeOf(Departed{}), ShouldBeTrue)
				So(result, ShouldEqual, "Success")
			})
			Convey("and the ship is at sea.", func() {
				So(ship.location, ShouldEqual, AtSea)
			})
		})
	})
}
