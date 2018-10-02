package main

import "GoRoguelike/routines"

const (
	levelsizex = 80
	levelsizey = 20
)

var (
	GAME_IS_RUNNING bool
	log             LOG
	curr_time       int
)

type game struct {
	dung dungeon
}

func (g *game) runGame() {
	routines.Randomize()
	GAME_IS_RUNNING = true
	d := dungeon{}
	d.initialize_level()
	log = LOG{}

	for GAME_IS_RUNNING {
		renderLevel(&d, true)
		plr_playerControl(&d)
		checkDeadPawns(&d)
		for i := 0; i < len(d.pawns); i++ {
			ai_decideMove(d.pawns[i], &d)
		}
		curr_time++
	}
}
