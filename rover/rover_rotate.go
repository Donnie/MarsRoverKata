package rover

// rotate is the rotating force
func (r *Rover) rotate(com string) *Rover {
	r.rotators[r.Dir](com)
	return r
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
