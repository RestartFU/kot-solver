package action

import (
	"github.com/restartfu/kot-solver/game"
	"time"
)

type DelayedRegularJump struct {
	Delay time.Duration
}

func (a DelayedRegularJump) Intervals(_ game.Action) []time.Duration {
	return []time.Duration{
		a.Delay,
	}
}
