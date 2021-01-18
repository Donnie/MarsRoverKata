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
	r.drivers["EAST"] = r.goEast
	r.drivers["WEST"] = r.goWest
	r.drivers["NORTH"] = r.goNorth
	r.drivers["SOUTH"] = r.goSouth

	r.rotators = make(map[string]Rotator)
	r.rotators["EAST"] = r.fromEast
	r.rotators["WEST"] = r.fromWest
	r.rotators["NORTH"] = r.fromNorth
	r.rotators["SOUTH"] = r.fromSouth

	r.steppers = make(map[string]int8)
	r.steppers["F"] = 1
	r.steppers["B"] = -1

	return
}

// drive is the driving force
func (r *Rover) drive(com string) {
	r.drivers[r.Dir](r.steppers[com])
}

// goEast takes the rover east
func (r *Rover) goEast(step int8) {
	r.LocY += step
}

// goWest takes the rover west
func (r *Rover) goWest(step int8) {
	r.LocY -= step
}

// goNorth takes the rover north
func (r *Rover) goNorth(step int8) {
	r.LocX += step
}

// goSouth takes the rover south
func (r *Rover) goSouth(step int8) {
	r.LocX -= step
}

// rotate is the rotating force
func (r *Rover) rotate(com string) {
	r.rotators[r.Dir](com)
}

// fromEast rotates the rover from east
func (r *Rover) fromEast(com string) {
	options := map[string]string{"L": "NORTH", "R": "SOUTH"}
	r.Dir = options[com]
}

// fromWest rotates the rover from west
func (r *Rover) fromWest(com string) {
	options := map[string]string{"L": "SOUTH", "R": "NORTH"}
	r.Dir = options[com]
}

// fromNorth rotates the rover from north
func (r *Rover) fromNorth(com string) {
	options := map[string]string{"L": "WEST", "R": "EAST"}
	r.Dir = options[com]
}

// fromSouth rotates the rover from south
func (r *Rover) fromSouth(com string) {
	options := map[string]string{"L": "EAST", "R": "WEST"}
	r.Dir = options[com]
}

// report outputs the current status of rover
func (r *Rover) report() string {
	return fmt.Sprintf(`(%d, %d) %s`, r.LocX, r.LocY, r.Dir)
}

// TouchDown on a Planet
func (r *Rover) TouchDown() string {
	return r.report()
}

// Forward moves the rover one step in extant direction
func (r *Rover) Forward() string {
	r.drive("F")
	return r.report()
}

// Backward moves the rover one step against extant direction
func (r *Rover) Backward() string {
	r.drive("B")
	return r.report()
}

// RotateLeft rotates the rover 90° left
func (r *Rover) RotateLeft() string {
	r.rotate("L")
	return r.report()
}

// RotateRight rotates the rover 90° right
func (r *Rover) RotateRight() string {
	r.rotate("R")
	return r.report()
}

// RunMission runs multiple commands and provides report
func (r *Rover) RunMission(com string) (report string) {
	var comms = map[rune]func() string{
		'F': r.Forward,
		'B': r.Backward,
		'R': r.RotateRight,
		'L': r.RotateLeft,
	}

	for _, c := range com {
		if val, ok := comms[c]; ok {
			report = val()
		}
	}

	return
}
