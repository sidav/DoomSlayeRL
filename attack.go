package main

import "fmt"

func m_meleeAttack(attacker *p_pawn, victim *p_pawn) {
	if attacker.isPlayer() && victim.aiData != nil && victim.aiData.state == AI_STAGGERED {
		m_gloryKill(attacker, victim)
		return 
	}
	damage := attacker.meleeData.damageDice.roll()
	victim.receiveDamage(damage)
	log.appendMessage(fmt.Sprintf("%s %s %s! (%d damage)", attacker.name, attacker.meleeData.meleeAttackString, victim.name, damage))
}

func m_gloryKill(attacker *p_pawn, victim *p_pawn){ // unused yet
	victim.hp = -victim.maxhp
	log.appendMessage(fmt.Sprintf("You glory kill the %s!", victim.name))
}

func (victim *p_pawn) receiveDamage(damage int) { //deals with armor, staggered state etc
	victim.hp -= damage
	if victim.isPlayer() == false {
		if victim.hp * 100 / victim.maxhp < 25 { //less than 25 percent hp remaining
			victim.aiData.state = AI_STAGGERED
		}
	}
}

func m_rangedAttack(attacker *p_pawn, victim *p_pawn) {

}
