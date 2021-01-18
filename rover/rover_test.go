package rover

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTouchDown(t *testing.T) {
	rover := Rover{
		LocX: 4,
		LocY: 2,
		Dir:  "EAST",
	}
	expected := "(4, 2) EAST"
	found := rover.TouchDown()
	require.Equal(t, expected, found)
}
