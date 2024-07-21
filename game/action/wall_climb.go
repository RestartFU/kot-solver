package action

import (
	"github.com/restartfu/kot-solver/game"
	"time"
)

const DefaultWallClimbDelay float64 = 400

type WallClimb struct {
	Steps            float64
	Jumps            int
	DelaysDifference []time.Duration
}

func (a WallClimb) Intervals(previous game.Action) []time.Duration {
	i := []time.Duration{
		stepsToMilliseconds(a.Steps) + time.Millisecond*509,
	}
	switch previous.(type) {
	case CloseCorner:
		i[0] -= time.Millisecond * 210
	}

	for j := 0; j < a.Jumps; j++ {
		if j < len(a.DelaysDifference) {
			if a.DelaysDifference[j] == 0 {
				i = append(i, time.Duration(DefaultWallClimbDelay)*time.Millisecond)
			} else {
				i = append(i, time.Duration(DefaultWallClimbDelay)+a.DelaysDifference[j])
			}
		} else {
			i = append(i, time.Duration(DefaultWallClimbDelay)*time.Millisecond)
		}
	}
	return i
}
