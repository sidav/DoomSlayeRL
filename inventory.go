package main

import "GoRoguelike/routines"

type inventory struct {
	items                           []*i_item
	maxItems                        int
	bullets, shells, rockets, cells int
}

func (inv *inventory) _addAmmo(i *i_item) {
	inv.bullets += i.ammoData.ammo[AMMO_BULL]
	inv.shells += i.ammoData.ammo[AMMO_SHEL]
	inv.rockets += i.ammoData.ammo[AMMO_RCKT]
	inv.cells += i.ammoData.ammo[AMMO_CELL]
}

func (inv *inventory) addItem(i *i_item) {
	if i.getType() == "ammo" {
		inv._addAmmo(i)
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

func (inv *inventory) getNamesSliceForAllItems() []string{
	var slice []string
	slice = make([]string, 0)
	for i := 0; i < len(inv.items); i++ {
		slice = append(slice, inv.items[i].name)
	}
	return slice
}

func (inv *inventory) selectItem(owner *p_pawn) {
	routines.ShowSingleChoiceMenu("INVENTORY", inv.getNamesSliceForAllItems())
}
