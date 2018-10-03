package main

import (
	"math"
)

type projectile struct {
	x, y, nextTurnToMove, turnsForOneTile int
	realx, realy, targX, targY            float32
	damageDice                            *dice
	cCell                                 *consoleCell
}

func (p *projectile) initTarget(tx, ty int) { // inits the target vector by absolute coordinates of the target.
	p.targX = float32(tx - p.x)
	p.targY = float32(ty - p.y)
	length := float32(math.Sqrt(float64(p.targX*p.targX + p.targY*p.targY)))
	p.targX /= length
	p.targY /= length
	p.realx, p.realy = float32(p.x), float32(p.y)
}

func (p *projectile) moveNextTile() {
	p.realx += p.targX
	p.realy += p.targY
	p.x = int(math.Round(float64(p.realx)))
	p.y = int(math.Round(float64(p.realy)))
}
