package main

import "DoomSlayeRL/routines"

type inventory struct {
	items                           []*i_item
	maxItems                        int
	ammo [4]int
}

func (inv *inventory) _addAmmo(itm *i_item) {
	for i := 0; i < 4; i++ {
		inv.ammo[i] += itm.ammoData.ammo[i]
	}
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

func (inv *inventory) selectItem() *i_item {
	items := inv.items
	index := routines.ShowSingleChoiceMenu("INVENTORY", "Your items:", inv.getNamesSliceForAllItems())
	if index != -1 {
		return items[index]
	}
	return nil
}
