package main

import (
	cw "GoConsoleWrapper/console_wrapper"
	"GoRoguelike/routines"
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
	vismap := d.GetFieldOfVisionFrom(d.player.x, d.player.y)
	// render level
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			if vismap[x][y] {
				cw.Set_color(cw.BEIGE, nil)
				cw.Put_char(d.tiles[x][y].Appearance, x, y)
			} else {
				cw.Set_color(cw.BLUE, nil)
				cw.Put_char(d.tiles[x][y].Appearance, x, y)
			}
		}
	}
	//render items
	for _, item := range d.items {
		if vismap[item.x][item.y] {
			renderItem(item)
		}
	}

	//render pawns
	for _, pawn := range d.pawns {
		if vismap[pawn.x][pawn.y] {
			renderPawn(pawn)
		}
	}

	//render player
	renderPawn(d.player)

	renderPlayerStats(d)
	renderLog(false)

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
	player := d.player

	var weaponline string
	if player.weaponInHands != nil {
		weaponline = fmt.Sprintf("%s (%d/%d)", player.weaponInHands.name, player.weaponInHands.weaponData.ammo,
			player.weaponInHands.weaponData.maxammo)
	} else {
		weaponline = "fists"
	}
	ammoLine := fmt.Sprintf("BULL:%d SHLL:%d RCKT:%d CELL:%d",
		player.inventory.bullets, player.inventory.shells, player.inventory.rockets, player.inventory.cells)
	cw.Set_color(cw.RED, nil)
	cw.Put_string(fmt.Sprintf("HP: (%d/%d) TIME: %d.%d WEAP: %s", player.hp, player.maxhp, curr_time/10, curr_time%10, weaponline), 0, levelsizey)
	cw.Set_color(cw.RED, nil)
	cw.Put_string(ammoLine, 0, levelsizey+1)
}

func renderLine(char rune, fromx, fromy, tox, toy int) {
	line := routines.GetLine(fromx, fromy, tox, toy)
	cw.Set_color(cw.RED, nil)
	for i := 0; i < len(line); i++ {
		cw.Put_char(char, line[i].X, line[i].Y)
	}
}

func renderLog(flush bool) {
	cw.Set_color(cw.WHITE, nil)
	for i := 0; i < LOG_HEIGHT; i++ {
		cw.Put_string(log.last_msgs[i], 0, levelsizey+i+2)
	}
	if flush {
		cw.Flush_console()
	}
}
