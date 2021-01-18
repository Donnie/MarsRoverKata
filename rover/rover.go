package rover

import "fmt"

// Driver takes rover in a direction
type Driver func(step int)

// Rotator takes rover in a direction
type Rotator func(com string)

// Rover properties
type Rover struct {
	Dir       string
	LocX      int
	LocY      int
	Obstacles [][]int

	drivers  map[string]Driver
	rotators map[string]Rotator
	steppers map[string]int
}

// NewRover creates a new rover
func NewRover(x, y int, dir string) (r *Rover) {
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

	r.steppers = make(map[string]int)
	r.steppers["F"] = 1
	r.steppers["B"] = -1

	return
}

// Report outputs the current status of rover
func (r *Rover) Report() string {
	return fmt.Sprintf(`(%d, %d) %s`, r.LocX, r.LocY, r.Dir)
}

// TouchDown on a Planet
func (r *Rover) TouchDown() *Rover {
	return r
}

// Forward moves the rover one step in extant direction
func (r *Rover) Forward() *Rover {
	r.drive("F")
	return r
}

// Backward moves the rover one step against extant direction
func (r *Rover) Backward() *Rover {
	r.drive("B")
	return r
}

// RotateLeft rotates the rover 90° left
func (r *Rover) RotateLeft() *Rover {
	r.rotate("L")
	return r
}

// RotateRight rotates the rover 90° right
func (r *Rover) RotateRight() *Rover {
	r.rotate("R")
	return r
}

// RunMission runs multiple commands and provides report
func (r *Rover) RunMission(com string) *Rover {
	var comms = map[rune]func() *Rover{
		'F': r.Forward,
		'B': r.Backward,
		'R': r.RotateRight,
		'L': r.RotateLeft,
	}

	for _, c := range com {
		if val, ok := comms[c]; ok {
			val()
		}
	}

	return r
}

// AddObstacles adds obstacles' maps
func (r *Rover) AddObstacles(obs [][]int) *Rover {
	r.Obstacles = obs
	return r
}
