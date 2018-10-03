package main

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
	newp := &projectile{x: x, y: y, turnsForOneTile: w.projectileExample.turnsForOneTile, nextTurnToMove: CURRENT_TURN,
		damageDice: w.projectileExample.damageDice}
	newp.initTarget(tx, ty)
	// calculate current x, y
	newp.moveNextTile()
	return newp
}
