package main

import (
	cw "GoConsoleWrapper/console_wrapper"
	"fmt"
)

func renderLevel(d *dungeon) {
	cw.Clear_console()
	// render level
	cw.Set_color(cw.BEIGE, nil)
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			cw.Put_char(d.tiles[x][y].Appearance, x, y)
		}
	}
	//render player
	cw.Set_color(cw.GREEN, nil)
	cw.Put_char(d.player.appearance, d.player.x, d.player.y)
	//render pawns
	for i := 0; i < len(d.pawns); i++ {
		app := d.pawns[i].appearance
		x := d.pawns[i].x
		y := d.pawns[i].y
		cw.Put_char(app, x, y)
	}

	renderPlayerStats(d)
	renderLog()

	cw.Flush_console()
}

func renderPlayerStats(d *dungeon) {
	player := &d.player
	cw.Set_color(cw.RED, nil)
	cw.Put_string(fmt.Sprintf("HP: (%d/%d)", player.hp, player.maxhp), 0, levelsizey)
}

func renderLog() {
	cw.Set_color(cw.WHITE, nil)
	for i := 0; i < LOG_HEIGHT; i++ {
		cw.Put_string(log.last_msgs[i], 0, levelsizey+i+1)
	}
}
