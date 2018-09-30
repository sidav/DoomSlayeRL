package main

import "fmt"

func movePawn(p *pawn, d *dungeon, x, y int) {
	// px, py := p.x, p.y
	nx, ny := p.x+x, p.y+y
	if d.isTilePassableAndNotOccupied(nx, ny) {
		p.x += x
		p.y += y
	}
}

func moveOrMeleeAttackPawn(p *pawn, d *dungeon, x, y int) {
	nx, ny := p.x+x, p.y+y
	if d.isTilePassableAndNotOccupied(nx, ny) {
		movePawn(p, d, x, y)
	} else if d.isPawnPresent(nx, ny) {
		victim := d.getPawnAt(nx, ny)
		victim.hp -= 10
		log.appendMessage(fmt.Sprintf("You punch %s!", victim.name))
	}
}

func checkDeadPawns(d *dungeon) {
	var indicesOfPawnsToRemove []int
	for i := 0; i < len(d.pawns); i++ {
		p := &d.pawns[i]
		if p.hp < 0 {
			indicesOfPawnsToRemove = append(indicesOfPawnsToRemove, i)
		}
	}
	for i := 0; i < len(indicesOfPawnsToRemove); i++ {
		index := indicesOfPawnsToRemove[i]
		log.appendMessage(fmt.Sprintf("%s drops dead!", d.pawns[index].name))
		//let's create a corpse
		d.items = append(d.items, i_createCorpseFor(&d.pawns[index]))
		d.pawns = append(d.pawns[:index], d.pawns[index+1:]...) // this fucking magic removes indexth element from a slice
	}
}
