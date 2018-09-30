package main

type inventory struct {
	items                           []*i_item
	maxItems                        int
	bullets, shells, rockets, cells int
}

func (inv *inventory) addAmmo(i *i_item) {
	inv.bullets += i.ammoData.bullets
	inv.shells += i.ammoData.shells
	inv.rockets += i.ammoData.rockets
	inv.cells += i.ammoData.cells
}

func (inv *inventory) addItem(i *i_item) {
	if i.getType() == "ammo" {
		inv.addAmmo(i)
		return
	}
	inv.items = append(inv.items, i)
}

func (inv *inventory) removeItem(i *i_item) {
	for j := 0; j < len(inv.items); j++ {
		if i == inv.items[j] {
			inv.items = append(inv.items[:j], inv.items[j+1:]...)
		}
	}
}
