package main

import (
	cw "GoConsoleWrapper/console_wrapper"
	"fmt"
)

var (
	cons_pawnColors = map[rune]int{
		'@': cw.GREEN,
		'z': cw.BEIGE,
		'i': cw.RED,
	}
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
	//render items
	for _, item := range d.items {
		renderItem(&item)
	}

	//render pawns
	for i := 0; i < len(d.pawns); i++ {
		renderPawn(&d.pawns[i])
	}

	//render player
	renderPawn(&d.player)

	renderPlayerStats(d)
	renderLog()

	cw.Flush_console()
}

func renderPawn(p *p_pawn) {
	app := p.appearance
	cw.Set_color(cons_pawnColors[p.appearance], nil)
	x := p.x
	y := p.y
	cw.Put_char(app, x, y)
}

func renderItem(i *i_item) {
	app := i.appearance
	cw.Set_color(cw.RED, nil)
	x := i.x
	y := i.y
	cw.Put_char(app, x, y)
}

func renderPlayerStats(d *dungeon) {
	player := &d.player
	cw.Set_color(cw.RED, nil)
	cw.Put_string(fmt.Sprintf("HP: (%d/%d) TIME: %d.%d", player.hp, player.maxhp, curr_time/10, curr_time%10), 0, levelsizey)
}

func renderLog() {
	cw.Set_color(cw.WHITE, nil)
	for i := 0; i < LOG_HEIGHT; i++ {
		cw.Put_string(log.last_msgs[i], 0, levelsizey+i+1)
	}
}
