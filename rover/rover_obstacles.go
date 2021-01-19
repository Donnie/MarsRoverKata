package rover

func (r *Rover) detectObstacles(com rune) (out bool) {
	status := "STOPPED"
	v := NewRover(r.LocX, r.LocY, r.Dir) // we create a virtual rover to simulate
	comms := map[rune]func() *Rover{
		'F': v.Forward,
		'B': v.Backward,
		'R': v.RotateRight,
		'L': v.RotateLeft,
	}
	comms[com]()

	for _, obs := range r.Obstacles {
		if obs[0] == v.LocX && obs[1] == v.LocY {
			r.Status = &status
			return true
		}
	}
	return
}
