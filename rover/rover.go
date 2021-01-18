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
	rotators map[string]func()
	steppers map[string]int8
}

// NewRover creates a new rover
func NewRover(x, y int8, dir string) (r *Rover) {
	r = &Rover{
		LocX: x,
		LocY: y,
		Dir:  dir,
	}

	r.drivers = make(map[string]Driver)
	r.drivers["EAST"] = r.GoEast
	r.drivers["WEST"] = r.GoWest
	r.drivers["NORTH"] = r.GoNorth
	r.drivers["SOUTH"] = r.GoSouth

	r.rotators = make(map[string]func())
	r.rotators["EAST"] = r.FromEast
	r.rotators["WEST"] = r.FromWest
	r.rotators["NORTH"] = r.FromNorth
	r.rotators["SOUTH"] = r.FromSouth

	r.steppers = make(map[string]int8)
	r.steppers["F"] = 1
	r.steppers["B"] = -1

	return
}

// Drive is the driving force
func (r *Rover) Drive(com string) {
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

// Rotate is the rotating force
func (r *Rover) Rotate() {
	r.rotators[r.Dir]()
}

// FromEast rotates the rover east
func (r *Rover) FromEast() {
	r.Dir = "NORTH"
}

// FromWest rotates the rover west
func (r *Rover) FromWest() {
	r.Dir = "SOUTH"
}

// FromNorth rotates the rover north
func (r *Rover) FromNorth() {
	r.Dir = "WEST"
}

// FromSouth rotates the rover south
func (r *Rover) FromSouth() {
	r.Dir = "EAST"
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
	r.Drive("F")
	return r.Report()
}

// Backward moves the rover one step against extant direction
func (r *Rover) Backward() string {
	r.Drive("B")
	return r.Report()
}

// RotateLeft rotates the rover 90° left
func (r *Rover) RotateLeft() string {
	r.Rotate()
	return r.Report()
}

// RotateRight rotates the rover 90° right
func (r *Rover) RotateRight() string {
	r.Rotate()
	return r.Report()
}
