package main

const (
	levelsizex = 10
	levelsizey = 10
)

var (
	GAME_IS_RUNNING bool
)

type game struct {
	dung dungeon
}

func (g *game) runGame() {
	d := dungeon{}
	d.initialize_level()
	for GAME_IS_RUNNING {
		renderLevel(&d)
		playerControl(&d)
	}
}
