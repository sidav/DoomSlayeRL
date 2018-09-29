package main

const (
	levelsizex = 80
	levelsizey = 25
)

type game struct {
	dung dungeon
}

func (g *game) runGame(){
	d := dungeon{}
	d.initialize_level()
	for {
		renderLevel(&d)
		playerControl(&d.player)
	}
}
