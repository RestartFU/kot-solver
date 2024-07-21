package main

import (
	"github.com/restartfu/kot-solver/game"
	"github.com/restartfu/kot-solver/game/action"
	hook "github.com/robotn/gohook"
	"log"
	"time"
)

func Base3SawBhRg() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.GuardRightCornerDownwards{},
		action.CloseCorner{Steps: 6},
		action.WallClimb{Steps: 5, Jumps: 2, DelaysDifference: []time.Duration{0, time.Millisecond * 100}},

		action.DelayedRegularJump{Delay: time.Millisecond * 1200},
		action.DelayedRegularJump{Delay: time.Millisecond * 303},
		action.DelayedRegularJump{Delay: time.Millisecond*557 + (time.Millisecond * 80)},
		action.DelayedRegularJump{Delay: time.Millisecond * 607},
	}...)
}

func base130() game.Dungeon {
	return game.NewDungeon([]game.Action{
		//action.DelayedRegularJump{Delay: time.Millisecond * 21455},
		action.DelayedRegularJump{Delay: time.Millisecond * 1400},
		action.DelayedRegularJump{Delay: time.Millisecond * 945},
		action.DelayedRegularJump{Delay: time.Millisecond * 1200},
		action.DelayedRegularJump{Delay: time.Millisecond * 740},
		action.DelayedRegularJump{Delay: time.Millisecond * 495},
		action.DelayedRegularJump{Delay: time.Millisecond * 400},
		action.DelayedRegularJump{Delay: time.Millisecond * 750},
		action.DelayedRegularJump{Delay: time.Millisecond * 850},
	}...)
}

// Base131 : Saw Jump, Cannon, Spinner
func Base131() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.DelayedRegularJump{Delay: time.Millisecond * 5540},
		action.DelayedRegularJump{Delay: time.Millisecond * 530},
		action.DelayedRegularJump{Delay: time.Millisecond * 400},
		//action.DelayedRegularJump{Delay: time.Millisecond * 2000},
		//action.DelayedRegularJump{Delay: time.Millisecond * 600},
	}...)
}

func basex() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.DragonClean{Steps: 5},
		action.CloseCorner{Steps: 4},
		action.DragonClean{Steps: 4},
	}...)
}

func magnet() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.DelayedRegularJump{Delay: time.Millisecond * 1500},
		action.DelayedRegularJump{Delay: time.Millisecond * 600},
		action.DelayedRegularJump{Delay: time.Millisecond * 1500},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 610},
		action.DelayedRegularJump{Delay: time.Millisecond * 550},
		//action.DelayedRegularJump{Delay: time.Millisecond * 800},
	}...)
}

func basexdd() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.GuardRightCornerUpwards{},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 600},
		action.DelayedRegularJump{Delay: time.Millisecond * 800},
		action.DelayedRegularJump{Delay: time.Millisecond * 700},
		action.DelayedRegularJump{Delay: time.Millisecond * 1020},
	}...)
}

func BaseLoL() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.DelayedRegularJump{Delay: time.Millisecond * 550},
		action.DelayedRegularJump{Delay: time.Millisecond * 800},
		action.DelayedRegularJump{Delay: time.Millisecond * 555},
		action.DelayedRegularJump{Delay: time.Millisecond * 525},
		action.DelayedRegularJump{Delay: time.Millisecond * 705},
	}...)
}

func Base30() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.WallDelayed{},
		action.DelayedRegularJump{Delay: time.Millisecond * 1735},
		action.DelayedRegularJump{Delay: time.Millisecond * 630},
		action.DelayedRegularJump{Delay: time.Millisecond * 620},
		action.DelayedRegularJump{Delay: time.Millisecond * 400},
	}...)
}

func BaseTrampolineOverChest() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.DelayedRegularJump{Delay: time.Millisecond * 590},
		action.DelayedRegularJump{Delay: time.Millisecond * 670},
		action.DelayedRegularJump{Delay: time.Millisecond * 400},
		action.DelayedRegularJump{Delay: time.Millisecond * 650},
		action.DelayedRegularJump{Delay: time.Millisecond * 900},
		action.DelayedRegularJump{Delay: time.Millisecond * 300},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 750},
		action.DelayedRegularJump{Delay: time.Millisecond * 400},
		action.DelayedRegularJump{Delay: time.Millisecond * 1200},
		action.DelayedRegularJump{Delay: time.Millisecond * 580},
		action.DelayedRegularJump{Delay: time.Millisecond * 1065},
		action.DelayedRegularJump{Delay: time.Millisecond * 300},
		action.DelayedRegularJump{Delay: time.Millisecond * 300},
		action.DelayedRegularJump{Delay: time.Millisecond * 300},
	}...)
}

func BaseRgWarderCannon() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.DelayedRegularJump{Delay: time.Millisecond * 1300},
		action.DelayedRegularJump{Delay: time.Millisecond * 400},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 2000},
		action.DelayedRegularJump{Delay: time.Millisecond * 2350},
		action.DelayedRegularJump{Delay: time.Millisecond * 980},
		action.DelayedRegularJump{Delay: time.Millisecond * 860},
		action.DelayedRegularJump{Delay: time.Millisecond * 700},
		action.DelayedRegularJump{Delay: time.Millisecond * 3800},
		action.DelayedRegularJump{Delay: time.Millisecond * 480},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 500},
		action.DelayedRegularJump{Delay: time.Millisecond * 1700},
		action.DelayedRegularJump{Delay: time.Millisecond * 1600},
		action.DelayedRegularJump{Delay: time.Millisecond * 820},
		action.DelayedRegularJump{Delay: time.Millisecond * 1920},
	}...)
}

func Pegasus() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.DelayedRegularJump{Delay: time.Millisecond * 930},
		action.DelayedRegularJump{Delay: time.Millisecond * 570},
		action.DelayedRegularJump{Delay: time.Millisecond * 1000},
	}...)
}

func Base124() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.DelayedRegularJump{Delay: time.Millisecond * 7050},
		action.DelayedRegularJump{Delay: time.Millisecond * 300},
		action.DelayedRegularJump{Delay: time.Millisecond * 885},
		action.DelayedRegularJump{Delay: time.Millisecond * 350},
		action.DelayedRegularJump{Delay: time.Millisecond * 910},
		action.DelayedRegularJump{Delay: time.Millisecond * 850},
		action.DelayedRegularJump{Delay: time.Millisecond * 2000},
		action.DelayedRegularJump{Delay: time.Millisecond * 300},
	}...)
}

func Base1242() game.Dungeon {
	return game.NewDungeon([]game.Action{
		action.DelayedRegularJump{Delay: time.Millisecond * 13000},
		action.DelayedRegularJump{Delay: time.Millisecond * 300},
		action.DelayedRegularJump{Delay: time.Millisecond * 885},
		action.DelayedRegularJump{Delay: time.Millisecond * 320},
		action.DelayedRegularJump{Delay: time.Millisecond * 900},
		action.DelayedRegularJump{Delay: time.Millisecond * 760},
		action.DelayedRegularJump{Delay: time.Millisecond * 2000},
		action.DelayedRegularJump{Delay: time.Millisecond * 300},
	}...)
}

func main() {
	var lastJump time.Time

	hook.Register(hook.KeyDown, []string{"ctrl", "shift", "d"}, func(event hook.Event) {
		lastJump = time.Now()
		//for {

		BaseLoL().Solve()
		//time.Sleep(time.Second * 3)
		//}

	})

	hook.Register(hook.KeyDown, []string{"space"}, func(event hook.Event) {
		log.Printf("Time between last and recent jump: %0.3fs", time.Since(lastJump).Seconds())
		lastJump = time.Now()
	})

	s := hook.Start()
	<-hook.Process(s)
}
