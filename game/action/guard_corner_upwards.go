package action

import (
	"github.com/restartfu/kot-solver/game"
	"time"
)

type GuardRightCornerUpwards struct{}

func (a GuardRightCornerUpwards) Intervals(_ game.Action) []time.Duration {
	return []time.Duration{
		time.Millisecond * 900,
		time.Millisecond * 30,
	}
}
