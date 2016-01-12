package ship

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArrivingAtNoLocation(t *testing.T) {
	Convey("Given a ship", t, func() {
		ship, _, _ := Register("FlyingDutchman", "UlanBataar")
		noOfChangesAfterRegistration := len(ship.Changes)
		Convey("when trying to record an arrival without a location", func() {
			result, err := ship.Arrive("")
			if err != nil {
				panic(err)
			}
			noOfChangesAfterTryingToRecordArrival := len(ship.Changes)
			Convey("then nothing is recorded because a ship cannot arrive nowhere.", nil)
			So(noOfChangesAfterRegistration, ShouldEqual, noOfChangesAfterTryingToRecordArrival)
			So(result, ShouldEqual, "Arrival port cannot be blank.")
		})
	})
}

func TestArrivingAtAPort(t *testing.T) {
	Convey("Given a ship that isn't at sea", t, func() {
		ship, _, _ := Register("FlyingDutchman", "Weymouth")
		Convey("when trying to record an arrival at a new port", func() {
			result, err := ship.Arrive("Rotterdam")
			if err != nil {
				panic(err)
			}
			Convey("then no arrival is recorded because the ship has not departed to sea from its current port.", func() {
				So(len(ship.Changes), ShouldEqual, 1)
				So(result, ShouldEqual, "Cannot arrive at a port: Ship is not at sea.")
			})
		})
	})
	Convey("Given a ship that is at a port", t, func() {
		ship, _, _ := Register("FlyingDutchman", "Portsmouth")
		Convey("when trying to record an arrival at the same port", func() {
			result, err := ship.Arrive("Portsmouth")
			if err != nil {
				panic(err)
			}
			Convey("then no arrival is recorded because the ship is already there.", func() {
				So(len(ship.Changes), ShouldEqual, 1)
				So(result, ShouldEqual, "Cannot record arrival: Ship is already at this port.")
			})
		})
	})

	Convey("Given a ship that is at sea", t, func() {
		ship, _, _ := Register("FlyingDutchman", AtSea)
		Convey("when recording arrival at a port", func() {
			result, err := ship.Arrive("Hull")
			if err != nil {
				panic(err)
			}
			Convey("then a 'Arrived' event occurs", func() {
				arrived := ship.Changes[1].(Arrived)
				So(arrived, ShouldNotBeNil)
				Convey("and 'Success' is returned.", func() {
					So(result, ShouldEqual, "Success")
				})
			})
		})
	})
}
