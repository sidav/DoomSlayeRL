package main

import "fmt"

func m_meleeAttack(attacker *p_pawn, victim *p_pawn) {
	if attacker.isPlayer() && victim.aiData != nil && victim.aiData.state == AI_STAGGERED {
		m_gloryKill(attacker, victim)
		return
	}
	damage := attacker.meleeData.damageDice.roll()
	victim.receiveDamage(damage)
	attacker.spendTurnsForAction(turnCostFor("melee_attack"))
	log.appendMessage(fmt.Sprintf("%s %s %s! (%d damage)", attacker.name, attacker.meleeData.meleeAttackString, victim.name, damage))
}

func m_gloryKill(attacker *p_pawn, victim *p_pawn){ // unused yet
	attacker.spendTurnsForAction(turnCostFor("glory_kill"))
	victim.hp = -666
	log.appendMessage(fmt.Sprintf("You glory kill the %s!", victim.name))
}

func (victim *p_pawn) receiveDamage(damage int) { //deals with armor, staggered state etc
	const (
		STAGGER_PERCENT_THRESHOLD = 50
		STAGGERED_TIME_AMOUNT = 60
	)
	victim.hp -= damage
	if victim.isPlayer() == false {
		if victim.getHpPercent() < STAGGER_PERCENT_THRESHOLD {
			victim.aiData.state = AI_STAGGERED
			victim.aiData.stateTimeoutTurn = CURRENT_TURN + STAGGERED_TIME_AMOUNT
		}
	}
}

func m_rangedAttack(attacker *p_pawn, victim *p_pawn) {

}
