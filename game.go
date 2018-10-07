package main

import "GoRoguelike/routines"

const (
	levelsizex = 80
	levelsizey = 20
)

var (
	GAME_IS_RUNNING bool
	log             LOG
	CURRENT_TURN    int
)

type game struct {
	dung dungeon
}

func areCoordinatesValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < levelsizex && y < levelsizey
}

func areCoordinatesInRangeFrom(fx, fy, tx, ty, srange int) bool {
	return (tx-fx)*(tx-fx) + (ty-fy)*(ty-fy) < srange * srange 
}

func (g *game) runGame() {
	routines.Randomize()
	GAME_IS_RUNNING = true
	d := dungeon{}
	d.initialize_level()
	log = LOG{}

	for GAME_IS_RUNNING {
		m_moveProjectiles(&d)
		if d.player.isTimeToAct() {
			renderLevel(&d, true)
			plr_playerControl(&d)
		}
		checkDeadPawns(&d)
		for i := 0; i < len(d.pawns); i++ {
			if d.pawns[i].isTimeToAct() {
				ai_decideMove(d.pawns[i], &d)
			}
		}
		CURRENT_TURN++
	}
}
