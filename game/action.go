package game

import (
	"time"
)

type Action interface {
	Intervals(previous Action) []time.Duration
}

/*func RegularJump(milliseconds float64) Action {
	return Action{[]time.Duration{
		time.Duration(milliseconds) * time.Millisecond,
	}}
}

const DefaultWallClimbDelay float64 = 400

func WallClimb(jumps int, delay float64, individualDelays []float64) Action {
	intervals := []time.Duration{
		time.Duration(delay) * time.Millisecond,
	}
	for i := 0; i < jumps; i++ {
		if i < len(individualDelays) {
			if individualDelays[i] == 0 {
				intervals = append(intervals, time.Duration(DefaultWallClimbDelay)*time.Millisecond)
			} else {
				intervals = append(intervals, time.Duration(individualDelays[i])*time.Millisecond)
			}
		} else {
			intervals = append(intervals, time.Duration(DefaultWallClimbDelay)*time.Millisecond)
		}
	}
	return Action{intervals}
}
*/
