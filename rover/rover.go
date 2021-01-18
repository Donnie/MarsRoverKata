package rover

import "fmt"

// Driver takes rover in a direction
type Driver func(step int8)

// Rotator takes rover in a direction
type Rotator func(com string)

// Rover properties
type Rover struct {
	LocX int8
	LocY int8
	Dir  string

	drivers  map[string]Driver
	rotators map[string]Rotator
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

	r.rotators = make(map[string]Rotator)
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
func (r *Rover) Rotate(com string) {
	r.rotators[r.Dir](com)
}

// FromEast rotates the rover east
func (r *Rover) FromEast(com string) {
	options := map[string]string{"L": "NORTH", "R": "SOUTH"}
	r.Dir = options[com]
}

// FromWest rotates the rover west
func (r *Rover) FromWest(com string) {
	options := map[string]string{"L": "SOUTH", "R": "NORTH"}
	r.Dir = options[com]
}

// FromNorth rotates the rover north
func (r *Rover) FromNorth(com string) {
	options := map[string]string{"L": "WEST", "R": "EAST"}
	r.Dir = options[com]
}

// FromSouth rotates the rover south
func (r *Rover) FromSouth(com string) {
	options := map[string]string{"L": "EAST", "R": "WEST"}
	r.Dir = options[com]
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
	r.Rotate("L")
	return r.Report()
}

// RotateRight rotates the rover 90° right
func (r *Rover) RotateRight() string {
	r.Rotate("R")
	return r.Report()
}

// RunMission runs multiple commands
func (r *Rover) RunMission(com string) (report string) {
	return
}
