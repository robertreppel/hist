package ship

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadingShipFromEventHistory(t *testing.T) {
	Convey("Given a ship", t, func() {
		ship := Ship{}
		ship.transition(Registered{"FlyingDutchman", "At Sea"})
		Convey("when it's registered", func() {
			Convey("then it should be at sea.", func() {
				So(ship.location, ShouldEqual, "At Sea")
			})
		})
		Convey("when it arrives", func() {
			ship.transition(Arrived{Port: "Beaulieu", ShipID: "FlyingDutchman"})
			Convey("then it should be in port.", func() {
				So(ship.location, ShouldEqual, "Beaulieu")
			})
		})
		Convey("when it departs", func() {
			ship.transition(Departed{Port: "Beaulieu", ShipID: "FlyingDutchman"})
			Convey("then it should be at sea.", func() {
				So(ship.location, ShouldEqual, "At Sea")
			})
		})
	})
}
