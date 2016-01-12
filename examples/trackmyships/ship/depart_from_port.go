package ship

func (ship *Ship) Depart() (string, error) {
	if ship.location == AtSea {
		return "Cannot depart from port: Ship is at sea.", nil
	}
	ship.trackChange(Departed{ShipID: ship.name, Port: ship.location})
	return "Success", nil
}
