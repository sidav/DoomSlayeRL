package main

import "GoRoguelike/routines"

type projectile struct {
	x, y, targetX, targetY, nextTurnToMove, turnsForOneTile int
	damageDice                                              *dice
	cCell                                                   *consoleCell
}

func (p *projectile) moveNextTile() {
	line := routines.GetLineOver(p.x, p.y, p.targetX, p.targetY)
	p.x, p.y = line[1].X, line[1].Y
	p.targetX, p.targetY = line[len(line)-1].GetCoords()
}
