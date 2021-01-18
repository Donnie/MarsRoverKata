package rover

import "fmt"

// Rover properties
type Rover struct {
	LocX int8
	LocY int8
	Dir  string
}

// TouchDown on a Planet
func (r *Rover) TouchDown() string {
	return fmt.Sprintf(`(%d, %d) %s`, r.LocX, r.LocY, r.Dir)
}
