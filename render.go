package main

import (
	"GoRoguelike/routines"
	cw "TCellConsoleWrapper/tcell_wrapper"
	"fmt"
	"time"
)

var (
	R_VIEWPORT_SIZE_X = 30
	R_VIEWPORT_SIZE_Y = 20
	R_VIEWPORT_CURR_X = 0
	R_VIEWPORT_CURR_Y = 0
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

//func r_areRealCoordsInViewport(x, y int) bool {
//	return x - R_VIEWPORT_CURR_X < R_VIEWPORT_SIZE_X && y - R_VIEWPORT_CURR_Y < R_VIEWPORT_SIZE_Y
//}

func r_CoordsToViewport(x, y int) (int, int) {
	vpx, vpy := x - R_VIEWPORT_CURR_X, y - R_VIEWPORT_CURR_Y
	if vpx > R_VIEWPORT_SIZE_X || vpy > R_VIEWPORT_SIZE_Y {
		return -1, -1
	}
	return vpx, vpy
}

func updateViewportCoords(p *p_pawn) {
	R_VIEWPORT_CURR_X = p.x - R_VIEWPORT_SIZE_X / 2
	R_VIEWPORT_CURR_Y = p.y - R_VIEWPORT_SIZE_Y / 2
}

func renderLevel(d *dungeon, flush bool) {
	cw.Clear_console()
	vismap := d.GetFieldOfVisionFrom(d.player.x, d.player.y)
	updateViewportCoords(d.player)
	// render level. vpx, vpy are viewport coords, whereas x, y are real coords.
	for x := R_VIEWPORT_CURR_X; x < R_VIEWPORT_CURR_X + R_VIEWPORT_SIZE_X; x++ {
		for y := 0; y < R_VIEWPORT_CURR_Y+ R_VIEWPORT_SIZE_Y; y++ {
			vpx, vpy := r_CoordsToViewport(x, y)
			if !areCoordinatesValid(x, y) {
				continue
			}
			cellRune := d.tiles[x][y].cCell.appearance
			cellColor := d.tiles[x][y].cCell.color
			if RENDER_DISABLE_LOS || vismap[x][y] {
				if !RENDER_DISABLE_LOS {
					d.tiles[x][y].wasSeenByPlayer = true
				}
				setFgColor(cellColor)
			} else {
				if d.tiles[x][y].wasSeenByPlayer {
					setFgColor(FogOfWarColor)
				}
			}
			if d.tiles[x][y].wasSeenByPlayer {
				cw.PutChar(cellRune, vpx, vpy)
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
			renderPawn(pawn, false)
		}
	}

	//render projectiles
	for _, proj := range d.projectiles {
		if areCoordinatesValid(proj.x, proj.y) && (RENDER_DISABLE_LOS || vismap[proj.x][proj.y]) {
			renderProjectile(proj)
		}
	}

	//render player
	renderPawn(d.player, false)

	renderPlayerStats(d)
	renderLog(false)

	if flush {
		cw.Flush_console()
	}
}

func renderProjectile(p *projectile) {
	setColor(cw.RED, cw.BLACK)
	x, y := r_CoordsToViewport(p.x, p.y)
	cw.PutChar('*', x, y)
}

func renderPawn(p *p_pawn, inverse bool) {
	app := p.appearance
	if inverse {
		setColor(cw.BLACK, cons_pawnColors[p.appearance])
	}	else {
		setFgColor(cons_pawnColors[p.appearance])
		if p.isPlayer() == false && p.aiData.state == AI_STAGGERED {
			setColor(cw.BLACK, cw.DARK_YELLOW)
		}
	}
	x, y := r_CoordsToViewport(p.x, p.y)
	cw.PutChar(app, x, y)
	setBgColor(cw.BLACK)
}

func renderItem(i *i_item) {
	setFgColor(i.ccell.color)
	x, y := r_CoordsToViewport(i.x, i.y)
	cw.PutChar(i.ccell.appearance, x, y)
}

func renderPlayerStats(d *dungeon) {
	player := d.player

	statsline := fmt.Sprintf("HP: (%d/%d) TIME: %d.%d", player.hp, player.maxhp,
		CURRENT_TURN/10, CURRENT_TURN%10)
	setFgColor(cw.DARK_RED)
	cw.PutString(statsline, 0, levelsizey)

	var weaponline string
	if player.weaponInHands != nil {
		weaponline = fmt.Sprintf("%s (%d/%d)", player.weaponInHands.name, player.weaponInHands.weaponData.ammo,
			player.weaponInHands.weaponData.maxammo)
	} else {
		weaponline = "fists"
	}
	cw.PutString(fmt.Sprintf("WEAP: %s", weaponline), len(statsline)+1, levelsizey)

	ammoLine := fmt.Sprintf("BULL:%d SHLL:%d RCKT:%d CELL:%d",
		player.inventory.ammo[AMMO_BULL], player.inventory.ammo[AMMO_SHEL], player.inventory.ammo[AMMO_RCKT], player.inventory.ammo[AMMO_CELL])
	setColor(cw.DARK_RED, cw.BLACK)
	cw.PutString(ammoLine, 0, levelsizey+1)
}

//func renderLine(char rune, fromx, fromy, tox, toy int, flush, exceptFirstAndLast bool) {
//	line := routines.GetLine(fromx, fromy, tox, toy)
//	setFgColor(cw.RED)
//	if exceptFirstAndLast {
//		for i := 1; i < len(line)-1; i++ {
//			cw.PutChar(char, line[i].X, line[i].Y)
//		}
//	} else {
//		for i := 0; i < len(line); i++ {
//			cw.PutChar(char, line[i].X, line[i].Y)
//		}
//	}
//	if flush {
//		cw.Flush_console()
//	}
//}

func renderTargetingLine(fromx, fromy, tox, toy int, flush bool, d *dungeon) {
	line := routines.GetLine(fromx, fromy, tox, toy)
	char := '?'
	if len(line) > 1  {
		char = getTargetingChar(line[1].X - line[0].X, line[1].Y - line[0].Y)
	}
	for i := 1; i < len(line); i++ {
		x, y := line[i].X, line[i].Y
		if d.isPawnPresent(x, y) {
			renderPawn(d.getPawnAt(x, y), true)
		} else {
			setFgColor(cw.YELLOW)
			if i == len(line)-1 {
				char = 'X'
			}
			viewx, viewy := r_CoordsToViewport(line[i].X, line[i].Y)
			cw.PutChar(char,viewx, viewy)
		}
	}
	if flush {
		cw.Flush_console()
	}
}

func renderBullet(currx, curry, tox, toy int, d *dungeon) {
	renderLevel(d, false)
	setFgColor(cw.YELLOW)
	bulletRune := '*'
	if !d.isPawnPresent(currx, curry) && !d.isTileOpaque(currx, curry) {
		bulletRune = getTargetingChar(tox-currx, toy-curry)
	}
	x, y := r_CoordsToViewport(currx, curry)
	cw.PutChar(bulletRune, x, y)
	cw.Flush_console()
	time.Sleep(25 * time.Millisecond)
}

func getTargetingChar(x, y int) rune{
	if abs(x) > 1 {
		x /= abs(x)
	}
	if abs(y) > 1 {
		y /= abs(y)
	}
	if x == 0 {
		return '|'
	}
	if y == 0 {
		return '-'
	}
	if x*y == 1 {
		return '\\'
	}
	if x*y == -1 {
		return '/'
	}
	return '?'
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func renderLog(flush bool) {
	setFgColor(cw.WHITE)
	for i := 0; i < LOG_HEIGHT; i++ {
		cw.PutString(log.last_msgs[i], 0, levelsizey+i+2)
	}
	if flush {
		cw.Flush_console()
	}
}
