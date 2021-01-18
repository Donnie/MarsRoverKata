package rover

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var rover = Rover{
	LocX: 4,
	LocY: 2,
	Dir:  "EAST",
}

func TestTouchDown(t *testing.T) {
	expected := "(4, 2) EAST"
	found := rover.TouchDown()
	require.Equal(t, expected, found)
}

func TestMoveForward(t *testing.T) {
	expected := "(4, 3) EAST"
	found := rover.Forward()
	require.Equal(t, expected, found)
}
