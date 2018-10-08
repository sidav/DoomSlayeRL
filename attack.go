package main

import (
	"DoomSlayeRL/routines"
	"fmt"
)

func m_meleeAttack(attacker *p_pawn, victim *p_pawn, d *dungeon) {
	if attacker.isPlayer() && victim.aiData != nil && victim.aiData.state == AI_STAGGERED {
		m_gloryKill(attacker, victim, d)
		return
	}
	damage := attacker.meleeData.damageDice.roll()
	victim.receiveDamage(damage)
	attacker.spendTurnsForAction(turnCostFor("melee_attack"))
	log.appendMessage(fmt.Sprintf("%s %s %s! (%d damage)", attacker.name, attacker.meleeData.meleeAttackString, victim.name, damage))
}

func (victim *p_pawn) receiveDamage(damage int) { //deals with armor, staggered state etc
	const (
		STAGGER_PERCENT_THRESHOLD = 30
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
		// TODO: multi-shot weapons
		m_traceBullet(attacker, vx, vy, dung)
	}
	if aw.weaponData.maxammo > 0 {
		aw.weaponData.ammo -= 1 // TODO: investigate
	}
}

func m_traceBullet(attacker *p_pawn, tox, toy int, d *dungeon) {
	aw := attacker.weaponInHands
	ax, ay := attacker.getCoords()
	damage := aw.weaponData.hitscanData.damageDice.roll()
	traceLine := routines.GetLineOver(ax, ay, tox, toy, 20)
	for i, cell := range traceLine {
		if i == 0 {
			continue
		}
		victim := d.getPawnAt(cell.X, cell.Y)
		renderBullet(cell.X, cell.Y, tox, toy, d)
		if victim != nil {
			// TODO: miss shots
			victim.receiveDamage(damage)
			log.appendMessagef("The %s is hit!", victim.name)
			return
		}
		if d.isTileOpaque(cell.X, cell.Y) {
			return
		}
	}

}
