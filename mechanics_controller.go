package main

import (
	"fmt"
)

func m_movePawn(p *p_pawn, d *dungeon, x, y int) {
	// px, py := p.x, p.y
	nx, ny := p.x+x, p.y+y
	if d.isTilePassableAndNotOccupied(nx, ny) {
		p.x += x
		p.y += y
		if x*y != 0 { // diagonal movement
			p.spendTurnsForAction(turnCostFor("step_diag"))
		} else { // non-diagonal movement
			p.spendTurnsForAction(turnCostFor("step"))
		}
	}
}

func m_moveOrMeleeAttackPawn(p *p_pawn, d *dungeon, x, y int) {
	nx, ny := p.x+x, p.y+y
	if d.isTilePassableAndNotOccupied(nx, ny) {
		m_movePawn(p, d, x, y)
	} else if d.isPawnPresent(nx, ny) {
		victim := d.getPawnAt(nx, ny)
		if victim.isPlayer() || p.isPlayer() {
			m_meleeAttack(p, victim, d)
		}
	}
}

func m_moveProjectiles(d *dungeon) {
	for _, p := range d.projectiles {
		px, py := p.x, p.y
		if d.isPawnPresent(px, py) {
			victim := d.getPawnAt(px, py)
			dmg := p.damageDice.roll()
			victim.receiveDamage(dmg)
			log.appendMessage(fmt.Sprintf("%s is hit! (%d damage)", victim.name, dmg))
			d.removeProjectileFromList(p)
			continue
		}
		if d.isTileOpaque(px, py) {
			d.removeProjectileFromList(p)
			continue
		}
		if p.nextTurnToMove <= CURRENT_TURN {
			p.moveNextTile()
			p.nextTurnToMove = CURRENT_TURN + p.turnsForOneTile
		}
	}
}

func checkDeadPawns(d *dungeon) {
	var indicesOfPawnsToRemove []int
	for i := 0; i < len(d.pawns); i++ {
		p := d.pawns[i]
		if p.hp < 0 {
			indicesOfPawnsToRemove = append(indicesOfPawnsToRemove, i)
		}
	}
	for i := 0; i < len(indicesOfPawnsToRemove); i++ {
		index := indicesOfPawnsToRemove[i]
		pawn := d.pawns[index]
		// add blood splats if neccessary
		if pawn.hp == -666 { // exactly 666 hp means that this enemy was glory killed
			d.addBloodSplats(pawn.x, pawn.y, 1)
		} else {
			negHpPercent := -pawn.getHpPercent()
			if negHpPercent < 50 {
				log.appendMessage(fmt.Sprintf("%s drops dead!", d.pawns[index].name))
				//let's create a corpse
				d.addItemToFloor(i_createCorpseFor(d.pawns[index]))
			} else {
				log.appendMessage(fmt.Sprintf("%s is obliterated!", d.pawns[index].name))
				d.addBloodSplats(pawn.x, pawn.y, 2)
			}
		}
		d.removePawn(d.pawns[index])
	}
}
