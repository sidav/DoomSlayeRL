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
	newp := w.projectileExample
	newp.x = x
	newp.y = y
	newp.targetX = tx
	newp.targetY = ty
	return newp
}
