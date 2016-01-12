package ship

type Arrived struct {
	Port string
}

type Departed struct {
	Port   string
	ShipID string
}

type Registered struct {
	ShipID   string
	Location string
}
