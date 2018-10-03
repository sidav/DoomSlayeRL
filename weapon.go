package main

import "GoRoguelike/routines"

type i_weaponData struct {
	slot, ammo, maxammo int
	damageDice          *dice
	projectileExample   *projectile
}

func (w *i_weaponData) getType() string {
	if w.projectileExample != nil {
		return "projectile"
	}
	return "hitscan"
}

// Pointers may fuck the thing up. Checks needed
func (w *i_weaponData) createProjectile(x, y, tx, ty int) *projectile {
	newp := &projectile{targetX: tx, targetY: ty, turnsForOneTile: w.projectileExample.turnsForOneTile, nextTurnToMove: CURRENT_TURN,
		damageDice: &dice{3, 3, 3}} // TODO: remove this temp values
	// calculate current x, y
	line := routines.GetLine(x, y, tx, ty)
	newp.x, newp.y = line[1].X, line[1].Y
	return newp
}
