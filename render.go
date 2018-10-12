package main

import (
	"DoomSlayeRL/routines"
	cw "TCellConsoleWrapper/tcell_wrapper"
	"fmt"
	"time"
)

var (
	R_VIEWPORT_WIDTH   = 40
	R_VIEWPORT_HEIGHT  = 20
	R_VIEWPORT_CURR_X  = 0
	R_VIEWPORT_CURR_Y  = 0
	RENDER_DISABLE_LOS bool
)

const (
	FogOfWarColor = cw.DARK_GRAY
)

//func r_areRealCoordsInViewport(x, y int) bool {
//	return x - R_VIEWPORT_CURR_X < R_VIEWPORT_WIDTH && y - R_VIEWPORT_CURR_Y < R_VIEWPORT_HEIGHT
//}

func r_CoordsToViewport(x, y int) (int, int) {
	vpx, vpy := x-R_VIEWPORT_CURR_X, y-R_VIEWPORT_CURR_Y
	if vpx >= R_VIEWPORT_WIDTH || vpy >= R_VIEWPORT_HEIGHT {
		return -1, -1
	}
	return vpx, vpy
}

func updateViewportCoords(p *p_pawn) {
	R_VIEWPORT_CURR_X = p.x - R_VIEWPORT_WIDTH/2
	R_VIEWPORT_CURR_Y = p.y - R_VIEWPORT_HEIGHT/2
}

func renderLevel(d *dungeon, flush bool) {
	cw.Clear_console()
	vismap := d.GetFieldOfVisionFor(d.player)
	updateViewportCoords(d.player)
	// render level. vpx, vpy are viewport coords, whereas x, y are real coords.
	for x := R_VIEWPORT_CURR_X; x < R_VIEWPORT_CURR_X+R_VIEWPORT_WIDTH; x++ {
		for y := 0; y < R_VIEWPORT_CURR_Y+R_VIEWPORT_HEIGHT; y++ {
			vpx, vpy := r_CoordsToViewport(x, y)
			if !areCoordinatesValid(x, y) {
				continue
			}
			cellRune := d.tiles[x][y].getAppearance().appearance
			cellColor := d.tiles[x][y].getAppearance().color
			if RENDER_DISABLE_LOS || vismap[x][y] {
				d.tiles[x][y].wasSeenByPlayer = true
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
	app := p.ccell.appearance
	clr := p.ccell.color
	if inverse {
		setColor(cw.BLACK, clr)
	} else {
		setFgColor(clr)
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

func renderBullets(currCoords []*routines.Vector, currDirs []*routines.Vector, d *dungeon) {
	renderLevel(d, false)

	for i:=0; i<len(currCoords); i++ {
		currx, curry := currCoords[i].GetRoundedCoords()
		tox, toy := currDirs[i].GetRoundedCoords()
		setFgColor(cw.YELLOW)
		bulletRune := '*'
		if !d.isPawnPresent(currx, curry) && !d.isTileOpaque(currx, curry) {
			bulletRune = getTargetingChar(tox, toy)
		}
		x, y := r_CoordsToViewport(currx, curry)
		cw.PutChar(bulletRune, x, y)
	}
	cw.Flush_console()
	time.Sleep(35 * time.Millisecond)
}

//
// UI-related stuff below
//

func renderPlayerStats(d *dungeon) {
	player := d.player
	pinv := player.inventory
	statusbarsWidth := 80 - R_VIEWPORT_WIDTH - 3

	hpPercent := player.getHpPercent()
	var hpColor int
	switch {
	case hpPercent < 33:
		hpColor = cw.RED
		break
	case hpPercent < 66:
		hpColor = cw.YELLOW
		break
	default:
		hpColor = cw.DARK_GREEN
		break
	}
	setFgColor(hpColor)

	renderStatusbar(fmt.Sprintf("HP: (%d/%d)", player.hp, player.maxhp), player.hp, player.maxhp,
		R_VIEWPORT_WIDTH+1, 0, statusbarsWidth, hpColor)

	if player.wearedArmor == nil {
		setFgColor(cw.BEIGE)
		cw.PutString("No armor", R_VIEWPORT_WIDTH+1, 1)
	} else {
		setFgColor(player.wearedArmor.ccell.color)
		renderStatusbar(fmt.Sprintf("ARMOR: (%d/%d)", player.wearedArmor.armorData.currArmor, player.wearedArmor.armorData.maxArmor),
			player.wearedArmor.armorData.currArmor, player.wearedArmor.armorData.maxArmor, R_VIEWPORT_WIDTH+1, 1, statusbarsWidth, player.wearedArmor.ccell.color)
	}

	setFgColor(cw.BEIGE)
	if player.weaponInHands != nil {
		renderStatusbar(fmt.Sprintf("%s (%d/%d)", player.weaponInHands.name, player.weaponInHands.weaponData.ammo,
			player.weaponInHands.weaponData.maxammo), player.weaponInHands.weaponData.ammo,
			player.weaponInHands.weaponData.maxammo, R_VIEWPORT_WIDTH+1, 2, statusbarsWidth, cw.DARK_YELLOW)
	} else {
		cw.PutString("Barehanded", R_VIEWPORT_WIDTH+1, 2)
	}

	setFgColor(cw.BEIGE)
	cw.PutString(fmt.Sprintf("INV: %d/%d", len(pinv.items), pinv.maxItems), R_VIEWPORT_WIDTH+1, 3)

	setColor(cw.BEIGE, cw.BLACK)
	ammoLine := fmt.Sprintf("BULL:%d/%d", pinv.ammo[AMMO_BULL], pinv.maxammo[AMMO_BULL])
	cw.PutString(ammoLine, R_VIEWPORT_WIDTH+1, 4)
	ammoLine = fmt.Sprintf("SHLL:%d/%d", pinv.ammo[AMMO_SHEL], pinv.maxammo[AMMO_SHEL])
	cw.PutString(ammoLine, R_VIEWPORT_WIDTH+1, 5)
	ammoLine = fmt.Sprintf("RCKT:%d/%d", pinv.ammo[AMMO_RCKT], pinv.maxammo[AMMO_RCKT])
	cw.PutString(ammoLine, R_VIEWPORT_WIDTH+1, 6)
	ammoLine = fmt.Sprintf("CELL:%d/%d", pinv.ammo[AMMO_CELL], pinv.maxammo[AMMO_CELL])
	cw.PutString(ammoLine, R_VIEWPORT_WIDTH+1, 7)

	timeline := fmt.Sprintf("TIME: %d.%d (%d.%d)", CURRENT_TURN/10, CURRENT_TURN%10,
		player.playerData.lastSpentTimeAmount/10, player.playerData.lastSpentTimeAmount%10)
	cw.PutString(timeline, R_VIEWPORT_WIDTH+1, 9)
}

func renderTargetingLine(fromx, fromy, tox, toy int, flush bool, d *dungeon) {
	renderLevel(d, false)
	line := routines.GetLine(fromx, fromy, tox, toy)
	char := '?'
	if len(line) > 1 {
		dirVector := routines.CreateVectorByStartAndEndInt(fromx, fromy, tox, toy)
		dirVector.TransformIntoUnitVector()
		dirx, diry := dirVector.GetRoundedCoords()
		char = getTargetingChar(dirx, diry)
	}
	if fromx == tox && fromy == toy {
		renderPawn(d.player, true)
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
			cw.PutChar(char, viewx, viewy)
		}
	}
	if flush {
		cw.Flush_console()
	}
}

func renderStatusbar(name string, curvalue, maxvalue, x, y, width, barColor int) {
	barTitle := name
	cw.PutString(barTitle, x, y)
	barWidth := width - len(name)
	filledCells := barWidth * curvalue / maxvalue
	barStartX := x + len(barTitle) + 1
	for i := 0; i < barWidth; i++ {
		if i < filledCells {
			setFgColor(barColor)
			cw.PutChar('=', i+barStartX, y)
		} else {
			setFgColor(cw.DARK_BLUE)
			cw.PutChar('-', i+barStartX, y)
		}
	}
}

func getTargetingChar(x, y int) rune {
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
		cw.PutString(log.last_msgs[i].getText(), 0, R_VIEWPORT_HEIGHT+i)
	}
	if flush {
		cw.Flush_console()
	}
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
