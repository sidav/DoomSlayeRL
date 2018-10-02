package main

func p_createPawn(name string, x, y int) *p_pawn {
	var p p_pawn
	switch name {
	case "player":
		p = p_pawn{appearance: '@', name: "you", maxhp: 100, playerData: &p_playerData{},
			meleeData: &p_meleeAttackData{meleeAttackString: "punch", damageDice: &dice{dnum: 2, dval: 6, dmod: 0}}, inventory: &inventory{}}
	case "zombie":
		p = p_pawn{appearance: 'z', name: name, maxhp: 10, meleeData: &p_meleeAttackData{meleeAttackString: "hits", damageDice: &dice{dnum: 1, dval: 6, dmod: 0}}}
	case "imp":
		p = p_pawn{appearance: 'i', name: name, maxhp: 25, meleeData: &p_meleeAttackData{meleeAttackString: "claws", damageDice: &dice{dnum: 3, dval: 5, dmod: 1}}}
	case "archvile":
		p = p_pawn{appearance: 'A', name: name, maxhp: 125, meleeData: &p_meleeAttackData{meleeAttackString: "burns", damageDice: &dice{dnum: 10, dval: 2, dmod: 1}}}
	default:
		p = p_pawn{appearance: '?', name: "Unknown monster " + name, maxhp: 25, meleeData: &p_meleeAttackData{meleeAttackString: "claws", damageDice: &dice{dnum: 3, dval: 5, dmod: 1}}}
	}
	p.x = x
	p.y = y
	p.hp = p.maxhp
	p.aiData = &p_aiData{}
	return &p
}

func p_createPlayer(x, y int) *p_pawn {
	var p p_pawn
	p = p_pawn{appearance: '@', name: "you", maxhp: 100, playerData: &p_playerData{},
	meleeData: &p_meleeAttackData{meleeAttackString: "punch", damageDice: &dice{dnum: 2, dval: 6, dmod: 0}}, inventory: &inventory{}}
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
