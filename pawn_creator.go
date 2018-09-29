package main

func p_createPawn(name string, x, y int) pawn {
	var p pawn
	switch name {
	case "player":
		p = pawn{appearance: '@', name: "you", maxhp: 100, x: x, y: y}
	case "zombie":
		p = pawn{appearance: 'z', name: name, maxhp: 10, x: x, y: y}
	case "imp":
		p = pawn{appearance: 'i', name: name, maxhp: 25, x: x, y: y}
	}
	p.hp = p.maxhp
	return p
}

//func p_createPlayer(x, y int) pawn {
//	return pawn{appearance:'@', maxhp:100, hp:100, x: x, y: y}
//}
//
//func p_createZombie(x, y int) pawn {
//	return pawn{appearance:'z', maxhp:10, hp:10, x: x, y: y}
//}
//
//func p_createImp(x, y int) pawn {
//	return pawn{appearance:'i', maxhp:25, hp:25, x: x, y: y}
//}
