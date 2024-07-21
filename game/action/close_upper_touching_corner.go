package action

import (
	"github.com/restartfu/kot-solver/game"
	"time"
)

type CloseUpperTouchingCorner struct {
	Steps float64
}

func (a CloseUpperTouchingCorner) Intervals(previous game.Action) []time.Duration {
	i := []time.Duration{
		stepsToMilliseconds(a.Steps) + time.Millisecond*30,
		time.Millisecond * 600,
	}

	switch previous.(type) {
	case CloseCorner:
		i[0] -= time.Millisecond * 210
	}

	return i
}
