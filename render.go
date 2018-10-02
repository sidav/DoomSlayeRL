package main

import (
	"GoRoguelike/routines"
	cw "TCellConsoleWrapper/tcell_wrapper"
	"fmt"
)

var (
	RENDER_DISABLE_LOS bool
	cons_pawnColors    = map[rune]int{
		'@': cw.GREEN,
		'z': cw.BEIGE,
		'i': cw.RED,
	}
)

const (
	FogOfWarColor = cw.DARK_GRAY
)

func renderLevel(d *dungeon, flush bool) {
	cw.Clear_console()
	vismap := d.GetFieldOfVisionFrom(d.player.x, d.player.y)
	// render level
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			cellRune := d.tiles[x][y].cCell.appearance
			cellColor := d.tiles[x][y].cCell.color
			if RENDER_DISABLE_LOS || vismap[x][y] {
				if !RENDER_DISABLE_LOS {
					d.tiles[x][y].wasSeenByPlayer = true
				}
				cw.SetFgColor(cellColor)
				cw.PutChar(cellRune, x, y)
			} else {
				if d.tiles[x][y].wasSeenByPlayer {
					cw.SetFgColor(FogOfWarColor)
					cw.PutChar(cellRune, x, y)
				}
			}
		}
	}
	//render items
	for _, item := range d.items {
		if RENDER_DISABLE_LOS || vismap[item.x][item.y] {
			renderItem(item)
		}
	}

	//render pawns
	for _, pawn := range d.pawns {
		if RENDER_DISABLE_LOS || vismap[pawn.x][pawn.y] {
			renderPawn(pawn)
		}
	}

	//render player
	renderPawn(d.player)

	renderPlayerStats(d)
	renderLog(false)

	if flush {
		cw.Flush_console()
	}
}

func renderPawn(p *p_pawn) {
	app := p.appearance
	cw.SetFgColor(cons_pawnColors[p.appearance])
	if p.isPlayer() == false && p.aiData.state == AI_STAGGERED {
		cw.SetColor(cw.BLACK, cw.DARK_YELLOW)
	}
	x := p.x
	y := p.y
	cw.PutChar(app, x, y)
	cw.SetBgColor(cw.BLACK)
}

func renderItem(i *i_item) {
	app := i.appearance
	cw.SetFgColor(cw.DARK_RED)
	x := i.x
	y := i.y
	cw.PutChar(app, x, y)
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
	cw.SetFgColor(cw.DARK_RED)
	cw.PutString(fmt.Sprintf("HP: (%d/%d) TIME: %d.%d WEAP: %s", player.hp, player.maxhp,
		CURRENT_TURN/10, CURRENT_TURN%10, weaponline), 0, levelsizey)
	cw.SetFgColor(cw.DARK_RED)
	cw.PutString(ammoLine, 0, levelsizey+1)
}

func renderLine(char rune, fromx, fromy, tox, toy int, flush, exceptFirstAndLast bool) {
	line := routines.GetLine(fromx, fromy, tox, toy)
	cw.SetFgColor(cw.RED)
	if exceptFirstAndLast {
		for i := 1; i < len(line)-1; i++ {
			cw.PutChar(char, line[i].X, line[i].Y)
		}
	} else {
		for i := 0; i < len(line); i++ {
			cw.PutChar(char, line[i].X, line[i].Y)
		}
	}
	if flush {
		cw.Flush_console()
	}
}

func renderLog(flush bool) {
	cw.SetFgColor(cw.WHITE)
	for i := 0; i < LOG_HEIGHT; i++ {
		cw.PutString(log.last_msgs[i], 0, levelsizey+i+2)
	}
	if flush {
		cw.Flush_console()
	}
}
