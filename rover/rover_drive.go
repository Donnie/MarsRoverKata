package rover

// drive is the driving force
func (r *Rover) drive(com string) *Rover {
	r.drivers[r.Dir](r.steppers[com])
	return r
}

// goEast takes the rover east
func (r *Rover) goEast(step int) {
	r.LocY += step
}

// goWest takes the rover west
func (r *Rover) goWest(step int) {
	r.LocY -= step
}

// goNorth takes the rover north
func (r *Rover) goNorth(step int) {
	r.LocX += step
}

// goSouth takes the rover south
func (r *Rover) goSouth(step int) {
	r.LocX -= step
}
