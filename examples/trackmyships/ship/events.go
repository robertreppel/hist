package ship

// Arrived at port
type Arrived struct {
	Port   string
	ShipID string
}

// Departed from port
type Departed struct {
	Port   string
	ShipID string
}

// Registered ship in the ship's register
type Registered struct {
	ShipID   string
	Location string
}
