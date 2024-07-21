package action

import (
	"github.com/restartfu/kot-solver/game"
	"time"
)

type DragonClean struct {
	Steps float64
}

func (a DragonClean) Intervals(previous game.Action) []time.Duration {
	i := []time.Duration{
		stepsToMilliseconds(a.Steps) + time.Millisecond*500,
	}
	switch previous.(type) {
	case CloseCorner:
		i[0] -= time.Millisecond * 210
	}
	return i
}
