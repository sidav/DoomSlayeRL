package main

import "fmt"

func m_meleeAttack(attacker *p_pawn, victim *p_pawn) {
	damage := attacker.meleeData.rollForDamage()
	victim.hp -= damage
	log.appendMessage(fmt.Sprintf("%s %s %s! (%d damage)", attacker.name, attacker.meleeData.meleeAttackString, victim.name, damage))
}
