package main

type (
	i_ammoData struct {
		bullets, shells, rockets, cells int
	}

	i_item struct {
		x, y       int
		ccell      *consoleCell
		name       string
		weaponData *i_weaponData
		ammoData   *i_ammoData
	}
)

func (i *i_item) getType() string {
	if i.weaponData != nil {
		return "weapon"
	}
	if i.ammoData != nil {
		return "ammo"
	}
	return "item"
}
