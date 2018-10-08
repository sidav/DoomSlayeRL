package main

import (
	"TCellConsoleWrapper/tcell_wrapper"
	"fmt"
)

func i_createItem(name string, x, y int) *i_item {
	var i i_item
	switch name {

	// medicals
	case "health bonus":
		i = i_item{ccell: &consoleCell{'+', tcell_wrapper.CYAN}, name: name, instantlyPickupable: true,
		medicalData: &i_medicalData{healAmount:2, ignoresMaximum:false}}


	// ammo
	case "clip":
		i = i_item{ccell: &consoleCell{'"', tcell_wrapper.DARK_YELLOW}, name: name, instantlyPickupable: true, ammoData: &i_ammoData{ammo: [4]int{6, 0, 0, 0}}}
	case "cell":
		i = i_item{ccell: &consoleCell{'"', tcell_wrapper.DARK_CYAN}, name: name, instantlyPickupable: true, ammoData: &i_ammoData{ammo: [4]int{0, 0, 0, 5}}}
	case "ammunition crate":
		i = i_item{ccell: &consoleCell{'=', tcell_wrapper.DARK_MAGENTA}, name: name, ammoData: &i_ammoData{ammo: [4]int{10, 10, 1, 5}}}


	// weapons
	case "pistol":
		i = i_item{ccell: &consoleCell{')', tcell_wrapper.BEIGE}, name: name,
			weaponData: &i_weaponData{maxammo: 6, hitscanData: &w_hitscan{damageDice: &dice{dnum: 1, dval: 6, dmod: 0}}}}
	case "bolt-action rifle":
		i = i_item{ccell: &consoleCell{')', tcell_wrapper.DARK_GREEN}, name: name,
			weaponData: &i_weaponData{maxammo: 1, hitscanData: &w_hitscan{damageDice: &dice{dnum: 5, dval: 3, dmod: 0}}}}
	case "gauss rifle":
		i = i_item{ccell: &consoleCell{')', tcell_wrapper.DARK_CYAN}, name: name,
			weaponData: &i_weaponData{ammoType: AMMO_CELL, maxammo: 1, hitscanData: &w_hitscan{damageDice: &dice{dnum: 10, dval: 6, dmod: 10}}}}

	default:
		i = i_item{ccell: &consoleCell{'?', tcell_wrapper.MAGENTA}, name: "UNKNOWN ITEM " + name}
	}
	if i.getType() == "weapon" {
		i.weaponData.ammo = i.weaponData.maxammo
	}
	i.x = x
	i.y = y
	return &i
}

func i_createCorpseFor(p *p_pawn) *i_item {
	x, y := p.x, p.y
	name := fmt.Sprintf("%s corpse", p.name)
	return &i_item{name: name, x: x, y: y, ccell: &consoleCell{'%', tcell_wrapper.DARK_RED}}
}

//func i_createWeapon(name string, x, y int) *i_item {
//	var i i_item
//	switch name {
//	case "pistol":
//		i = i_item{appearance: ')', name: name, weaponData: &i_weaponData{maxammo: 6, damageDice: &dice{dnum: 1, dval: 6, dmod: 0}}}
//	default:
//		i = i_item{appearance: '?', name: "UNKNOWN ITEM " + name}
//	}
//	i.weaponData.ammo = i.weaponData.maxammo
//	i.x = x
//	i.y = y
//	return &i
//}
