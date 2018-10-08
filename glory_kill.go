package main

import (
	"GoRoguelike/routines"
	"fmt"
)

func m_gloryKill(attacker *p_pawn, victim *p_pawn, d *dungeon) {

	const(
		THRESHOLD_TO_DROP_HEALTH = 50
		HEALTH_BONUS_CHANCE = 65
	)

	attacker.spendTurnsForAction(turnCostFor("glory_kill"))
	victim.hp = -666
	log.appendMessage(fmt.Sprintf("You glory kill the %s!", victim.name))
	// spawn health bonuses
	d.addItemToFloor(i_createItem("health bonus", victim.x, victim.y))
	if attacker.getHpPercent() < THRESHOLD_TO_DROP_HEALTH {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if routines.RandomPercent() < HEALTH_BONUS_CHANCE {
					d.addItemToFloor(i_createItem("health bonus", victim.x+x, victim.y+y))
				}
			}
		}
	}
}
