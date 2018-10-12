package main

type (
	p_meleeAttackData struct {
		meleeAttackString string
		// 3d6 + 1 == dnum d dval + dmod
		damageDice *dice
	}

	p_playerData struct {
		lastSpentTimeAmount int
	}

	p_pawn struct {
		ccell                                      *consoleCell
		hp, maxhp, x, y, nextTurnToAct, sightRange int
		name                                       string
		meleeData                                  *p_meleeAttackData
		playerData                                 *p_playerData
		wearedArmor                                *i_item
		weaponInHands                              *i_item
		inventory                                  *inventory
		aiData                                     *p_aiData
	}
)

func (p *p_pawn) isDead() bool {
	return p.hp <= 0
}

func (p *p_pawn) canMelee() bool {
	return p.meleeData != nil
}

func (p *p_pawn) canShoot() bool {
	return p.weaponInHands != nil
}

func (p *p_pawn) spendTurnsForAction(turns int) {
	p.nextTurnToAct = CURRENT_TURN + turns
	if p.playerData != nil {
		p.playerData.lastSpentTimeAmount = turns
	}
}

func (p *p_pawn) isTimeToAct() bool {
	return p.nextTurnToAct <= CURRENT_TURN
}

func (p *p_pawn) isPlayer() bool {
	return p.playerData != nil
}

func (p *p_pawn) getCoords() (int, int) {
	return p.x, p.y
}

func (p *p_pawn) getHpPercent() int {
	return p.hp * 100 / p.maxhp
}
