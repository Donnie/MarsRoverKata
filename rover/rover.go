package rover

// Rover properties
type Rover struct {
	LocX int8
	LocY int8
	Dir  string
}

// TouchDown on a Planet
func (r *Rover) TouchDown() string {
	return "(0, 0) NORTH"
}
