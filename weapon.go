package main

type i_weaponData struct {
	ammo, maxammo     int
	projectileExample *projectile
	hitscanData       *w_hitscan
}

func (w *i_weaponData) getType() string {
	if w.projectileExample != nil {
		return "projectile"
	}
	if w.hitscanData != nil {
		return "hitscan"
	}
	return "WEAPON_UNDEFINED"
}

func (w *i_weaponData) hasEnoughAmmoToShoot() bool {
	return w.ammo > 0 // TODO: variable ammo cost
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
