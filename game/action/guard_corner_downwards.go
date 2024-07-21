package action

import (
	"github.com/restartfu/kot-solver/game"
	"time"
)

type GuardRightCornerDownwards struct{}

func (a GuardRightCornerDownwards) Intervals(_ game.Action) []time.Duration {
	return []time.Duration{
		time.Millisecond * 790,
		time.Millisecond * 250,
	}
}
