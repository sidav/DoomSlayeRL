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
		medicalData: &i_medicalData{healAmount:4, ignoresMaximum:false}}
	case "stimpack":
		i = i_item{ccell: &consoleCell{'+', tcell_wrapper.RED}, name: name, instantlyPickupable: false,
			medicalData: &i_medicalData{healAmount:15, ignoresMaximum:false}}
	case "small medikit":
		i = i_item{ccell: &consoleCell{'+', tcell_wrapper.DARK_RED}, name: name, instantlyPickupable: false,
			medicalData: &i_medicalData{healAmount:15, ignoresMaximum:false}}
	case "soulsphere":
		i = i_item{ccell: &consoleCell{'o', tcell_wrapper.BLUE}, name: name, instantlyPickupable: true,
			medicalData: &i_medicalData{healAmount:100, ignoresMaximum:true}}

	// armor
	case "green armor":
		i = i_item{ccell: &consoleCell{'[', tcell_wrapper.GREEN}, name: name, instantlyPickupable: false,
			armorData: &i_armorData{maxArmor: 100, damageConsumingPercent: 25}}
	case "red armor":
		i = i_item{ccell: &consoleCell{'[', tcell_wrapper.RED}, name: name, instantlyPickupable: false,
			armorData: &i_armorData{maxArmor: 125, damageConsumingPercent: 40}}
	case "blue armor":
		i = i_item{ccell: &consoleCell{'[', tcell_wrapper.BLUE}, name: name, instantlyPickupable: false,
			armorData: &i_armorData{maxArmor: 200, damageConsumingPercent: 66}}

	// ammo
	case "clip":
		i = i_item{ccell: &consoleCell{'"', tcell_wrapper.DARK_YELLOW}, name: name, instantlyPickupable: true, ammoData: &i_ammoData{ammo: [4]int{6, 0, 0, 0}}}
	case "cell":
		i = i_item{ccell: &consoleCell{'"', tcell_wrapper.DARK_CYAN}, name: name, instantlyPickupable: true, ammoData: &i_ammoData{ammo: [4]int{0, 0, 0, 5}}}
	case "shells":
		i = i_item{ccell: &consoleCell{'"', tcell_wrapper.DARK_RED}, name: name, instantlyPickupable: true, ammoData: &i_ammoData{ammo: [4]int{0, 4, 0, 0}}}
	case "ammunition crate":
		i = i_item{ccell: &consoleCell{'=', tcell_wrapper.DARK_MAGENTA}, name: name, ammoData: &i_ammoData{ammo: [4]int{20, 10, 1, 5}}}


	// weapons
		//BULLS:
	case "pistol":
		i = i_item{ccell: &consoleCell{')', tcell_wrapper.BEIGE}, name: name,
			weaponData: &i_weaponData{maxammo: 6, hitscanData: &w_hitscan{damageDice: &dice{dnum: 1, dval: 6, dmod: 0}}}}
	case "heavy pistol":
		i = i_item{ccell: &consoleCell{')', tcell_wrapper.BEIGE}, name: name,
			weaponData: &i_weaponData{maxammo: 5, hitscanData: &w_hitscan{damageDice: &dice{dnum: 2, dval: 5, dmod: 0}}}}
	case "assault rifle":
		i = i_item{ccell: &consoleCell{'\\', tcell_wrapper.BEIGE}, name: name,
			weaponData: &i_weaponData{maxammo: 20, hitscanData: &w_hitscan{shotsPerAttack: 3, damageDice: &dice{dnum: 1, dval: 6, dmod: 0}}}}
	case "chaingun":
		i = i_item{ccell: &consoleCell{')', tcell_wrapper.YELLOW}, name: name,
			weaponData: &i_weaponData{maxammo: 20, hitscanData: &w_hitscan{shotsPerAttack: 4, damageDice: &dice{dnum: 2, dval: 3, dmod: 0}}}}
	case "bolt-action rifle":
		i = i_item{ccell: &consoleCell{')', tcell_wrapper.DARK_GREEN}, name: name,
			weaponData: &i_weaponData{maxammo: 1, hitscanData: &w_hitscan{damageDice: &dice{dnum: 5, dval: 3, dmod: 0}}}}
		// SHELLS:
	case "shotgun":
		i = i_item{ccell: &consoleCell{')', tcell_wrapper.BLUE}, name: name,
			weaponData: &i_weaponData{ammoType: AMMO_SHEL, maxammo: 5, hitscanData: &w_hitscan{pelletsPerShot: 5, spreadAngle: 30, damageDice: &dice{dnum: 2, dval: 3, dmod: 0}}}}
	case "super shotgun":
		i = i_item{ccell: &consoleCell{')', tcell_wrapper.DARK_RED}, name: name,
			weaponData: &i_weaponData{ammoType: AMMO_SHEL, maxammo: 1, hitscanData: &w_hitscan{pelletsPerShot: 16, spreadAngle: 45, damageDice: &dice{dnum: 2, dval: 3, dmod: 0}}}}
	case "Pancor Jackhammer":
		i = i_item{ccell: &consoleCell{'\\', tcell_wrapper.DARK_RED}, name: name,
			weaponData: &i_weaponData{ammoType: AMMO_SHEL, maxammo: 10, hitscanData: &w_hitscan{shotsPerAttack: 3, pelletsPerShot: 4, spreadAngle: 30, damageDice: &dice{dnum: 2, dval: 3, dmod: 0}}}}
		// CELLS:
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
