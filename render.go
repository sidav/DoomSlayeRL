package main

import (
	cw "GoConsoleWrapper/console_wrapper"
)

func renderLevel(l *dungeon) {
	// render level
	for x:=0; x<levelsizex; x++ {
		for y:=0; y<levelsizey; y++{
			cw.Put_char(l.tiles[x][y].Appearance, x, y)
		}
	}
	//render player
	cw.Put_char(l.player.appearance, l.player.x, l.player.y)
	//render pawns
	for i:=0; i<len(l.pawns); i++ {
		app := l.pawns[i].appearance
		x := l.pawns[i].x
		y := l.pawns[i].y
		cw.Put_char(app, x, y)
	}
	cw.Flush_console()
}

