package game

import (
	"github.com/restartfu/kot-solver/game/windows"
	"log"
	"time"
)

type Dungeon struct {
	actions []Action
}

func NewDungeon(actions ...Action) Dungeon {
	return Dungeon{actions: actions}
}

func (d Dungeon) Solve() {
	log.Printf("Begining to solve dungeon . . .")

	var lastJump = time.Now()

	time.Sleep(time.Second / 2)
	windows.Jump()

	for j, i := range d.actions {
		var previous Action
		if j > 0 {
			previous = d.actions[j-1]
		}
		for _, j := range i.Intervals(previous) {
			for {
				if time.Since(lastJump) >= j {
					lastJump = time.Now()
					windows.Jump()
					break
				}
				time.Sleep(time.Millisecond)
			}
		}
	}
}

/*
 ______________
|    **        |
|    **        |
|**      **  **|
|**      **  **|
|****  **xx\  \|
|****  **xx   \|
|****          |
|****  /\      |
 --------------
*/

/*func Pegasus() Dungeon {
	return Dungeon{actions: []Action{
		RegularJump(930),
		RegularJump(570),
		RegularJump(1000),
	}}
}*/
