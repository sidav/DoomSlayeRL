package main

type (
	i_item struct {
		x, y                int
		ccell               *consoleCell
		name                string
		weaponData          *i_weaponData
		ammoData            *i_ammoData
		instantlyPickupable bool
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
