package action

import (
	"fmt"
	"github.com/restartfu/kot-solver/game"
	"time"
)

type CloseCorner struct {
	Steps float64
}

func (a CloseCorner) Intervals(previous game.Action) []time.Duration {

	i := []time.Duration{
		stepsToMilliseconds(a.Steps),
		time.Millisecond * 600,
	}

	if a.Steps == 10 {
		i[0] = 1930 * time.Millisecond
	}

	compensated := false

	switch previous.(type) {
	case CloseCorner:
		switch a.Steps {
		case 6:
			fmt.Println("6")
			i[0] -= time.Millisecond * 217
		case 10:
			i[0] -= time.Millisecond * 205
		//case 12:
		//i[0] -= time.Millisecond * 180
		default:
			i[0] -= time.Millisecond * 200
		}
		compensated = true
	case WallDelayed:
		i[0] -= time.Millisecond * 185
		compensated = true
	case GuardRightCornerDownwards:
		switch a.Steps {
		case 6:
			i[0] -= time.Millisecond * 260
		case 10:
		default:
			i[0] -= time.Millisecond * 255
		}
	case DragonClean:
		i[0] += time.Millisecond * 300
		compensated = true
	case CloseUpperTouchingCorner:
		i[0] -= time.Millisecond * 210
	}

	if a.Steps == 10 && !compensated {
		steps := i[0]
		delay := i[1]

		i[0] = time.Millisecond * 600
		i[1] = steps
		i = append(i, delay)
	} else if a.Steps == 10 && compensated {
		i[0] += time.Millisecond * 530
	}
	return i
}

func stepsToMilliseconds(steps float64) time.Duration {
	switch steps {
	case 0:
		return 0
	case 3:
		return 430 * time.Millisecond
	case 4:
		return 720 * time.Millisecond
	case 5:
		return 1010 * time.Millisecond
	case 6:
		return 1310 * time.Millisecond
	case 7:
		return 1590 * time.Millisecond
	case 8:
		return 1890 * time.Millisecond
	case 9:
		return 2170 * time.Millisecond
	case 10:
		return 2450 * time.Millisecond
	case 11:
		return 2760 * time.Millisecond
	case 12:
		return 3045 * time.Millisecond
	case 13:
		return 3345 * time.Millisecond
	case 14:
		return 3625 * time.Millisecond
	}
	panic("should never happen")
}
