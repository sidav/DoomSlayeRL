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

func m_gloryKill(attacker *p_pawn, victim *p_pawn) { // unused yet
	attacker.spendTurnsForAction(turnCostFor("glory_kill"))
	victim.hp = -666
	log.appendMessage(fmt.Sprintf("You glory kill the %s!", victim.name))
}

func (victim *p_pawn) receiveDamage(damage int) { //deals with armor, staggered state etc
	const (
		STAGGER_PERCENT_THRESHOLD = 50
		STAGGERED_TIME_AMOUNT     = 60
	)
	victim.hp -= damage
	if victim.isPlayer() == false {
		if victim.getHpPercent() < STAGGER_PERCENT_THRESHOLD {
			victim.aiData.state = AI_STAGGERED
			victim.aiData.stateTimeoutTurn = CURRENT_TURN + STAGGERED_TIME_AMOUNT
		}
	}
}

func m_rangedAttack(attacker *p_pawn, vx, vy int, dung *dungeon) {
	aw := attacker.weaponInHands
	ax, ay := attacker.getCoords()
	// vx, vy := victim.getCoords()
	if aw.weaponData.getType() == "projectile" {
		attacker.spendTurnsForAction(turnCostFor("ranged_attack"))
		proj := aw.weaponData.createProjectile(ax, ay, vx, vy)
		dung.addProjectileToList(proj)
		log.appendMessagef("%s shoots!", attacker.name)
	}
	if aw.weaponData.getType() == "hitscan" {
		attacker.spendTurnsForAction(turnCostFor("ranged_attack"))
		damage := aw.weaponData.hitscanData.damageDice.roll() // what a mess of receiver methods...
		victim := dung.getPawnAt(vx, vy)
		if victim != nil {
			victim.receiveDamage(damage)
			log.appendMessagef("Placeholder: %s is hit!", victim.name)
		}
	}
	if aw.weaponData.maxammo > 0 {
		aw.weaponData.ammo -= 1 // TODO: investigate
	}
}
