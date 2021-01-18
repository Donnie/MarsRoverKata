package rover

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTouchDown(t *testing.T) {
	rover := NewRover(4, 2, "EAST")
	expected := "(4, 2) EAST"
	found := rover.TouchDown()
	require.Equal(t, expected, found)
}

func TestMoveForward(t *testing.T) {
	rover := NewRover(4, 2, "EAST")
	expected := "(4, 3) EAST"
	found := rover.Forward()
	require.Equal(t, expected, found)
}

func TestMoveBackward(t *testing.T) {
	rover := NewRover(4, 2, "EAST")
	expected := "(4, 1) EAST"
	found := rover.Backward()
	require.Equal(t, expected, found)
}

func TestRotateLeft(t *testing.T) {
	rover := NewRover(4, 2, "EAST")
	expected := "(4, 2) NORTH"
	found := rover.RotateLeft()
	require.Equal(t, expected, found)
}

func TestRotateRight(t *testing.T) {
	rover := NewRover(4, 2, "EAST")
	expected := "(4, 2) SOUTH"
	found := rover.RotateRight()
	require.Equal(t, expected, found)
}
