package rover

import "fmt"

// Driver takes rover in a direction
type Driver func()

// Rover properties
type Rover struct {
	LocX int8
	LocY int8
	Dir  string

	drivers map[string]Driver
}

// Drive is the driving force
func (r *Rover) Drive() {
	r.drivers = make(map[string]Driver)
	r.drivers["EAST"] = r.GoEast
	r.drivers["WEST"] = r.GoWest
	r.drivers["NORTH"] = r.GoNorth
	r.drivers["SOUTH"] = r.GoSouth
	r.drivers[r.Dir]()
}

// GoEast takes the rover east
func (r *Rover) GoEast() {
	r.LocY++
}

// GoWest takes the rover west
func (r *Rover) GoWest() {
	r.LocY--
}

// GoNorth takes the rover north
func (r *Rover) GoNorth() {
	r.LocX++
}

// GoSouth takes the rover south
func (r *Rover) GoSouth() {
	r.LocX--
}

// Report outputs the current status of rover
func (r *Rover) Report() string {
	return fmt.Sprintf(`(%d, %d) %s`, r.LocX, r.LocY, r.Dir)
}

// TouchDown on a Planet
func (r *Rover) TouchDown() string {
	return r.Report()
}

// Forward moves the rover one step in extant direction
func (r *Rover) Forward() string {
	r.Drive()
	return r.Report()
}
