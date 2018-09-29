package main

const (
	levelsizex = 10
	levelsizey = 10
)

var (
	GAME_IS_RUNNING bool
	log             LOG
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
	}
}
