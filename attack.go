package main

import "fmt"

func m_meleeAttack(attacker *p_pawn, victim *p_pawn) {
	damage := attacker.meleeData.damageDice.roll()
	victim.hp -= damage
	log.appendMessage(fmt.Sprintf("%s %s %s! (%d damage)", attacker.name, attacker.meleeData.meleeAttackString, victim.name, damage))
}

func m_rangedAttack(attacker *p_pawn, victim *p_pawn) {

}
