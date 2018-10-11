package main

import cw "TCellConsoleWrapper/tcell_wrapper"

func p_createPawn(name string, x, y int) *p_pawn {
	var p p_pawn
	switch name {
	case "unwilling":
		p = p_pawn{ccell:&consoleCell{'z', cw.BEIGE}, name: name, maxhp: 15,
		meleeData: &p_meleeAttackData{meleeAttackString: "hits", damageDice: &dice{dnum: 2, dval: 3, dmod: 0}}}
	case "zombie soldier":
		p = p_pawn{ccell:&consoleCell{'z', cw.DARK_YELLOW}, name: name, maxhp: 20,
		meleeData: &p_meleeAttackData{meleeAttackString: "hits", damageDice: &dice{dnum: 1, dval: 6, dmod: 0}},
		weaponInHands: &i_item{weaponData: &i_weaponData{
			hitscanData: &w_hitscan{spreadAngle: 30, damageDice: &dice{1,3,3}}}}}
	case "zombie sergeant":
		p = p_pawn{ccell:&consoleCell{'z', cw.RED}, name: name, maxhp: 20,
			meleeData: &p_meleeAttackData{meleeAttackString: "hits", damageDice: &dice{dnum: 1, dval: 6, dmod: 0}},
			weaponInHands: &i_item{weaponData: &i_weaponData{
				hitscanData: &w_hitscan{pelletsPerShot: 4, spreadAngle: 30, damageDice: &dice{2,3,0}}}}}
	case "imp":
		p = p_pawn{ccell:&consoleCell{'i', cw.RED}, name: name, maxhp: 25,
			meleeData: &p_meleeAttackData{meleeAttackString: "claws", damageDice: &dice{dnum: 3, dval: 5, dmod: 1}},
			weaponInHands: &i_item{weaponData: &i_weaponData{
				projectileExample: &projectile{turnsForOneTile: 6, damageDice: &dice{3, 3, 3}}}}}
	default:
		p = p_pawn{ccell:&consoleCell{'?', cw.MAGENTA}, name: "Unknown monster " + name, maxhp: 25, meleeData: &p_meleeAttackData{meleeAttackString: "claws", damageDice: &dice{dnum: 3, dval: 5, dmod: 1}}}
	}
	p.x = x
	p.y = y
	p.hp = p.maxhp
	p.aiData = &p_aiData{}
	return &p
}

func p_createPlayer(x, y int) *p_pawn {
	var p p_pawn
	p = p_pawn{ccell:&consoleCell{'@', cw.GREEN}, name: "you", maxhp: 100, playerData: &p_playerData{},
		weaponInHands: i_createItem("pistol", x, y),
		meleeData: &p_meleeAttackData{meleeAttackString: "punch", damageDice: &dice{dnum: 2, dval: 6, dmod: 0}},
		inventory: &inventory{maxItems: 4, maxammo: [4]int{20, 10, 1, 10}}}
	p.x = x
	p.y = y
	p.hp = p.maxhp
	return &p
}

//func p_createPlayer(x, y int) p_pawn {
//	return p_pawn{appearance:'@', maxhp:100, hp:100, x: x, y: y}
//}
//
//func p_createZombie(x, y int) p_pawn {
//	return p_pawn{appearance:'z', maxhp:10, hp:10, x: x, y: y}
//}
//
//func p_createImp(x, y int) p_pawn {
//	return p_pawn{appearance:'i', maxhp:25, hp:25, x: x, y: y}
//}
