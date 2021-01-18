package rover

import "fmt"

// Driver takes rover in a direction
type Driver func(step int8)

// Rover properties
type Rover struct {
	LocX int8
	LocY int8
	Dir  string

	drivers  map[string]Driver
	steppers map[string]int8
}

// NewRover creates a new rover
func NewRover(x, y int8, dir string) Rover {
	return Rover{
		LocX: x,
		LocY: y,
		Dir:  dir,
	}
}

// Drive is the driving force
func (r *Rover) Drive(com string) {
	r.steppers = make(map[string]int8)
	r.steppers["F"] = 1
	r.steppers["B"] = -1

	r.drivers = make(map[string]Driver)
	r.drivers["EAST"] = r.GoEast
	r.drivers["WEST"] = r.GoWest
	r.drivers["NORTH"] = r.GoNorth
	r.drivers["SOUTH"] = r.GoSouth
	r.drivers[r.Dir](r.steppers[com])
}

// GoEast takes the rover east
func (r *Rover) GoEast(step int8) {
	r.LocY += step
}

// GoWest takes the rover west
func (r *Rover) GoWest(step int8) {
	r.LocY -= step
}

// GoNorth takes the rover north
func (r *Rover) GoNorth(step int8) {
	r.LocX += step
}

// GoSouth takes the rover south
func (r *Rover) GoSouth(step int8) {
	r.LocX -= step
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
func (r *Rover) Forward() (out string) {
	r.Drive("F")
	return r.Report()
}

// Backward moves the rover one step against extant direction
func (r *Rover) Backward() (out string) {
	r.Drive("B")
	return r.Report()
}
