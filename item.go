package main

type (
	i_weaponData struct {
		slot, dnum, dval, dmod, ammo, maxammo int
	}

	i_item struct {
		x, y       int
		appearance rune
		name       string
		weaponData *i_weaponData
	}
)

func (i *i_item) getType() string {
	if i.weaponData != nil {
		return "weapon"
	}
	return "item"
}
