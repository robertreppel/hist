package ship

// Register a ship so it can be tracked.
func Register(name string, location string) (*Ship, string, error) {
	if len(name) == 0 {
		return nil, "Ship name cannot be blank.", nil
	}
	if len(location) == 0 {
		return nil, "Location cannot be blank.", nil
	}
	var ship Ship
	ship.trackChange(Registered{ShipID: name, Location: location})
	return &ship, "Success", nil
}
