package action

import (
	"github.com/restartfu/kot-solver/game"
	"time"
)

type WallDelayed struct{}

func (a WallDelayed) Intervals(_ game.Action) []time.Duration {
	return []time.Duration{
		time.Millisecond * 610,
		time.Millisecond * 100,
	}
}
