package main

func p_createPawn(name string, x, y int) p_pawn {
	var p p_pawn
	switch name {
	case "player":
		p = p_pawn{appearance: '@', name: "you", maxhp: 100, playerData: &p_playerData{}, melee: &p_meleeAttackData{meleeAttackString: "punch", dnum: 2, dval: 6, dmod: 0}}
	case "zombie":
		p = p_pawn{appearance: 'z', name: name, maxhp: 10, melee: &p_meleeAttackData{meleeAttackString: "hits", dnum: 1, dval: 6, dmod: 0}}
	case "imp":
		p = p_pawn{appearance: 'i', name: name, maxhp: 25, melee: &p_meleeAttackData{meleeAttackString: "claws", dnum: 3, dval: 5, dmod: 1}}
	}
	p.x = x
	p.y = y
	p.hp = p.maxhp
	return p
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
