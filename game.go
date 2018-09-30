package main

const (
	levelsizex = 10
	levelsizey = 10
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
	GAME_IS_RUNNING = true
	d := dungeon{}
	d.initialize_level()
	log = LOG{}

	for GAME_IS_RUNNING {
		renderLevel(&d)
		playerControl(&d)
		checkDeadPawns(&d)
		for i := 0; i < len(d.pawns); i++ {
			ai_decideMove(&d.pawns[i], &d)
		}
		curr_time++
	}
}
